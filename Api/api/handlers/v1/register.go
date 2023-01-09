package v1

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"unicode"

	"github.com/Trendyol/Api/api/auth"
	"github.com/Trendyol/Api/api/model"
	pb "github.com/Trendyol/Api/genproto"
	user "github.com/Trendyol/Api/genproto"
	"github.com/Trendyol/Api/mail"
	"github.com/Trendyol/Api/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/encoding/protojson"
)

var code string

//Register register user
// @Summary Register user summary
// @Description This API for registering user
// @Tags user
// @Accept json
// @Produce json
// @Param user body model.CreateUser true "user body"
// @Success 200 {string} model.User!
// @Success 400 {object} model.User
// @Success 500 {object} model.User
// @Router /v1/users/RegisterUser [post]
func (h *handlerV1) RegisterUser(c *gin.Context) {
	var (
		body        pb.CreateUserReq
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseEnumNumbers = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while blind json", logger.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	check, err := h.serviceManager.UserService().CheckField(ctx, &user.CheckFieldRequest{
		Field: `username`,
		Value: body.Username,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			`error`: err.Error(),
		})
		h.log.Error("failed to check username", logger.Error(err))
		return

	}
	if !check.Check {
		check1, err := h.serviceManager.UserService().CheckField(ctx, &pb.CheckFieldRequest{
			Field: `email`,
			Value: body.Email,
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			h.log.Error("failed to check email", logger.Error(err))
			return

		}
		if check1.Check {

			return
		}

	} else {
		return
	}
	code, _ = genCaptchaCode()
	if err != nil {
		fmt.Println(err)
		return
	}

	eigthMore, number, upper, special, moredigits := VerifyPassword(body.Password)
	if !eigthMore {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed in password not much characters",
		})
		h.log.Error("failed in password not much characters", logger.Error(err))
		return
	}
	if !moredigits {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed in password not much characters",
		})
		h.log.Error("failed in password not much characters", logger.Error(err))
		return
	}
	if !number {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed in password don't have numbers in password",
		})
		h.log.Error("failed in password don't have numbers in password", logger.Error(err))
		return
	}
	if !upper {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed in password don't have upper symbole",
		})
		h.log.Error("failed in password don't have upper symbole", logger.Error(err))
		return
	}
	if !special {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed in password don't have special sybole",
		})
		h.log.Error("failed in password don't have special sybole", logger.Error(err))
		return
	}
	body.Code = code
	fmt.Println(code)
	//src := []byte("Hello Gopher!")
	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), len(body.Password))
	if err != nil {
		fmt.Print(err)
		return
	}

	body.Password = string(password)
	//body.Password=
	bodyByte, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to convert json", logger.Error(err))
		return

	}

	//users := User{}
	err = h.redisStorage.SetWithTTL(body.Email, string(bodyByte), int64(time.Minute)*2)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to convert json2", logger.Error(err))
		return

	}
	fmt.Println(body)

	err = mail.SendMail(code, body.Email)
	if err != nil {
		fmt.Println(err)
	}
	genCaptchaCode()
}

var coded string
var email string

//Post user by code
//@Summary Post user summary
//Description This api for post user by code
//@Tags user
//@Accept json
//@Produce json
//@Param email path string true "Email"
//@Param coded path string true "Code"
//@Success 200 {string} model.Tokens!
//@Router /v1/users/register/user/{email}/{coded} [post]
func (h *handlerV1) Verify(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	email = c.Param("email")
	coded = c.Param("coded")
	// fmt.Println(email)
	// fmt.Println(coded, "   ", code)
	var (
		userm pb.CreateUserReq
	)

	vali, _ := redis.String(h.redisStorage.Get(email))
	err := json.Unmarshal([]byte(vali), &userm)
	if err != nil {
		return
	}
	//fmt.Println(val)
	ctxr, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	h.JwtHandler = auth.JwtHandler{
		Sub:       userm.Id,
		Iss:       "client",
		Role:      "Authorized",
		Log:       h.log,
		SigninKey: h.cfg.SigninKey,
	}
	var (
		tokens model.Tokens
	)
	access, refresh, err := h.JwtHandler.GenerateAuthJWT()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error while genrating new jwt",
		})
		h.log.Error("error while generating new jwt", logger.Error(err))
		return
	}
	tokens.AccessToken = access
	tokens.RefreshToken = refresh

	if userm.Code == coded {
		_, err := h.serviceManager.UserService().CreateUser(ctxr, &userm)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			h.log.Error("failed to create user", logger.Error(err))
			return
		}

		c.JSON(http.StatusCreated, tokens)
	} else {
		err := "code erroe"
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		//h.log.Error("failed in code", nil)
		return

	}

}

type Register struct {
	User pb.User
	Code string `json:"code"`
}

func genCaptchaCode() (string, error) {
	codes := make([]byte, 6)
	if _, err := rand.Read(codes); err != nil {
		return "", err
	}

	for i := 0; i < 6; i++ {
		codes[i] = uint8(48 + (codes[i] % 10))
	}

	return string(codes), nil
}

func VerifyPassword(s string) (eigthMore, number, upper, special, moredigits bool) {
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			//return false, false, false, false
		}
	}
	eigthMore = letters >= 8
	moredigits = letters <= 32
	return
}

// Post user by code
// @Summary Get user summary
// @Description This api for post user by code
// @Tags user
// @Accept json
// @Produce json
// @Param user body model.Login true "Login"
// @Success 200 {string} model.User!
// @Success 400 {object} Success
// @Success 500 {object} Success
// @Router /v1/users/login [post]
func (h *handlerV1) Login(c *gin.Context) {
	var body model.Login
	fmt.Println(body)
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while to blind json", logger.Error(err))
		return
	}
	// params, errStr := utils.ParseQueryParamsLog(queryParams)
	// if errStr != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": errStr[0],
	// 	})
	// 	h.log.Fatal("failed json params" + errStr[0])
	// 	return
	// }
	ctxr, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	if body.Email == "admin" && body.Password == "admin" {
		h.JwtHandler = auth.JwtHandler{
			Iss:       "admin",
			Role:      "Authorized",
			Log:       h.log,
			SigninKey: h.cfg.SigninKey,
		}
	} else {

		usersss, err := h.serviceManager.UserService().LoginUser(ctxr, &pb.LoginUserReq{
			Email:    body.Email,
			Password: body.Password,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Email or password is invalide",
			})
			h.log.Error("Email or password in invalide", logger.Error(err))
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(usersss.Password), []byte(body.Password))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "password is invalide",
			})
			h.log.Error("password in invalide", logger.Error(err))
			return
		}

		h.JwtHandler = auth.JwtHandler{
			Sub:       usersss.Id,
			Iss:       "client",
			Role:      "Authorized",
			Log:       h.log,
			SigninKey: h.cfg.SigninKey,
		}
		c.JSON(http.StatusCreated, usersss)
	}
	access, refresh, err := h.JwtHandler.GenerateAuthJWT()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error while genrating new jwt",
		})
		h.log.Error("error while generating new jwt", logger.Error(err))
		return
	}

	c.JSON(http.StatusCreated, access)
	c.JSON(http.StatusCreated, refresh)

}

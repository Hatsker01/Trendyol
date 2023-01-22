package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Trendyol/Api/api/model"
	pb "github.com/Trendyol/Api/genproto"
	"github.com/Trendyol/Api/pkg/logger"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateUser ...
// @Summary CreateUser
// @Description This API for creating a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body model.CreateUser true "CreateUser"
// @Success 200 {object} model.User!
// @Success 400 {object} Success
// @Success 500 {object} Success
// @Router   /v1/users [post]
func (h *handlerV1) CreateUser(c *gin.Context) {
	var body pb.CreateUserReq

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while to blind json", logger.Error(err))
		return

	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	check, err := h.serviceManager.UserService().CheckField(ctx, &pb.CheckFieldRequest{
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

	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), len(body.Password))
	if err != nil {
		fmt.Print(err)
		return
	}

	body.Password = string(password)

	user, err := h.serviceManager.UserService().CreateUser(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating user json", logger.Error(err))
	}
	c.JSON(http.StatusCreated, user)
}

// UpdateUser ...
// @Summary UpdateUser
// @Description This API for updating a new user
// @Security BearerAuth
// @Tags user
// @Accept json
// @Produce json
// @Param user body model.User true "UpdateUser"
// @Success 200 {object} model.User!
// @Success 400 {object} Success
// @Success 500 {object} Success
// @Router /v1/users [put]
func (h *handlerV1) UpdateUser(c *gin.Context) {
	claims, err := GetClaims(*h, c)
	if err == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized access",
		})
		return
	}

	if claims.Role != "Authorized" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "token not found",
		})
		return
	}
	body := pb.User{}
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while to blind json", logger.Error(err))

	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	user, err := h.serviceManager.UserService().UpdateUser(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating user", logger.Error(err))
	}
	c.JSON(http.StatusAccepted, user)

}

// GetUserById ...
// @Summary GetUserById
// @Description This API for getting user by id
// @Security BearerAuth
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User_ID"
// @Success 200 {object} model.User!
// @Success 400 {object} Success
// @Success 500 {object} Success
// @Router /v1/user/getUserbyId/{id} [get]
func (h *handlerV1) GetUserById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")

	er := CheckClaims(h, c)
	if er == nil {
		fmt.Println(er)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while comparing token",
		})
		h.log.Fatal("failed in token check")
		return

	}
	// if err == nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "Unauthorized access",
	// 	})
	// 	return
	// }

	// fmt.Println("asdfasfa")
	// if claims. != "Authorized" {
	// 	c.JSON(http.StatusForbidden, gin.H{
	// 		"error": "token not found",
	// 	})
	// 	return
	// }
	// fmt.Println("asdfasfa")

	// var jspbMarshal protojson.MarshalOptions
	// jspbMarshal.UseProtoNames = true
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	responce, err := h.serviceManager.UserService().GetUserById(ctx, &pb.WithId{Id: guid})
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting user by id", logger.Error(err))
	}
	c.JSON(http.StatusAccepted, responce)
}

// GetAllUsers
// @Summary GetAllUsers
// @Description This API for getting all Users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} model.Users!
// @Success 400 {object} Success
// @Success 500 {object} Success
// @Router /v1/users/getAll [get]
func (h *handlerV1) GetAllUsers(c *gin.Context) {
	// er := CheckClaims(h, c)
	// if er == nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "Unauthorized access",
	// 	})
	// 	return
	// }
	// if er. != "Authorized" || claims.Iss != "admin" {
	// 	c.JSON(http.StatusForbidden, gin.H{
	// 		"error": "user not admin",
	// 	})
	// 	return
	// }
	cxt, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	users, err := h.serviceManager.UserService().GetAllUsers(cxt, &pb.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error while getting all users",
		})
		return
	}
	c.JSON(http.StatusAccepted, users)
}

// DeleteUserById
// @Summary Delete User By Id
// @Description This API for deleting user by ID
// @Security BearerAuth
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User_id"
// @Success 200 {object} model.User!
// @Success 400 {object} Success
// @Success 500 {object} Success
// @Router /v1/user/delete/{id} [delete]
func (h *handlerV1) DeleteUserById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	er := CheckClaims(h, c)
	if er == nil {
		fmt.Println(er)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while comparing token",
		})
		h.log.Fatal("failed in token check")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	user, err := h.serviceManager.UserService().DeleteUserById(ctx, &pb.WithId{Id: guid})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error while deleting user by id",
		})
		return
	}
	c.JSON(http.StatusAccepted, user)
}

// ChangePassword ...
// @Summary Change User Password
// @Description This API for changing user password
// @Security BearerAuth
// @Tags user
// @Accept json
// @Produce json
// @Param user body model.ChengePass true "ChangePassword"
// @Success 200 {object} model.ChangePassRes
// @Failure 400 {object} model.StandardErrorModel
// @Failure 500 {object} model.StandardErrorModel
// @Router /v1/user/changePass [put]
func (h *handlerV1) ChengePass(c *gin.Context) {
	er := CheckClaims(h, c)
	if er == nil {
		newResponse(c, http.StatusUnauthorized, "failed while checking token")
		h.log.Error("failed while checking token")
		return
	}
	var body model.ChengePass
	err := c.ShouldBindJSON(&body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "failed while blinding JSON")
		h.log.Error("failed while blindig json", logger.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	user, err := h.serviceManager.UserService().GetUserById(ctx, &pb.WithId{Id: body.Id})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, model.ErrIdNotFound.Error())
		h.log.Error("error while getting user by id", logger.Error(err))
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(body.NewPass), len(body.NewPass))
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "failed while changing pass")
		h.log.Error("error while changing pass and comparing")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.OldPass))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "password is invalide",
		})
		h.log.Error("password in invalide", logger.Error(err))
		return
	}
	if body.NewPass == body.VerifyNew {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
		defer cancel()

		pass, err := h.serviceManager.UserService().ChangePassword(ctx, &pb.ChangePassReq{
			Id:          user.Id,
			NewPassword: string(password),
		})
		if err != nil {
			newResponse(c, http.StatusInternalServerError, "failed while changing password")
			h.log.Error("failed while changing password", logger.Error(err))
			return
		}
		c.JSON(http.StatusAccepted, pass)
	} else if body.NewPass != body.VerifyNew {
		newResponse(c, http.StatusInternalServerError, "new password are not the same")
		h.log.Error("new password are not th same", logger.Error(err))
		return
	} else {
		newResponse(c, http.StatusInternalServerError, "password is wrong")
		h.log.Error("password is wrong", logger.Error(err))
		return
	}

}

type Success struct {
	Message string `json:"message"`
}

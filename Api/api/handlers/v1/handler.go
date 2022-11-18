package v1

import (
	"errors"
	"net/http"

	"github.com/Trendyol/Api/api/auth"
	"github.com/Trendyol/Api/api/model"
	"github.com/Trendyol/Api/config"
	"github.com/Trendyol/Api/pkg/logger"
	"github.com/Trendyol/Api/services"
	// "github.com/Trendyol/Api/storage/repo"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
	JwtHandler     auth.JwtHandler
}

type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
	jwtHandler     auth.JwtHandler
}

func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
		JwtHandler:     c.jwtHandler,
	}
}

func CheckClaims(h *handlerV1, c *gin.Context) jwt.MapClaims {
	var (
		ErrUnauthorized = errors.New("unauthorized")
		authorization   model.JwtRequestModel
		claims          jwt.MapClaims
		err             error
	)

	authorization.Token = c.GetHeader("Authorization")
	if c.Request.Header.Get("Authorization") == "" {
		c.JSON(http.StatusUnauthorized, ErrUnauthorized)
		h.log.Error("Unauthorized request:", logger.Error(err))

	}
	h.JwtHandler.Token = authorization.Token
	claims, err = h.JwtHandler.ExtractClaims()
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrUnauthorized)
		h.log.Error("token is invalid:", logger.Error(err))
		return nil
	}
	return claims
}

// func New1(c *HandlerV1Config)*handlerV1{
// 	return &handlerV1{
// 		log: c.Logger,
// 		serviceManager: c.ServiceManager,
// 		cfg: c.Cfg,
// 	}
// }

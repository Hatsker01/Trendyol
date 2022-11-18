package v1

import (
	"errors"
	"net/http"

	"github.com/Trendyol/api/config"
	"github.com/Trendyol/api/internal/service"
	"github.com/Trendyol/api/internal/transport/rest/auth"
	"github.com/Trendyol/api/internal/transport/rest/v1/model"
	"github.com/Trendyol/api/pkg/logger"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	serviceManager service.IServiceManager
	log            logger.Logger
	JwtHandler     auth.JwtHandler
	cfg            config.Config
}

// HandlerConfig ...
type HandlerV1Config struct {
	ServiceManager service.IServiceManager
	Logger         logger.Logger
	Jwthandler     auth.JwtHandler
	Cfg            config.Config
}

func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		serviceManager: c.ServiceManager,
		log:            c.Logger,
		JwtHandler:     c.Jwthandler,
		cfg:            c.Cfg,
	}
}
func (h *handlerV1) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initUsersRoutes(v1)

	}
}
func GetClaims(h *handlerV1, c *gin.Context) jwt.MapClaims {
	var (
		ErrUnauthorized = errors.New("unauthorized")
		authorization   model.JwtRequestModel
		claims          jwt.MapClaims
		err             error
	)
	authorization.Token = c.GetHeader("Authorized")
	if c.Request.Header.Get("Authorization") == "" {
		c.JSON(http.StatusUnauthorized, ErrUnauthorized)
		h.log.Error("Unauthorized request:", logger.Error(err))
	}
	h.JwtHandler.Token = authorization.Token
	claims, err = h.JwtHandler.ExtractClaims()
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrUnauthorized)
		h.log.Error("token is invaild:", logger.Error(err))
		return nil
	}
	return claims
}

package rest

import (
	"fmt"
	"net/http"

	"github.com/Trendyol/api/docs"
	v1 "github.com/Trendyol/api/internal/transport/rest/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	

	"github.com/Trendyol/api/config"
	"github.com/Trendyol/api/internal/service"
	"github.com/Trendyol/api/pkg/logger"
	"github.com/Trendyol/api/pkg/middleware"
)

type Handler struct {
	services *service.IServiceManager
	log      logger.Logger
}

func NewHandler(services *service.IServiceManager, log logger.Logger) *Handler {
	return &Handler{
		services: services,
		log:      log,
	}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.New()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		middleware.New(GinCorsMiddleware()),
	)
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTPHost, cfg.HTTPPort)
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	h.initAPI(router)
	return router
}

func (h *Handler) initAPI(router *gin.Engine) {

	handlerV1 := v1.New(&v1.HandlerV1Config{
		ServiceManager: *h.services,
		Logger: h.log,
	})
	api := router.Group("")
	{
		handlerV1.Init(api)
	}
}

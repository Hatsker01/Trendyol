package api

import (
	v1 "github.com/Trendyol/Api/api/handlers/v1"
	"github.com/Trendyol/Api/config"
	"github.com/Trendyol/Api/pkg/logger"
	"github.com/Trendyol/Api/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// New @BasePath /v1
// New ...
// @SecurityDefinitions.apikey BearerAuth
// @Description GetMyProfile
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()
	
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// jwtHandler := auth.JwtHandler{
	// 	SigninKey: option.Conf.SigninKey,
	// 	Log:       option.Logger,
	// }
	// router.Use(casbin.NewJwtRoleStruct( option.Conf, jwtHandler))

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")
	api.POST("/users/", handlerV1.CreateUser)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

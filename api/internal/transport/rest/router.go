// package rest
package rest

// import (
// 	"github.com/Trendyol/api/config"
// 	"github.com/Trendyol/api/internal/service"
// 	v1 "github.com/Trendyol/api/internal/transport/rest/v1"
// 	"github.com/Trendyol/api/pkg/logger"
// 	"github.com/gin-gonic/gin"

// 	// "github.com/gin-gonic/gin"
// 	swaggerFiles "github.com/swaggo/files"
// 	// _ "github.com/Trendyol/api/api/docs"
// 	ginSwagger "github.com/swaggo/gin-swagger"
// )

// type Option struct {
// 	Conf           config.Config
// 	Logger         logger.Logger
// 	ServiceManager service.IServiceManager
// }

// // New @BasePath /v1
// // New ...
// // @SecurityDefinitions.apikey BearerAuth
// // @Description GetMyProfile
// // @in header
// // @name Authorization
// func New(option Option) *gin.Engine {
// 	router := gin.New()

// 	router.Use(gin.Logger())
// 	router.Use(gin.Recovery())
// 	// jwtHandler := auth.JwtHandler{
// 	// 	SigninKey: option.Conf.SigninKey,
// 	// 	Log:       option.Logger,
// 	// }
// 	handlerV1 := v1.New(&v1.HandlerV1Config{
// 		ServiceManager: option.ServiceManager,
// 		Logger:         option.Logger,
// 		Cfg:            option.Conf,
// 	})
// 	api := router.Group("v1")
// 	api.POST("/users/", handlerV1.CreateUser)
// 	url := ginSwagger.URL("swagger/doc.json")
// 	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

// 	return router
// }

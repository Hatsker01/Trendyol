package api

import (
	_ "github.com/Trendyol/Api/api/docs"
	v1 "github.com/Trendyol/Api/api/handlers/v1"
	"github.com/Trendyol/Api/config"
	"github.com/Trendyol/Api/pkg/logger"
	"github.com/Trendyol/Api/services"
	"github.com/Trendyol/Api/storage/repo"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	RedisRepo      repo.RepositoryStorage
}

// @securityDefinitions.apikey BearerAuth
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
	// router.Use(casbin.NewJwtRoleStruct(option.Conf, jwtHandler))

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
		Redis:          option.RedisRepo,
	})

	api := router.Group("/v1")
	api.POST("/users", handlerV1.CreateUser)
	api.POST("/users/RegisterUser", handlerV1.RegisterUser)
	api.POST("users/register/user/:email/:coded", handlerV1.Verify)
	api.PUT("/users", handlerV1.UpdateUser)
	api.GET("/users/login/user", handlerV1.Login)
	api.GET("/user/getUserbyId/:id",handlerV1.GetUserById)
	api.GET("/users/getAll",handlerV1.GetAllUsers)
	api.DELETE("/user/delete/:id",handlerV1.DeleteUserById)
	api.PUT("/user/changePass",handlerV1.ChengePass)

	api.POST("/post",handlerV1.CreatePost)
	api.PUT("/post",handlerV1.UpdatePost)
	api.GET("/post/:id",handlerV1.GetPostById)
	api.GET("/post/getAllUserPosts/:id",handlerV1.GetAllUserPosts)
	api.GET("/posts",handlerV1.GetAllPosts)
	api.GET("/post/delete/:id",handlerV1.DeletePostbyId)
	api.GET("/posts/stars",handlerV1.SortByStars)

	api.GET("/category/:id",handlerV1.GetPostByCategory)
	api.POST("/category",handlerV1.CreateCategory)
	api.GET("/category/getAll",handlerV1.GetAllCategories)
	api.DELETE("/category/:id",handlerV1.DeleteCategory)
	api.GET("/category/getById/:id",handlerV1.GetCategory)

	api.POST("/like",handlerV1.PutLike)
	api.DELETE("/like/takeLike/:id",handlerV1.TakeLike)
	api.GET("/like/getAllLikeuser/:id",handlerV1.GetAllPostLikesUser)
	api.GET("/like/:id",handlerV1.GetLikeInfo)
	api.GET("/like/getPostLike/:id",handlerV1.GetPostLike)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

package app

import (
	"github.com/Trendyol/api/internal/service"
	"github.com/Trendyol/api/internal/transport/rest"
	"github.com/Trendyol/api/pkg/logger"
	"github.com/Trendyol/api/config"
)

// Run
// @title                      Golang CRM Swagger Documentation
// @version                    1.0
// @description                This is a sample server CRM server.
// @contact.name               API Support
// @SecurityDefinitions.apikey BearerAuth
// @Description                GetMyProfile
// @in                         header
// @name                       Authorization
func Run() {

	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "trendyol-service")

	// db, err := postsql.NewClient(cfg)
	// if err != nil {
	// 	log.Error("postgresql connection error", logger.Error(err))
	// 	return
	// }

	// repo := postgresql.NewRepositories(db)

	services,err := service.NewServiceManager(cfg)
	if err!=nil{
		log.Error("error while connecting services ",logger.Error(err))
		return
	}

	handlers := rest.NewHandler(&services, log)

	srv := handlers.Init(&cfg)

	err = srv.Run(cfg.HTTPHost + ":" + cfg.HTTPPort)
	if err != nil {
		log.Error("router running error", logger.Error(err))
		return
	}

}

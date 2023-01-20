package app

import (
	"net"

	"github.com/Trendyol/post_service/configs"
	pb "github.com/Trendyol/post_service/genproto"
	"github.com/Trendyol/post_service/internal/service"
	"github.com/Trendyol/post_service/internal/storage/postgres"
	postsql "github.com/Trendyol/post_service/pkg/database/postgresql"
	"github.com/Trendyol/post_service/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run() {

	cfg := configs.Load()
	log := logger.New(cfg.LogLevel, "trendyol-go-posts-service")
	db, err := postsql.NewClient(cfg)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))
	if err != nil {
		log.Error("postgres connection error", logger.Error(err))
		return
	}
	repo := postgres.NewRepasitories(db, log)
	postsService := service.NewPostsService(repo.Posts, log)
	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, postsService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

}

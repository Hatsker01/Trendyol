package app

import (
	"net"

	"github.com/Trendyol/service/configs"
	pb "github.com/Trendyol/service/genproto"
	"github.com/Trendyol/service/internal/service"
	"github.com/Trendyol/service/internal/storage/postgres"
	grpcClient "github.com/Trendyol/service/internal/transport/grpc_client"
	postsql "github.com/Trendyol/service/pkg/database/postgresql"
	"github.com/Trendyol/service/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run() {
	cfg := configs.Load()
	log := logger.New(cfg.LogLevel, "trendyol-go-service")
	db, err := postsql.NewClient(cfg)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))
	if err != nil {
		log.Error("postgres connection error", logger.Error(err))
		return
	}
	grpcC, err := grpcClient.New(cfg)
	if err != nil {
		log.Error("error establishing grpc connection", logger.Error(err))
		return
	}

	repo := postgres.NewRepasitories(db, log)
	usersService := service.NewUsersService(repo.Users, log, grpcC)
	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, usersService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

}

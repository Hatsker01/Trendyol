package service

import (
	"fmt"

	"github.com/Trendyol/api/config"
	pb "github.com/Trendyol/api/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IUserServiceManager interface {
	UserService() pb.UserServiceClient
}

type serviceManager struct {
	userService pb.UserServiceClient
}

func (s *serviceManager) UserService() pb.UserServiceClient {
	return s.userService
}

func NewServiceManager(conf config.Config) (IUserServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	ConnUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	serviceManager := &serviceManager{
		userService: pb.NewUserServiceClient(ConnUser),
	}
	return serviceManager, nil
}

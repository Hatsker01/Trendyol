package services

import (
	"fmt"

	"github.com/Trendyol/Api/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	pb "github.com/Trendyol/Api/genproto"
)

type IServiceManager interface {
	UserService() pb.UserServiceClient
	PostService() pb.PostServiceClient
}

type serviceManager struct {
	userService pb.UserServiceClient
	postService pb.PostServiceClient
}

func (s *serviceManager) UserService() pb.UserServiceClient {
	return s.userService
}
func (s *serviceManager) PostService()pb.PostServiceClient{
	return s.postService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	connPost,err:=grpc.Dial(
		fmt.Sprintf("%s:%d",conf.PostServiceHost, conf.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		return nil,err
	}

	serviceManager := &serviceManager{
		userService: pb.NewUserServiceClient(connUser),
		postService: pb.NewPostServiceClient(connPost),
	}
	return serviceManager, nil

}

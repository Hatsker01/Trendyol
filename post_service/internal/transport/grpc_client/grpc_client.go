// package grpcclient
package grpcclient

// import (
// 	"fmt"

// 	"github.com/Trendyol/post_service/configs"
// 	pb "github.com/Trendyol/post_service/genproto"
// 	"google.golang.org/grpc"
// )

// // type GrpcClientI interface {
// // 	PostService() pb.PostServiceClient
// // }
// type GrpcClientI interface {
// 	UserServise() pb.UserServiceClient
// }

// //GrpcClient ...
// type GrpcClient struct {
// 	cfg         configs.Config
// 	connections map[string]interface{}
// }

// //func New

// func New(cfg configs.Config) (*GrpcClient, error) {
// 	connuser, err := grpc.Dial(
// 		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
// 		grpc.WithInsecure(),
// 	)

// 	if err != nil {
// 		return nil, fmt.Errorf("user service dial host: %s port: %d err: %s",
// 			cfg.UserServiceHost, cfg.UserServicePort, err.Error())
// 	}

// 	return &GrpcClient{
// 		cfg: cfg,
// 		connections: map[string]interface{}{
// 			"user_service": pb.NewUserServiceClient(connuser),
// 		},
// 	}, nil
// }
// func (g *GrpcClient) UserServise() pb.UserServiceClient {
// 	return g.connections["service"].(pb.UserServiceClient)
// }

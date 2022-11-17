package grpcclient

import (
	"fmt"

	"github.com/Trendyol/service/configs"
	pb "github.com/Trendyol/service/genproto"
	"google.golang.org/grpc"
)

//GrpcClientI ...
type GrpcClientI interface {
	PostService() pb.PostServiceClient
}

//GrpcClient ...
type GrpcClient struct {
	cfg         configs.Config
	connections map[string]interface{}
}

//New ...
func New(cfg configs.Config) (*GrpcClient, error) {
	connpost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PostServiceHost, cfg.PostServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("post service dial host: %s port:%d err:%s",
			cfg.PostServiceHost, cfg.PostServicePort, err.Error())
	}
	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"post_service": pb.NewPostServiceClient(connpost),
		},
	}, nil
}

func (g *GrpcClient) PostService() pb.PostServiceClient {
	return g.connections["post_service"].(pb.PostServiceClient)
}

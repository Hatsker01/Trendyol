package service

import (
	"context"

	pb "github.com/Trendyol/post_service/genproto"
	"github.com/Trendyol/post_service/pkg/logger"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *PostsService) PutLike(ctx context.Context, like *pb.Like) (*pb.Like, error) {
	id, err := uuid.NewV4()
	if err != nil {
		s.logger.Error("failed while generating uuid for like", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while generating uuid")
	}
	like.Id = id.String()
	newLike, err := s.repo.PutLike(like)
	if err != nil {
		s.logger.Error("failed while putting like", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while putting like")
	}
	return newLike, nil
}

func (s *PostsService) TakeLike(ctx context.Context,id *pb.WithId)(*pb.Like,error){
	like,err:=s.repo.TakeLike(id.Id)
	if err!=nil{
		s.logger.Error("failed while taking like",logger.Error(err))
		return nil,status.Error(codes.Internal,"failed while taking like")
	}
	return like,err
}

func (s *PostsService) GetAllPostLikesUser(ctx context.Context,id *pb.WithId)(*pb.Likes,error){
	likes,err:=s.repo.GetAllLikesUser(id.Id)
	if err!=nil{
		s.logger.Error("failed while getting all likes user",logger.Error(err))
		return nil, status.Error(codes.Internal,"failed while getting all likes user")
	}
	return &pb.Likes{Likes: likes},nil
}

func (s *PostsService) GetLikeInfo(ctx context.Context,id *pb.LikeId)(*pb.Like,error){
	like,err:=s.repo.GetLikeInfo(id.Id)
	if err!=nil{
		s.logger.Error("failed while getting like info",logger.Error(err))
		return nil,status.Error(codes.Internal,"failed while getting like info")
	}
	return like,nil
}

func (s *PostsService) GetPostLike(ctx context.Context,id *pb.WithId)(*pb.Likes,error){
	likes,err:=s.repo.GetPostLike(id.Id)
	if err!=nil{
		 s.logger.Error("failed while getting post likes",logger.Error(err))
		 return nil,status.Error(codes.Internal,"failed while getting post likes")
	}
	return &pb.Likes{Likes: likes},nil
}
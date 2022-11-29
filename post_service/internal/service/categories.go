package service

import (
	"context"

	pb "github.com/Trendyol/post_service/genproto"
	"github.com/Trendyol/post_service/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *PostsService) GetPostByCategory(ctx context.Context, id *pb.CatID) (*pb.Posts, error) {
	posts, err := s.repo.GetPostByCategory(id.Id)
	if err != nil {
		s.logger.Error("failed while getting posts by category", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting posts by category id")
	}
	return &pb.Posts{Posts: posts}, nil
}

func (s *PostsService) GetAllCategories(cxt context.Context, emp *pb.Empty) (*pb.Categories, error) {
	categories, err := s.repo.GetAllCategories()
	if err != nil {
		s.logger.Error("failed while getting all categories", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting all categories")
	}
	return &pb.Categories{Categories: categories}, nil
}

func (s *PostsService) CreateCategory(ctx context.Context, cateReq *pb.CategoryReq) (*pb.Category, error) {
	category, err := s.repo.CreateCategory(cateReq)
	if err != nil {
		s.logger.Error("failed while creating category", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while creating category")
	}
	return category, nil
}

func (s *PostsService) DeleteCategory(ctx context.Context, id *pb.CatID) (*pb.Category, error) {
	category, err := s.repo.DeleteCategory(id.Id)
	if err != nil {
		s.logger.Error("failed while deleting category", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while deleting category")
	}
	return category, nil
}

func(s *PostsService) GetCategory(ctx context.Context,id *pb.CatID)(*pb.Category,error){
	category,err:=s.repo.GetCategory(id.Id)
	if err!=nil{
		s.logger.Error("failed while getting category by id",logger.Error(err))
		return nil,status.Error(codes.Internal,"failed while deleting category by id")
	}
	return category,nil
}

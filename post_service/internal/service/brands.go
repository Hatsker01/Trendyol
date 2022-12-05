package service

import (
	"context"

	pb "github.com/Trendyol/post_service/genproto"
	"github.com/Trendyol/post_service/pkg/logger"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *PostsService) CreateBrand(ctx context.Context, postBrand *pb.CreateBrandReq) (*pb.Brand, error) {
	id, err := uuid.NewV4()
	if err != nil {
		s.logger.Error("failed while generating uuid for brand", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while generating uuid for creating brand")
	}
	postBrand.Id = id.String()
	Brand, err := s.repo.CreateBrand(postBrand)
	if err != nil {
		s.logger.Error("failed while creating new brand", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while creating new brand")
	}
	return Brand, nil
}

func (s *PostsService) GetAllBrands(ctx context.Context,empty *pb.Empty)(*pb.Brands,error){
	brands,err:=s.repo.GetAllBrands()
	if err!=nil{
		s.logger.Error("failed while getting all brands",logger.Error(err))
		return nil,status.Error(codes.Internal,"failed while getting all brands")
	}
	return &pb.Brands{Brands: brands},nil
}

func (s *PostsService) DeleteBrand(ctx context.Context,id *pb.WithId)(*pb.Brand,error){
	brand,err:=s.repo.DeleteBrand(id.Id)
	if err!=nil{
		s.logger.Error("failed while deleting brand",logger.Error(err))
		return nil,status.Error(codes.Internal,"failed while deleting brand")
	}
	return brand,nil
}

func (s *PostsService) GetPostByBrand(ctx context.Context,id *pb.WithId)(*pb.Posts,error){
	posts,err:=s.repo.GetPostByBrand(id.Id)
	if err!=nil{
		s.logger.Error("failed while getting posts by brand",logger.Error(err))
		return nil,status.Error(codes.Internal,"failed while getting posts by brand")
	}
	return &pb.Posts{Posts: posts},nil
}

func (s *PostsService) GetBrandById(ctx context.Context,id *pb.WithId)(*pb.Brand,error){
	brand,err:=s.repo.GetBrandById(id.Id)
	if err!=nil{
		s.logger.Error("failed while getting brand by id",logger.Error(err))
		return nil,status.Error(codes.Internal,"failed while getting brand by id")
	}
	return brand,nil	
}


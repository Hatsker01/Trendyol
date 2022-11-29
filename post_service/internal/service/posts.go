package service

import (
	"context"

	pb "github.com/Trendyol/post_service/genproto"
	"github.com/Trendyol/post_service/pkg/logger"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Trendyol/post_service/internal/storage/postgres"
)

type PostsService struct {
	repo   postgres.Posts
	logger logger.Logger
}

func NewPostsService(repo postgres.Posts, logger logger.Logger) *PostsService {
	return &PostsService{
		repo:   repo,
		logger: logger,
	}
}

func (s *PostsService) CreatePost(ctx context.Context, post *pb.Post) (*pb.Post, error) {
	id, err := uuid.NewV4()
	if err != nil {
		s.logger.Error("failed while generating uuid", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while generating uuid")
	}
	post.Id = id.String()
	newPost, err := s.repo.CreatePost(post)
	if err != nil {
		s.logger.Error("failed while creating post", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while creating post")
	}
	return newPost, nil
}

func (s *PostsService) UpdatePost(ctx context.Context, upPost *pb.Post) (*pb.Post, error) {
	post, err := s.repo.UpdatePost(upPost)
	if err != nil {
		s.logger.Error("failed while updating post", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while updating post")
	}
	return post, nil
}

func (s *PostsService) GetPostById(ctx context.Context, id *pb.WithId) (*pb.Post, error) {
	post, err := s.repo.GetPostById(id.Id)
	if err != nil {
		s.logger.Error("failed while getting post by id", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting post by id")
	}
	return post, nil
}

func (s *PostsService) GetAllPosts(ctx context.Context, empty *pb.Empty) (*pb.Posts, error) {
	posts, err := s.repo.GetAllPosts()
	if err != nil {
		s.logger.Error("failed while getting all posts", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting all users")
	}
	return &pb.Posts{
		Posts: posts,
	}, nil
}

func (s *PostsService) DeletePostById(ctx context.Context, id *pb.WithId) (*pb.Post, error) {
	post, err := s.repo.DeletePostById(id.Id)
	if err != nil {
		s.logger.Error("failed while deleting post by id", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while deleting post by id")
	}
	return post, nil
}
func (s *PostsService) GetAllUserPosts(ctx context.Context, id *pb.WithId) (*pb.Posts, error) {
	posts, err := s.repo.GetAllUserPosts(id.Id)
	if err != nil {
		s.logger.Error("failed while getting user posts", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting user posts")
	}
	return &pb.Posts{
		Posts: posts,
	}, nil
}

func (s *PostsService) DeleteAllUserPosts(ctx context.Context, id *pb.WithId) (*pb.Posts, error) {
	posts, err := s.repo.DeleteAllUserPosts(id.Id)
	if err != nil {
		s.logger.Error("failed while deleting user posts postservice", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while deleteting user posts")
	}
	return &pb.Posts{
		Posts: posts,
	}, nil
}
func (s *PostsService) StarPosts(ctx context.Context,empty *pb.Empty)(*pb.Posts,error){
	posts,err:=s.repo.StarPosts()
	if err!=nil{
		s.logger.Error("failed while getting posts seperating by stars",logger.Error(err))
		return nil,status.Error(codes.Internal,"failed while getting post separating by stars")
	}
	return &pb.Posts{
		Posts: posts,
	},nil
}

func (s *PostsService) GetPostsByPrice(ctx context.Context,priceSep *pb.PriceSep)(*pb.Posts,error){
	posts,err:=s.repo.SeperatePostByPrice(priceSep)
	if err!=nil{
		s.logger.Error("failed while getting posts seperate by Price",logger.Error(err))
		return nil,status.Error(codes.Internal,"failed while getting posts seperating by price")
	}
	return &pb.Posts{
		Posts: posts,
	},nil
}

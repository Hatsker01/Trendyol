package service

import (
	"context"
	"fmt"

	pb "github.com/Trendyol/service/genproto"
	cl "github.com/Trendyol/service/internal/transport/grpc_client"
	"github.com/Trendyol/service/pkg/logger"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Trendyol/service/internal/storage/postgres"
)

type UsersService struct {
	repo   postgres.Users
	logger logger.Logger
	client cl.GrpcClientI
}

func NewUsersService(repo postgres.Users, logger logger.Logger, client cl.GrpcClientI) *UsersService {
	return &UsersService{
		repo:   repo,
		logger: logger,
		client: client,
	}
}
func (s *UsersService) CheckField(ctx context.Context, req *pb.CheckFieldRequest) (*pb.CheckFieldReponse, error) {
	check, err := s.repo.CheckField(req.Field, req.Value)
	if err != nil {
		s.logger.Error("failed while getting user", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting user")

	}
	return &pb.CheckFieldReponse{
		Check: check,
	}, nil
}

func (s *UsersService) CreateUser(ctx context.Context, user *pb.CreateUserReq) (*pb.User, error) {
	id, err := uuid.NewV4()
	if err != nil {
		s.logger.Error("failed while generating uuid", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while generating uuid")
	}
	user.Id = id.String()
	valEmail, err := s.repo.EmailValid(user.Email)
	if err != nil {
		s.logger.Error("failed while validing email", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while validating email")
	}
	if !valEmail {
		s.logger.Error("Email is in use", logger.Error(err))
		return nil, status.Error(codes.Internal, "Email is already in use")
	}
	newUser, err := s.repo.CreateUser(user)
	if err != nil {
		s.logger.Error("failed while creating user", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while creating user")
	}
	return newUser, nil
}

func (s *UsersService) UpdateUser(ctx context.Context, upuser *pb.User) (*pb.User, error) {
	user, err := s.repo.UpdateUser(upuser)
	if err != nil {
		s.logger.Error("failed while updating user", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while updating user")
	}
	return user, nil
}

func (s *UsersService) GetUserById(ctx context.Context, id *pb.WithId) (*pb.User, error) {
	user, err := s.repo.GetUserById(id.Id)
	if err != nil {
		s.logger.Error("failed while getting user by id", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting user by id")
	}
	// post,err:=s.client.PostService().GetPostById(ctx,id)
	// if err!=nil{
	// 	s.logger.Error("failed while getting post grc by id", logger.Error(err))
	// 	return nil, nil,status.Error(codes.Internal, "failed while getting post grpc by id")
	// }

	post, err := s.client.PostService().GetAllUserPosts(ctx, &pb.WithId{Id: id.Id})
	if err != nil {
		s.logger.Error("failed while getting user posts")
		return nil, status.Error(codes.Internal, "failed while getting user posts")
	}
	fmt.Println(id.Id)
	user.Posts = post.Posts
	return user, nil
}

func (s *UsersService) GetAllUsers(ctx context.Context, empty *pb.Empty) (*pb.Users, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		s.logger.Error("failed while getting all users", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting all users")
	}
	for _, user := range users {
		posts, err := s.client.PostService().GetAllUserPosts(ctx, &pb.WithId{Id: user.Id})
		if err != nil {
			s.logger.Error("failed while getting all user posts", logger.Error(err))
			return nil, status.Error(codes.Internal, "failed while getting all user posts")
		}
		user.Posts = posts.Posts
	}
	return &pb.Users{
		User: users,
	}, nil
}

func (s *UsersService) DeleteUserById(ctx context.Context, id *pb.WithId) (*pb.User, error) {
	user, err := s.repo.DeleteUserById(id.Id)
	if err != nil {
		s.logger.Error("failed while deleting user by id", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while deleting user by id")
	}
	post, err := s.client.PostService().DeleteAllUserPosts(ctx, &pb.WithId{Id: id.Id})
	if err != nil {
		s.logger.Error("failed while deleting user posts", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed while deleting user posts")

	}
	user.Posts = post.Posts
	return user, nil
}

func (s *UsersService) LoginUser(ctx context.Context, login *pb.LoginUserReq) (*pb.User, error) {
	user, err := s.repo.LoginUser(login)
	if err != nil {
		s.logger.Error("email or password wrong", logger.Error(err))
		return nil, status.Error(codes.Internal, "email ot password wrong")
	}
	return user, nil
}

func (s *UsersService) ChangePassword(ctx context.Context, newPass *pb.ChangePassReq)(*pb.ChangePassRes,error){
	pass,err:=s.repo.ChangePassword(newPass)
	if err!=nil{
		s.logger.Error("failed while changing password",logger.Error(err))
		return nil,status.Error(codes.Internal,"failed while changing password")
	}
	return pass,err
}

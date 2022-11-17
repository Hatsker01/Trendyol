package postgres

import (
	pb "github.com/Trendyol/service/genproto"
	"github.com/Trendyol/service/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type Users interface {
	CreateUser(user *pb.CreateUserReq) (*pb.User, error)
	GetUserById(id string) (*pb.User, error)
	UpdateUser(user *pb.User) (*pb.User, error)
	GetAllUsers() ([]*pb.User, error)
	DeleteUserById(id string) (*pb.User, error)
	LoginUser(login *pb.LoginUserReq) (*pb.User, error)
	EmailValid(email string) (bool, error)
}

type Repasitories struct {
	Users  Users
	logger logger.Logger
}

func NewRepasitories(db *sqlx.DB, log logger.Logger) *Repasitories {
	return &Repasitories{
		Users:  NewUserRepo(db),
		logger: log,
	}
}

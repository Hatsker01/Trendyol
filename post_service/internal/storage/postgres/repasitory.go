package postgres

import (
	pb "github.com/Trendyol/post_service/genproto"
	"github.com/Trendyol/post_service/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type Posts interface {
	CreatePost(post *pb.Post) (*pb.Post, error)
	GetPostById(id string) (*pb.Post, error)
	UpdatePost(post *pb.Post) (*pb.Post, error)
	GetAllPosts() ([]*pb.Post, error)
	DeletePostById(id string) (*pb.Post, error)
	GetAllUserPosts(id string) ([]*pb.Post, error)
	DeleteAllUserPosts(id string) ([]*pb.Post, error)
	StarPosts()([]*pb.Post,error)
	SeperatePostByPrice(priceSep *pb.PriceSep)([]*pb.Post,error)
	GetingPostsByColor(color *pb.ColorReq)([]*pb.Post,error)

	GetPostByPrice(price *pb.GetPostPriceReq)([]*pb.Post,error)

	GetPostByCategory(id string)([]*pb.Post,error)
	GetAllCategories()([]*pb.Category,error)
	CreateCategory(category *pb.CategoryReq)(*pb.Category,error)
	DeleteCategory(id string)(*pb.Category,error)
	GetCategory(id string)(*pb.Category,error)

	PutLike(like *pb.Like)(*pb.Like,error)
	TakeLike(id string)(*pb.Like,error)
	GetAllLikesUser(id string)([]*pb.Like,error)
	GetLikeInfo(id string)(*pb.Like,error)
	GetPostLike(id string)([]*pb.Like,error)

	PutStar(star *pb.StarReq) (*pb.Stars, error)
	GetStar(id string) (*pb.Stars, error)
	TakeStar(id string)(*pb.Empty,error)

	CreateBrand(brand *pb.CreateBrandReq)(*pb.Brand,error)
	GetAllBrands() ([]*pb.Brand, error)
	DeleteBrand(id string) (*pb.Brand, error)
	GetPostByBrand(id string) ([]*pb.Post, error)
	GetBrandById(id string) (*pb.Brand, error)
	

}



type Repasitories struct {
	Posts  Posts
	logger logger.Logger
}

func NewRepasitories(db *sqlx.DB, log logger.Logger) *Repasitories {
	return &Repasitories{
		Posts:  NewPostRepo(db),
		logger: log,
	}
}

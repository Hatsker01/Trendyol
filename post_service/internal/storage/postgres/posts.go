package postgres

import (
	"database/sql"
	"strconv"
	"time"

	pb "github.com/Trendyol/post_service/genproto"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type PostsRepo struct {
	db *sqlx.DB
}

func (r *PostsRepo) CreatePost(post *pb.Post) (*pb.Post, error) {

	newPost := pb.Post{}
	query := `INSERT INTO posts(id,title,description,body,author_id,rating,price,product_type,size,color,gen,brand_id,category_id,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14) 
	RETURNING id,title,description,body,author_id,stars,rating,price,product_type,size,color,gen,brand_id,category_id,created_at`
	err := r.db.QueryRow(query, post.Id, post.Title, post.Description, post.Body, post.AuthorId, post.Rating, post.Price, post.ProductType, pq.Array(post.Size_), post.Color, post.Gen, post.BrandId, post.CategoryId, time.Now().UTC()).Scan(
		&newPost.Id,
		&newPost.Title,
		&newPost.Description,
		&newPost.Body,
		&newPost.AuthorId,
		&newPost.Stars,
		&newPost.Rating,
		&newPost.Price,
		&newPost.ProductType,
		pq.Array(&newPost.Size_),
		&newPost.Color,
		&newPost.Gen,
		&newPost.BrandId,
		&newPost.CategoryId,
		&newPost.CreatedAt,
	)

	if err != nil {
		return &pb.Post{}, err
	}
	return &newPost, nil

}

func (r *PostsRepo) GetAllUserPosts(id string) ([]*pb.Post, error) {
	var posts []*pb.Post
	query := `SELECT id,title,description,body,author_id,stars,rating,price,product_type,size,color,gen,brand_id,category_id,created_at,updated_at from posts where deleted_at is null and author_id =$1`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		post := pb.Post{}
		var (
			stars      sql.NullString
			rating     sql.NullString
			updated_at sql.NullTime
		)
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Description,
			&post.Body,
			&post.AuthorId,
			&stars,
			&rating,
			&post.Price,
			&post.ProductType,
			pq.Array(&post.Size_),
			&post.Color,
			&post.Gen,
			&post.BrandId,
			&post.CategoryId,
			&post.CreatedAt,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			post.UpdatedAt = updated_at.Time.String()
		}
		if stars.Valid {
			post.Stars = stars.String
		}
		if rating.Valid {
			post.Rating = rating.String
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (r *PostsRepo) UpdatePost(post *pb.Post) (*pb.Post, error) {
	query := `UPDATE posts SET title=$1,description=$2,body=$3,author_id=$4,rating=$5,price=$6,product_type=$7,size=$8,color=9,gen=$10,brand_id=$11,category_id=$12,updated_at=$13 where id=$14
	RETURNING id,title,description,body,author_id,stars,rating,price,product_type,size,color,gen,brand_id,category_id,created_at`
	upPost := pb.Post{}
	var (
		stars  sql.NullString
		rating sql.NullString
	)
	err := r.db.QueryRow(query, post.Title, post.Description, post.Body, post.AuthorId, post.Rating, post.Price, post.ProductType, pq.Array(post.Size_), post.Color, post.Gen, post.BrandId, post.CategoryId, time.Now().UTC(), post.Id).Scan(
		&upPost.Id,
		&upPost.Title,
		&upPost.Description,
		&upPost.Body,
		&upPost.AuthorId,
		&stars,
		&rating,
		&upPost.Price,
		&upPost.ProductType,
		pq.Array(&upPost.Size_),
		&post.Color,
		&upPost.Gen,
		&upPost.BrandId,
		&upPost.CategoryId,
		&upPost.CreatedAt,
	)
	upPost.UpdatedAt = time.Now().UTC().String()
	if err != nil {
		return nil, err
	}
	if rating.Valid {
		upPost.Rating = rating.String
	}
	if stars.Valid {
		upPost.Stars = stars.String
	}
	return &upPost, nil
}

func (r *PostsRepo) GetPostById(id string) (*pb.Post, error) {
	query := `SELECT id,title,description,body,author_id,stars,rating,price,product_type,size,color,gen,brand_id,category_id,created_at,updated_at from posts where id=$1 and deleted_at is null`
	post := pb.Post{}
	var (
		stars      sql.NullString
		rating     sql.NullString
		updated_at sql.NullTime
	)
	err := r.db.QueryRow(query, id).Scan(
		&post.Id,
		&post.Title,
		&post.Description,
		&post.Body,
		&post.AuthorId,
		&stars,
		&rating,
		&post.Price,
		&post.ProductType,
		pq.Array(&post.Size_),
		&post.Color,
		&post.Gen,
		&post.BrandId,
		&post.CategoryId,
		&post.CreatedAt,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		post.UpdatedAt = updated_at.Time.String()
	}
	if stars.Valid {
		post.Stars = stars.String
	}
	if rating.Valid {
		post.Rating = rating.String
	}
	return &post, nil
}

func (r *PostsRepo) GetAllPosts() ([]*pb.Post, error) {
	query := `SELECT id,title,description,body,author_id,stars,rating,price,product_type,size,color,gen,brand_id,category_id,created_at,updated_at from posts WHERE deleted_at is null`
	var posts []*pb.Post

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		post := pb.Post{}
		var (
			stars      sql.NullString
			rating     sql.NullString
			updated_at sql.NullTime
		)
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Description,
			&post.Body,
			&post.AuthorId,
			&stars,
			&rating,
			&post.Price,
			&post.ProductType,
			pq.Array(&post.Size_),
			&post.Color,
			&post.Gen,
			&post.BrandId,
			&post.CategoryId,
			&post.CreatedAt,
			&updated_at,
		)

		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			post.UpdatedAt = updated_at.Time.String()
		}
		if rating.Valid {
			post.Rating = rating.String
		}
		if stars.Valid {
			post.Stars = stars.String
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (r *PostsRepo) DeletePostById(id string) (*pb.Post, error) {
	post, err := r.GetPostById(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE posts SET deleted_at=$2 where id=$1`
	_, err = r.db.Exec(query, id, time.Now().UTC())
	if err != nil {
		return nil, err
	}
	return post, nil

}

func (r *PostsRepo) DeleteAllUserPosts(id string) ([]*pb.Post, error) {
	posts, err := r.GetAllUserPosts(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE posts SET deleted_at=$1 where author_id=$2`
	_, err = r.db.Exec(query, time.Now().UTC(), id)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostsRepo) StarPosts() ([]*pb.Post, error) {
	query := `SELECT id,title,description,body,author_id,stars,rating,price,product_type,size,color,gen,brand_id,category_id from posts where deleted_at in null order by stars desc`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	var posts []*pb.Post
	for rows.Next() {
		var post pb.Post
		var (
			stars  sql.NullString
			rating sql.NullString
		)
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Description,
			&post.Body,
			&post.AuthorId,
			&stars,
			&rating,
			&post.ProductType,
			pq.Array(&post.Size_),
			&post.Color,
			&post.Gen,
			&post.BrandId,
			&post.CategoryId,
		)
		if err != nil {
			return nil, err
		}
		if rating.Valid {
			post.Rating = rating.String
		}
		if stars.Valid {
			post.Stars = stars.String
		}
		posts = append(posts, &post)
	}
	return posts, nil
}
func (r *PostsRepo) SeperatePostByPrice(priceSep *pb.PriceSep) ([]*pb.Post, error) {
	query := `SELECT id,title,description,body,author_id,stars,rating,price,product_type,size,color,gen,brand_id,category_id from posts where deleted_at is null order by price `
	if priceSep.High {
		query += "desc"
	}
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	var posts []*pb.Post
	for rows.Next() {
		var post pb.Post
		var (
			stars  sql.NullString
			rating sql.NullString
		)
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Description,
			&post.Body,
			&post.AuthorId,
			&stars,
			&rating,
			&post.Price,
			&post.ProductType,
			pq.Array(&post.Size_),
			&post.Color,
			&post.Gen,
		)
		if err != nil {
			return nil, err
		}
		if rating.Valid {
			post.Rating = rating.String
		}
		if stars.Valid {
			post.Stars = stars.String
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (r *PostsRepo) GetCountPosts() (int, error) {
	query := `select count(*) from posts where deleted_at is null`
	var count int
	err := r.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil

}

func (r *PostsRepo) GetPostByPrice(price *pb.GetPostPriceReq) ([]*pb.Post, error) {
	var low, high int
	var posts []*pb.Post
	var err error
	if price.High == "" {
		high, err = r.GetCountPosts()
		if err != nil {
			return nil, err
		}
	} else {
		high, _ = strconv.Atoi(price.High)
	}
	if price.Low == "" {
		low = 0
	} else {
		low, _ = strconv.Atoi(price.Low)
	}
	query := `SELECT id,title,description,body,author_id,stars,rating,price,product_type,size,color,gen,brand_id,category_id from posts where deleted_at is null and price>$1 and price<$2`
	rows, err := r.db.Query(query, low, high)
	for rows.Next() {
		var post pb.Post
		var (
			stars  sql.NullString
			rating sql.NullString
		)
		err := rows.Scan(
			&post.Id,
			&post.Description,
			&post.Body,
			&post.AuthorId,
			&stars,
			&rating,
			&post.Price,
			&post.ProductType,
			pq.Array(&post.Size_),
			&post.Color,
			&post.Gen,
			&post.BrandId,
			&post.CategoryId,
		)
		if err != nil {
			return nil, err
		}
		if rating.Valid {
			post.Rating = rating.String
		}
		if stars.Valid {
			post.Stars = stars.String
		}
		posts = append(posts, &post)
	}
	return posts, nil

}

func (r *PostsRepo) GetingPostsByColor(color *pb.ColorReq) ([]*pb.Post, error) {
	query := `SELECT id,title,description,body,author_id,stars,rating,price,product_type,size,color,gen,brand_id,category_id from posts where deleted_at is null and color=$1`
	var posts []*pb.Post
	rows, err := r.db.Query(query, color.Color)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var post pb.Post
		var (
			stars  sql.NullString
			rating sql.NullString
		)
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Description,
			&post.Body,
			&post.AuthorId,
			&stars,
			&rating,
			&post.Price,
			&post.ProductType,
			pq.Array(&post.Size_),
			&post.Color,
			&post.Gen,
			&post.BrandId,
			&post.CategoryId,
		)
		if err != nil {
			return nil, err
		}
		if rating.Valid {
			post.Rating = rating.String
		}
		if stars.Valid {
			post.Stars = stars.String
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (r *PostsRepo) PutStar(star *pb.StarReq) (*pb.Stars, error) {
	var id string

	query := `INSERT INTO stars(id,post_id,user_id,star,created_at) returning post_id`
	err := r.db.QueryRow(query, star.Id, star.PostId, star.UserId, time.Now().UTC()).Scan(&id)
	if err != nil {
		return nil, err
	}
	stars, err := r.GetStar(id)
	if err != nil {
		return nil, err
	}
	return stars, nil
}

func (r *PostsRepo) GetStar(id string) (*pb.Stars, error) {
	var star pb.Stars
	query := `SELECT ROUND(avg(star),2) as avarage from stars where post_id=$1 and deleted_at is null`
	err := r.db.QueryRow(query, id).Scan(
		&star.AvaregeStar,
	)
	if err != nil {
		return nil, err
	}
	star.PostId = id
	return &star, nil

}

func (r *PostsRepo) TakeStar(id string) (*pb.Empty, error) {
	query := `UPTATE stars SET deleted_at =$1 where id=$2`
	_, err := r.db.Exec(query, time.Now().UTC(), id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//functions for categories

func NewPostRepo(db *sqlx.DB) *PostsRepo {
	return &PostsRepo{
		db: db,
	}
}

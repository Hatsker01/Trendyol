package postgres

import (
	"database/sql"
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
	query := `INSERT INTO posts(id,title,description,body,author_id,stars,rating,price,product_type,size,color,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) 
	RETURNING id,title,description,body,author_id,stars,rating,price,product_type,size,created_at`
	err := r.db.QueryRow(query, post.Id, post.Title, post.Description, post.Body, post.AuthorId, post.Stars, post.Rating, post.Price, post.ProductType, pq.Array(post.Size_),post.Color, time.Now().UTC()).Scan(
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
		&newPost.CreatedAt,
	)

	if err != nil {
		return &pb.Post{}, err
	}
	return &newPost, nil

}

func (r *PostsRepo) GetAllUserPosts(id string) ([]*pb.Post, error) {
	var posts []*pb.Post
	query := `SELECT id,title,description,body,author_id,stars,rating,price,product_type,size,color,created_at,updated_at from posts where deleted_at is null and author_id =$1`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		post := pb.Post{}
		var updated_at sql.NullTime
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Description,
			&post.Body,
			&post.AuthorId,
			&post.Stars,
			&post.Rating,
			&post.Price,
			&post.ProductType,
			pq.Array(&post.Size_),
			&post.Color,
			&post.CreatedAt,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			post.UpdatedAt = updated_at.Time.String()
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (r *PostsRepo) UpdatePost(post *pb.Post) (*pb.Post, error) {
	query := `UPDATE posts SET title=$1,description=$2,body=$3,author_id=$4,stars=$5,rating=$6,price=$7,product_type=$8,size=$9,updated_at=$10 where id=$11
	RETURNING id,title,description,body,author_id,stars,rating,price,product_type,size,created_at`
	upPost := pb.Post{}
	err := r.db.QueryRow(query, post.Title, post.Description, post.Body, post.AuthorId, post.Stars, post.Rating, post.Price, post.ProductType, pq.Array(post.Size_), time.Now().UTC(), post.Id).Scan(
		&upPost.Id,
		&upPost.Title,
		&upPost.Description,
		&upPost.Body,
		&upPost.AuthorId,
		&upPost.Stars,
		&upPost.Rating,
		&upPost.Price,
		&upPost.ProductType,
		pq.Array(&upPost.Size_),
		&upPost.CreatedAt,
	)
	upPost.UpdatedAt = time.Now().UTC().String()
	if err != nil {
		return nil, err
	}
	return &upPost, nil
}

func (r *PostsRepo) GetPostById(id string) (*pb.Post, error) {
	query := `SELECT id,title,description,body,author_id,stars,rating,price,product_type,size,color,created_at,updated_at from posts where id=$1 and deleted_at is null`
	post := pb.Post{}
	var update_at sql.NullTime
	err := r.db.QueryRow(query, id).Scan(
		&post.Id,
		&post.Title,
		&post.Description,
		&post.Body,
		&post.AuthorId,
		&post.Stars,
		&post.Rating,
		&post.Price,
		&post.ProductType,
		pq.Array(&post.Size_),
		&post.Color,
		&post.CreatedAt,
		&update_at,
	)
	if err != nil {
		return nil, err
	}
	if update_at.Valid{
		post.UpdatedAt=update_at.Time.String()
	}
	return &post, nil
}

func (r *PostsRepo) GetAllPosts() ([]*pb.Post, error) {
	query := `SELECT id,title,description,body,author_id,stars,rating,price,product_type,size,color,created_at,updated_at from posts WHERE deleted_at is null`
	var posts []*pb.Post
	var update_at sql.NullTime
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		post := pb.Post{}
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Description,
			&post.Body,
			&post.AuthorId,
			&post.Stars,
			&post.Rating,
			&post.Price,
			&post.ProductType,
			pq.Array(&post.Size_),
			&post.Color,
			&post.CreatedAt,
			&update_at,
		)

		if err != nil {
			return nil, err
		}
		if update_at.Valid {
			post.UpdatedAt = update_at.Time.String()
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

func (r *PostsRepo) StarPosts()([]*pb.Post,error){
	query:=`SELECT id,title,description,body,author_id,stars,rating,price,product_type,size,color from posts where deleted_at in null order by stars desc`
	rows,err:=r.db.Query(query)
	if err!=nil{
		return nil,err
	}
	var posts []*pb.Post
	for rows.Next(){
		var post pb.Post
		err:=rows.Scan(
			&post.Id,
			&post.Title,
			&post.Description,
			&post.Body,
			&post.AuthorId,
			&post.Stars,
			&post.Rating,
			&post.ProductType,
			pq.Array(&post.Size_),
			&post.Color,
		)
		if err!=nil{
			return nil,err
		}
		posts = append(posts, &post)
	}
	return posts,nil
}
func (r *PostsRepo) SeperatePostByPrice(priceSep *pb.PriceSep)([]*pb.Post,error){
	query:=`SELECT id,title,description,body,author_id,stars,rating,price,product_type,size,color from posts where deleted_at is null order by price `
	if priceSep.High{
		query+="desc"
	}
	rows,err:=r.db.Query(query)
	if err!=nil{
		return nil,err
	}
	var posts []*pb.Post
	for rows.Next(){
		var post pb.Post
		err:=rows.Scan(
			&post.Id,
			&post.Title,
			&post.Description,
			&post.Body,
			&post.AuthorId,
			&post.Stars,
			&post.Stars,
			&post.Rating,
			&post.Price,
			&post.ProductType,
			pq.Array(&post.Size_),
			&post.Color,
		)
		if err!=nil{
			return nil,err
		}
		posts = append(posts, &post)
	}
	return posts,nil
}

//functions for categories



func NewPostRepo(db *sqlx.DB) *PostsRepo {
	return &PostsRepo{
		db: db,
	}
}

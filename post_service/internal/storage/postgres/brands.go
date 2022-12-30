package postgres

import (
	"database/sql"
	"time"

	pb "github.com/Trendyol/post_service/genproto"
	"github.com/lib/pq"
)

func (r *PostsRepo) CreateBrand(brand *pb.CreateBrandReq) (*pb.Brand, error) {
	query := `INSERT INTO brands(id,name,created_at) values($1,$2,$3) returning id,name,created_at`
	var newBrand pb.Brand
	err := r.db.QueryRow(query, brand.Id, brand.Name, time.Now().UTC()).Scan(
		&newBrand.Id,
		&newBrand.Name,
		&newBrand.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &newBrand, nil
}

func (r *PostsRepo) GetAllBrands() ([]*pb.Brand, error) {
	query := `SELECT id,name,created_at,updated_at from brands where deleted_at is null`
	var brands []*pb.Brand
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var brand pb.Brand
		var updated_at sql.NullTime
		err := rows.Scan(
			&brand.Id,
			&brand.Name,
			&brand.CreatedAt,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			brand.UpdatedAt = updated_at.Time.String()
		}
		brands = append(brands, &brand)
	}
	return brands, nil
}

func (r *PostsRepo) DeleteBrand(id string) (*pb.Brand, error) {
	brand, err := r.GetBrandById(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE brands SET deleted_at=$1 where id =$2`
	_, err = r.db.Exec(query, time.Now().UTC(), id)
	if err != nil {
		return nil, err
	}
	return brand, nil

}

func (r *PostsRepo) GetBrandById(id string) (*pb.Brand, error) {
	query := `SELECT id,name,created_at,updated_at from brands where id=$1 and deleted_at is null`
	var updated_at sql.NullTime
	var brand *pb.Brand
	err := r.db.QueryRow(query, id).Scan(
		&brand.Id,
		&brand.Name,
		&brand.CreatedAt,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		brand.UpdatedAt = updated_at.Time.String()
	}
	return brand, nil
}

func (r *PostsRepo) GetPostByBrand(id string) ([]*pb.Post, error) {
	query := `SELECT id,title,description,body,author_id,rating,price,product_type,size,color,gen,brand_id,category_id,created_at,updated_at where brand_id=$1 and deleted_at is null`
	var posts []*pb.Post
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var updated_at sql.NullTime
		var post pb.Post
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Description,
			&post.Body,
			&post.AuthorId,
			&post.Rating,
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
		posts = append(posts, &post)
	}
	return posts, nil
}

package postgres

import (
	"database/sql"
	"time"

	pb "github.com/Trendyol/post_service/genproto"
	"github.com/lib/pq"
)

func (r *PostsRepo)CreateCategory(category *pb.CategoryReq)(*pb.Category,error){
	var id string
	query:=`INSERT INTO category(name,created_at) VALUES ($1,$2) RETURNING category_id`
	err:=r.db.QueryRow(query,category.Name,time.Now().UTC()).Scan(&id)
	if err!=nil{
		return nil,err
	}
	newCategory,err:=r.GetCategory(id)
	if err!=nil{
		return nil,err
	}
	return newCategory,nil

}

func (r *PostsRepo) GetPostByCategory(id string)([]*pb.Post,error){
	var posts []*pb.Post
	query :=`select id,title,description,body,author_id,stars,rating,price,product_type,size,color,gen,brand_id,category_id from posts join category on posts.category_id=category.id where posts.deleted_at is null and category.deleted_at is null`
	rows,err:=r.db.Query(query,id)
	if err!=nil{
		return nil,err
	}
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
			&post.Price,
			&post.ProductType,
			pq.Array(&post.Size_),
			&post.Color,
			&post.Gen,
			&post.BrandId,
			&post.CategoryId,
		)
		if err!=nil{
			return nil,err
		}
		posts = append(posts, &post)
	}
	return posts,nil
}

func (r *PostsRepo) GetCategory(id string)(*pb.Category,error){
	category:= pb.Category{}
	var update_at sql.NullTime
	query:=`SELECT category_id,name,created_at,updated_at from category where deleted_at is null and category_id = $1`
	err:=r.db.QueryRow(query,id).Scan(
		&category.Id,
		&category.Name,
		&category.CreatedAt,
		&update_at,
	)
	if err!=nil{
		return nil,err
	}
	if update_at.Valid{
		category.UpdatedAt=update_at.Time.String()
	}
	return &category,nil
}

func (r *PostsRepo) GetAllCategories()([]*pb.Category,error){
	var categories []*pb.Category
	query :=`SELECT category_id,name,created_at,updated_at from category where deleted_at is null`
	rows,err:=r.db.Query(query)
	if err!=nil{
		return nil,err
	}
	for rows.Next(){
		var update_at sql.NullTime
		var category pb.Category
		err:=rows.Scan(
			&category.Id,
			&category.Name,
			&category.CreatedAt,
			&update_at,
		)
		if err!=nil{
			return nil,err
		}
		if update_at.Valid{
			category.UpdatedAt=update_at.Time.String()
		}
		categories = append(categories, &category)
	}
	return categories,nil
}

func (r *PostsRepo) DeleteCategory(id string)(*pb.Category,error){
	category,err:=r.GetCategory(id)
	if err!=nil{
		return nil,err
	}
	query :=`UPDATE category set deleted_at = $1 where category_id=$2`
	_,err=r.db.Exec(query,time.Now().UTC(),id)
	if err!=nil{
		return nil,err
	}
	return category,nil
}
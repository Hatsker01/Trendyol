package postgres

import (
	"time"

	pb "github.com/Trendyol/post_service/genproto"
)

func (r *PostsRepo)PutLike(like *pb.Like)(*pb.Like,error){
	var newlike *pb.Like
	query:=`INSERT INTO likes(id,user_id,post_id,created_at) VALUES($1,$2,$3,$4) RETURNING id`
	err:=r.db.QueryRow(query,like.Id,like.UserId,like.PostId,time.Now().UTC()).Scan(
		&newlike.Id,
	)
	if err!=nil{
		return nil,err
	}
	newlike,err=r.GetLikeInfo(newlike.Id)
	if err!=nil{
		return nil,err
	}
	return newlike,nil

}

func (r *PostsRepo)TakeLike(id string)(*pb.Like,error){
	query:=`UPDATE TABLE likes SET deleted_at = $1 WHERE id = $2`
	_,err:=r.db.Exec(query,time.Now().UTC(),id)
	if err!=nil{
		return nil,err
	}
	like,err:=r.GetLikeInfo(id)
	if err!=nil{
		return nil,err
	}
	return like,nil
}

func (r *PostsRepo)GetAllLikesUser(id string)([]*pb.Like,error){
	query:=`SELECT id,user_id,post_id,created_at from likes where deleted_at is null and user_id=$1`
	rows,err:=r.db.Query(query,id)
	var userLikes []*pb.Like
	if err!=nil{
		return nil,err
	}
	for rows.Next(){
		var like pb.Like
		err:=rows.Scan(
			&like.Id,
			&like.UserId,
			&like.PostId,
			&like.CreatedAt,
		)
		if err!=nil{
			return nil,err
		}
		userLikes = append(userLikes, &like)
	}
	return userLikes,nil

}

func (r *PostsRepo)GetPostLike(id string)([]*pb.Like,error){
	query:=	`SELECT id,user_id,post_id,created_at from likes where deleted_at is null and post_id=$1`
	var postLikes []*pb.Like
	rows,err:=r.db.Query(query,id)
	if err!=nil{
		return nil,err
	}
	for rows.Next(){
		var like pb.Like
		err:=rows.Scan(
			&like.Id,
			&like.UserId,
			&like.PostId,
			&like.CreatedAt,
		)
		if err!=nil{
			return nil,err
		}
		postLikes = append(postLikes, &like)
	}
	return postLikes,nil
}


func (r *PostsRepo)GetLikeInfo(id string)(*pb.Like,error){
	var like *pb.Like
	query:=`SELECT id,user_id,post_id,created_at from likes where deleted_at is null and id =$1`
	err:=r.db.QueryRow(query,id).Scan(
		&like.Id,
		&like.UserId,
		&like.PostId,
		&like.CreatedAt,
	)
	if err!=nil{
		return nil,err
	}
	return like,nil

}
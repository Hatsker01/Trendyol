package postgres

import (
	"database/sql"
	"time"

	pb "github.com/Trendyol/post_service/genproto"
)


func (r *PostsRepo)Product_sale(prod *pb.ProductSaleReq)(*pb.Productsale,error){
	newProsale:=pb.Productsale{}
	var saled_at sql.NullTime
	query:=`INSERT INTO product_sale (id,user_id,post_id,count,price,created_at) VALUES($1,$2,$3,$4,$5,$6) returning id,user_id,post_id,count,price,created_at`
	err:=r.db.QueryRow(query,prod.Id,&prod.UserId,&prod.PostId,&prod.Count,&prod.Price,time.Now().UTC()).Scan(
		&newProsale.Id,
		&newProsale.UserId,
		&newProsale.PostId,
		&newProsale.Count,
		&newProsale.Price,
		&saled_at,
		&newProsale.CreatedAt,
	)
	if err!=nil{
		return nil,err
	}
	return &newProsale,nil
}

func (r *PostsRepo) SaleProductDel(id string)(*pb.Productsale,error){
	product,err:=r.InfoProduct(id)
	if err!=nil{
		return nil,err
	}
	query:=`UPDATE product_sale SET deleted_at=$1 where id=$1`
	_,err=r.db.Exec(query,id)
	if err!=nil{
		return nil,err
	}
	return product,nil
}

func (r *PostsRepo) InfoProduct(id string)(*pb.Productsale,error){
	product:=pb.Productsale{}
	query:=`SELECT id,user_id,post_id,count,price,created_at from product_sale where deleted_at is null and saled_at is null adn id =$1`
	err:=r.db.QueryRow(query,id).Scan(
		&product.Id,
		&product.UserId,
		&product.PostId,
		&product.Count,
		&product.Price,
		&product.CreatedAt,
	)
	if err!=nil{
		return nil,err
	}
	return &product,nil

}

func (r *PostsRepo) GetAllProductsUser(id string)([]*pb.Productsale,error){
	products:=[]*pb.Productsale{}
	query:=`SELECT id,user_id,post_id,count,price,created_at from product_sale where deleted_at is null adn user_id=$1`
	rows,err:=r.db.Query(query,id)
	if err!=nil{
		return nil,err
	}
	for rows.Next(){
		product:=pb.Productsale{}
		err:=rows.Scan(
			&product.Id,
			&product.UserId,
			&product.Count,
			&product.Price,
			&product.CreatedAt,
		)
		if err!=nil{
			return nil,err
		}
		products = append(products, &product)
	}
	return products,nil
}

func (r *PostsRepo)GetingCountSaledPro(id string)(*pb.SaledCount,error){
	query:=`SELECT count(*) from product_sale where deleted_at is null and saled_at is not null and post_id=$1`
	var count int
	err:=r.db.QueryRow(query,id).Scan(&count)
	if err!=nil{
		return nil,err
	}
	return &pb.SaledCount{Count: int64(count)},nil
}
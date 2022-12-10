package service

import (
	"context"

	pb "github.com/Trendyol/post_service/genproto"
	"github.com/Trendyol/post_service/pkg/logger"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (s *PostsService) ProductSale(ctx context.Context,product *pb.ProductSaleReq)(*pb.Productsale,error){
	id,err:=uuid.NewV4()
	if err!=nil{
		s.logger.Error("failed while generating uuid for product_sale",logger.Error(err))
		return nil,status.Error(codes.Internal,"failed while generating uuid for product_sale")
	}
	product.Id=id.String()
	newPro,err:=s.repo.Product_sale(product)
	if err!=nil{
		s.logger.Error("failed while adding product to basket",logger.Error(err))
		return nil,status.Error(codes.Internal,"failed while adding product to basket")
	}
	return newPro,nil
}

func (s *PostsService) SaleProductDel(ctx context.Context,id *pb.WithId)(*pb.Productsale,error){
	product,err:=s.repo.SaleProductDel(id.Id)
	if err!=nil{
		s.logger.Error("failed while deleting product from basket",logger.Error(err))
		return nil,status.Error(codes.Internal,"failed shile deleting product from basket")
	}
	return product,nil
}

func(s *PostsService) GetAllProductsUser(ctx context.Context,id *pb.WithId)(*pb.ProductSales,error){
	products,err:=s.repo.GetAllProductsUser(id.Id)
	if err!=nil{
		s.logger.Error("failed while getting all products of user",logger.Error(err))
		return nil,status.Error(codes.Internal,"failed while getting all products of user")
	}
	return &pb.ProductSales{Products: products},nil
}

func (s *PostsService) InfoProduct(ctx context.Context,id *pb.WithId)(*pb.Productsale,error){
	products,err:=s.repo.InfoProduct(id.Id)
	if err!=nil{
		s.logger.Error("failed while gatting info product",logger.Error(err))
		return nil,status.Error(codes.Internal,"failed while getting info product")
	}
    return products,nil
}

func (s *PostsService) GetingCountSaledPro(ctx context.Context,id *pb.WithId)(*pb.SaledCount,error){
	count,err:=s.repo.GetingCountSaledPro(id.Id)
	if err!=nil{
		s.logger.Error("failed while getting count saled products",logger.Error(err))
		return nil,status.Error(codes.Internal,"failed while getting count saled products")
	}
	return count,nil
}

func (s *PostsService) GettingAllSalePro(ctx context.Context,empty *pb.Empty)(*pb.ProductSales,error){
	products,err:=s.repo.GettingAllSalePro()
	if err!=nil{
		s.logger.Error("failed while getting all sale Products")
		return nil,status.Error(codes.Internal,"failed while getting all saleproducts")	
	}
	return &pb.ProductSales{Products: products},nil

}
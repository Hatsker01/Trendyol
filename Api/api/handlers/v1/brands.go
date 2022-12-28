package v1

import (
	"context"
	"net/http"
	"time"

	pb "github.com/Trendyol/Api/genproto"
	"github.com/Trendyol/Api/pkg/logger"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateBrand ...
// @Sumamry CreateBrand
// @Description This API for creating new brand
// @Security BearerAuth
// @Tags brand
// @Accept json
// @Produce json
// @Param brand body model.CreateBrandReq true "CreateBrand"
// @Success 200 {object} model.Brand!
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/post/brand [post]
func (h *handlerV1) CreateBrand(c *gin.Context) {
	er := CheckClaims(h, c)
	if er == nil {
		newResponse(c, http.StatusUnauthorized, "failed while checking token")
		h.log.Error("Failed while checking token")
		return
	}
	var body pb.CreateBrandReq

	err := c.ShouldBindJSON(&body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "failed while blinding json")
		h.log.Error("failed while blinding json", logger.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	brand, err := h.serviceManager.PostService().CreateBrand(ctx, &body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "failed while creating new brand")
		h.log.Error("failed while creating new brand", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, brand)
}

// GetAllBrands ...
// @Summary Get All Brands
// @Description This API for getting all brands
// @Tags brand
// @Accept json
// @Produce json
// @Success 200 {object} model.Brands!
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/post/brand/getAll [get]
func (h *handlerV1) GetAllBrands(c *gin.Context) {
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	brands, err := h.serviceManager.PostService().GetAllBrands(ctx, &pb.Empty{})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("failed while getting all brands", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, brands)
}

// DeleteBrand ...
// @Summary DeleteBrand
// @Description This API for deleting brand
// @Security BearerAuth
// @Tags brand
// @Accept json
// @Produce json
// @Param id path string true "Brand_id"
// @Success 200 {object} model.Brand
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/post/brand/{id} [delete]
func (h *handlerV1) DeleteBrand(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	er := CheckClaims(h, c)
	if er == nil {
		newResponse(c, http.StatusUnauthorized, "failed while checking token")
		h.log.Error("failed while checking token")
		return
	}
	id :=c.Param("id")

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	brand,err:=h.serviceManager.PostService().DeleteBrand(ctx,&pb.WithId{Id: id})
	if err!=nil{
		newResponse(c,http.StatusInternalServerError,"failed while deleting brand with id")
		h.log.Error("failed while delting brand by id",logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted,brand)
}

// GetPostByBrand ...
// @Summary GetPostByBrand
// @Description This API for getting post bu brand
// @Security BearerAuth
// @Tags brand
// @Accept json
// @Produce json
// @Param id path string true "Brand_id"
// @Success 200 {object} model.Posts
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/post/brand/getPost/{id} [get]
func (h *handlerV1) GetPostByBrand(c *gin.Context)  {
	er:=CheckClaims(h,c)
	if er==nil{
		newResponse(c,http.StatusUnauthorized,"failed while checking token")
		h.log.Error("failed while checking token")
		return
	}

	id:=c.Param("id")

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	posts,err:=h.serviceManager.PostService().GetPostByBrand(ctx,&pb.WithId{Id: id})
	if err!=nil{
		newResponse(c,http.StatusInternalServerError,"failed while getting posts by brand")
		h.log.Error("failed while getting posts by brand",logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted,posts)
}

// GetBrandById ...
// @Summary GetBrandById
// @Description This API for getting brand by id
// @Security BearerAuth
// @Tags brand
// @Accept json
// @Produce json
// @Param id path string true "Brand_id"
// @Success 200 {object} model.Brand
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/post/brand/getByid/{id} [get]
func (h *handlerV1) GetBrandById(c *gin.Context){
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	er:=CheckClaims(h,c)
	if er==nil{
		newResponse(c,http.StatusUnauthorized,"failed while checking token")
		h.log.Error("failed while checking token")
		return
	}
	id :=c.Param("id")
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	brand,err:=h.serviceManager.PostService().GetBrandById(ctx,&pb.WithId{Id: id})
	if err!=nil{
		newResponse(c,http.StatusInternalServerError,"failed while getting brand by id")
		h.log.Error("failed while getting brand by id",logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted,brand)
}

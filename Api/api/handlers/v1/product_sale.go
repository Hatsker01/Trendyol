package v1

import (
	"context"
	"net/http"
	"time"

	pb "github.com/Trendyol/Api/genproto"
	"github.com/Trendyol/Api/pkg/logger"
	"github.com/gin-gonic/gin"
)

// ProductSale ...
// @Summary ProductSale
// @Description This API for adding product to basket
// @Security BearerAuth
// @Tags product
// @Accept json
// @Produce json
// @Param product body model.ProductSaleReq true "ProductSale"
// @Success 200 {object} model.Productsale
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/post/product [post]
func (h *handlerV1) ProductSale(c *gin.Context) {
	er := CheckClaims(h, c)
	if er == nil {
		newResponse(c, http.StatusUnauthorized, "failed while checking token")
		h.log.Error("failed while checking token")
		return
	}
	var body pb.ProductSaleReq
	err := c.ShouldBindJSON(&body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("failed while blinding json", logger.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	product, err := h.serviceManager.PostService().ProductSale(ctx, &body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("failed while adding product into basket", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, product)
}

// SaleProductDel ....
// @Summary SaleProductDelete
// @Description This API for deleting product from basket
// @Security BearerAuth
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "Product_id"
// @Success 200 {object} model.Productsale
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/post/product/{id} [delete]
func (h *handlerV1) SaleProductDel(c *gin.Context) {
	er := CheckClaims(h, c)
	if er == nil {
		newResponse(c, http.StatusUnauthorized, "failed while checking token")
		h.log.Error("failed while checking token")
		return
	}
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	product, err := h.serviceManager.PostService().SaleProductDel(ctx, &pb.WithId{Id: id})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("failed while deleting product from basket", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, product)
}

// GetAllProductsUser ...
// @Summary GetAllProductsUser
// @Description This API for getting all products in user basket
// @Security BearerAuth
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "UserId"
// @Success 200 {object} model.ProductSales
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/post/product/userPro/{id} [get]
func (h *handlerV1) GetAllProductsUser(c *gin.Context) {
	er := CheckClaims(h, c)
	if er == nil {
		newResponse(c, http.StatusUnauthorized, "failed while checking token")
		h.log.Error("failed while checking token")
		return
	}
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	products, err := h.serviceManager.PostService().GetAllProductsUser(ctx, &pb.WithId{Id: id})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("failed while getting all products from user basket", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, products)
}

// InfoProduct ...
// @Summary InfoProduct
// @Description This API for taking information abot product saling
// @Security BearerAuth
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "Product_id"
// @Success 200 {object} model.Productsale
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/post/product/info/{id} [get]
func (h *handlerV1) InfoProduct(c *gin.Context) {
	er := CheckClaims(h, c)
	if er == nil {
		newResponse(c, http.StatusUnauthorized, "failed while checking token")
		h.log.Error("failed while checking token")
		return
	}
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(),time.Second * time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	product, err := h.serviceManager.PostService().InfoProduct(ctx, &pb.WithId{Id: id})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("failed while getting info about product", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, product)
}

// GetingCountSaledPro ...
// @Summary GetingCountSaledProduct
// @Description This API for getting count salled product
// @Security BearerAuth
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "Product_id"
// @Success 200 {object} model.SeledCount
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/post/product/{id} [get]
func (h *handlerV1) GettingCountSaledPro(c *gin.Context) {
	er := CheckClaims(h, c)
	if er == nil {
		newResponse(c, http.StatusUnauthorized, "failed while checking token")
		h.log.Error("failed while checking token")
		return
	}
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	product, err := h.serviceManager.PostService().GetingCountSaledPro(ctx, &pb.WithId{Id: id})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("failed while getting count saled product", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, product)
}

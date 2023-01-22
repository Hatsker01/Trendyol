package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/Trendyol/Api/api/model"
	pb "github.com/Trendyol/Api/genproto"
	"github.com/Trendyol/Api/pkg/logger"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// Get Post By Category ...
// @Summary Get All posts by Category
// @Description This API for getting posts by category ID
// @Security BearerAuth
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "Category_id"
// @Success 200 {object} model.Posts!
// @Failure 400 {object} model.StandardErrorModel
// @Failure 500 {object} model.StandardErrorModel
// @Router /v1/category/{id} [get]
func (h *handlerV1) GetPostByCategory(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	er := CheckClaims(h, c)
	if er == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "error while comparing token",
		})
		h.log.Fatal("failed in check token")
		return
	}
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	posts, err := h.serviceManager.PostService().GetPostByCategory(ctx, &pb.CatID{
		Id: id,
	})
	if err != nil {
		newResponse(c, http.StatusBadRequest, model.ErrInputBody.Error())
		h.log.Error("failed while getting post by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, posts)
}

// Create Category ...
// @Summary CreateCategory
// @Decription This API for create category
// @Security BearerAuth
// @Tags category
// @Accept json
// @Produce json
// @Param category body model.CategiryReq true "CreateCategory"
// @Success 200 {object} model.Category!
// @Failure 400 {object} model.StandardErrorModel
// @Failure 500 {object} model.StandardErrorModel
// @Router /v1/category [post]
func (h *handlerV1) CreateCategory(c *gin.Context) {
	er := CheckClaims(h, c)
	if er == nil {
		newResponse(c, http.StatusUnauthorized, model.ErrInputBody.Error())
		h.log.Error("failed while checking token")
		return
	}
	var body pb.CategoryReq

	err := c.ShouldBindJSON(&body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, model.ErrInputBody.Error())
		h.log.Error("error while creating category", logger.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	category, err := h.serviceManager.PostService().CreateCategory(ctx, &body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, model.ErrWhileCreate.Error())
		h.log.Error("error while creting category", logger.Error(err))
		return
	}
	c.JSON(http.StatusCreated, category)
}

// Get All Categories ...
// @Summary GetAllCategories
// @Description This API for getting all categories
// @Tags category
// @Accept json
// @Produce json
// @Success 200 {object} model.Categories!
// @Failure 400 {object} model.StandardErrorModel
// @Failure 500 {object} model.StandardErrorModel
// @Router /v1/category/getAll [get]
func (h *handlerV1) GetAllCategories(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	categories, err := h.serviceManager.PostService().GetAllCategories(ctx, &pb.Empty{})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, model.ErrInputBody.Error())
		h.log.Error("failed while getting all categories ", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, categories)
}

// Delete Category ...
// @Summary Delete Category By ID
// @Description This API for deleting category by ID
// @Security BearerAuth
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "Category_ID"
// @Success 200 {object} model.Category!
// @Failure 400 {object} model.StandardErrorModel
// @Failure 500 {object} model.StandardErrorModel
// @Router /v1/category/{id} [delete]
func (h *handlerV1) DeleteCategory(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	er := CheckClaims(h, c)
	if er == nil {
		newResponse(c, http.StatusUnauthorized, "error while checking token")
		h.log.Fatal("failed while check token")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	id := c.Param("id")

	category, err := h.serviceManager.PostService().DeleteCategory(ctx, &pb.CatID{Id: id})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, model.ErrIdNotFound.Error())
		h.log.Error("error while deleting category by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, category)
}

// GetCategory ...
// @Summary GetCategory
// @Description This API for getting Category by ID
// @Security BearerAuth
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "Category_ID"
// @Success 200 {object} model.Category!
// @Failure 400 {object} model.StandardErrorModel
// @Failure 500 {object} model.StandardErrorModel
// @Router /v1/category/getById/{id} [get]
func (h *handlerV1) GetCategory(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	er := CheckClaims(h, c)
	if er == nil {
		newResponse(c, http.StatusUnauthorized, "failed while checking token")
		h.log.Fatal("failed while cheking token")
		return
	}
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	category, err := h.serviceManager.PostService().GetCategory(ctx, &pb.CatID{Id: id})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, model.ErrIdNotFound.Error())
		h.log.Error("failed while getting category by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, category)
}

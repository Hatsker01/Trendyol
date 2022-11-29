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

// Create Post...
// @Summary CreatePost
// @Description This API for creating new post
// @Tags post
// @Accept json
// @Produce json
// @Param post body model.Post true "CreatePost"
// @Success 200 {object} model.Post!
// @Router /v1/post [post]
func (h *handlerV1) CreatePost(c *gin.Context) {
	var body pb.Post

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while to blind json", logger.Error(err))
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	post, err := h.serviceManager.PostService().CreatePost(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating post", logger.Error(err))
		return
	}
	c.JSON(http.StatusCreated, post)
}

// Update Post ...
// @Summary UpdatePost
// @Descrtiption This API for updating post
// @Security BearerAuth
// @Tags post
// @Accept json
// @Produce json
// @Param post body model.Post true "UpdatePost"
// @Success 200 {object} model.Post!
// @Router /v1/post [put]
func (h *handlerV1) UpdatePost(c *gin.Context) {
	var body pb.Post
	er := CheckClaims(h, c)
	if er == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "error while comparing token",
		})
		h.log.Fatal("failed while checking token")
		return
	}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while to blind json", logger.Error(err))
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	post, err := h.serviceManager.PostService().UpdatePost(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while to update post", logger.Error(err))
	}
	c.JSON(http.StatusAccepted, post)
}

// GetPostById
// @Summary GetPostById
// @Description This API for getting post by id
// @Security BearerAuth
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "Post_id"
// @Success 200 {object} model.Post!
// @Router /v1/post/{id} [get]
func (h *handlerV1) GetPostById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	er := CheckClaims(h, c)
	if er == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "error while comparing json",
		})
		h.log.Fatal("error while comparing token")
		return
	}

	idd := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	post, err := h.serviceManager.PostService().GetPostById(ctx, &pb.WithId{Id: idd})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting post by id ", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, post)
}

// GetAllPost
// @Summary GetAllPosts
// @Description This API for getting all posts
// @Security BearerAuth
// @Tags post
// @Accept json
// @Produce json
// @Success 200 {object} model.Posts!
// @Router /v1/posts [get]
func (h *handlerV1) GetAllPosts(c *gin.Context) {
	er := CheckClaims(h, c)
	if er == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while comparing token",
		})
		h.log.Fatal("failed in chech token")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	posts, err := h.serviceManager.PostService().GetAllPosts(ctx, &pb.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error(err.Error())
		return
	}
	c.JSON(http.StatusAccepted, posts)
}

// GetAlluser posts
// @Summary Get all user posts
// @Description This API for getting user posts
// @Security BearerAuth
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "User_id"
// @Success 200 {object} model.Posts!
// @Router /v1/post/getAllUserPosts/{id} [get]
func (h *handlerV1) GetAllUserPosts(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")
	er := CheckClaims(h, c)
	if er == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while comparing token",
		})
		h.log.Fatal("failed in comparing token")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	posts, err := h.serviceManager.PostService().GetAllUserPosts(ctx, &pb.WithId{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error(err.Error())
		return
	}
	c.JSON(http.StatusAccepted, posts)
}

// DeletePostById
// @Summary Delete Post By ID
// @Description This API for deleting post by ID
// @Security BearerAuth
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "Post_id"
// @Success 200 {object} model.Post!
// @Router /v1/post/delete/{id} [delete]
func (h *handlerV1) DeletePostbyId(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	er := CheckClaims(h, c)
	if er == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while comparing token",
		})
		h.log.Fatal("failed in token check")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	post, err := h.serviceManager.PostService().DeletePostById(ctx, &pb.WithId{Id: guid})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error(err.Error())
		return
	}
	c.JSON(http.StatusAccepted, post)
}

// StarPosts
// @Summary Sort posts by stars
// @Description This API for getting post sorting by stars
// @Security BearerAuth
// @Tags post
// @Accept json
// @Produce json
// @Success 200 {object} model.Posts!
// @Router /v1/posts/stars [get]
func (h *handlerV1) SortByStars(c *gin.Context) {
	er := CheckClaims(h, c)
	if er == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "error while comparing token",
		})
		h.log.Fatal("failed while token check")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	posts, err := h.serviceManager.PostService().StarPosts(ctx, &pb.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error(err.Error())
		return
	}
	c.JSON(http.StatusAccepted, posts)
}




















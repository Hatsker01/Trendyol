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

// Putting Like ...
// @Summary Putting Like
// @Description This API for putting like
// @Security BearerAuth
// @Tags like
// @Accept json
// @Produce json
// @Param like body model.Like true "PutLike"
// @Success 200 {object} model.Like!
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/like [post]
func (h *handlerV1) PutLike(c *gin.Context) {
	er := CheckClaims(h, c)
	if er == nil {
		newResponse(c, http.StatusUnauthorized, "failed while checking token")
		h.log.Error("failed while checking token")
		return
	}

	var body pb.Like
	err := c.ShouldBindJSON(&body)
	if err != nil {
		newResponse(c, http.StatusBadRequest, model.ErrInputBody.Error())
		h.log.Error("failed while unmarshaling body", logger.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	like, err := h.serviceManager.PostService().PutLike(ctx, &body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, model.ErrInputBody.Error())
		h.log.Error("failed while putting post", logger.Error(err))
		return
	}
	c.JSON(http.StatusCreated, like)

}

// TakeLike ...
// @Summary TakeLike
// @Description This API for taking like
// @Security BearerAuth
// @Tags like
// @Accept json
// @Produce json
// @Param id path string true "Like_id"
// @Success 200 {object} model.Like
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/like/takeLike/{id} [delete]
func (h *handlerV1) TakeLike(c *gin.Context) {
	er := CheckClaims(h, c)
	if er == nil {
		newResponse(c, http.StatusUnauthorized, "failed while checking token")
		h.log.Error("failed while checking token")
		return
	}

	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	like, err := h.serviceManager.PostService().TakeLike(ctx, &pb.WithId{Id: id})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, model.ErrInputBody.Error())
		h.log.Error("failed while taking like", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, like)
}

// GetAllPostLikesUser ...
// @Summary GetAllPostLikesUser
// @Description This API for getting all likes user
// @Security BearerAuth
// @Tags like
// @Accept json
// @Produce json
// @Param id path string true "User_id"
// @Success 200 {object} model.Like
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/like/getAllLikeuser/{id} [get]
func (h *handlerV1) GetAllPostLikesUser(c *gin.Context) {
	er := CheckClaims(h, c)
	if er == nil {
		newResponse(c, http.StatusUnauthorized, "failed while checking token in get all like user")
		h.log.Error("failed while checking token")
		return
	}
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	like, err := h.serviceManager.PostService().GetAllPostLikesUser(ctx, &pb.WithId{Id: id})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, model.ErrIdNotFound.Error())
		h.log.Error("error while gtting likes user put post", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, like)
}

// GetLikeInfo ...
// @Summary GetLikeInfo
// @Description This API for getting like by id
// @Security BearerAuth
// @Tags like
// @Accept json
// @Produce json
// @Param id path string true "Like_Id"
// @Success 200 {object} model.Like
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/like/{id} [get]
func (h *handlerV1) GetLikeInfo(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	er := CheckClaims(h, c)
	if er == nil {
		newResponse(c, http.StatusUnauthorized, "failed while checking token")
		h.log.Error("error while checking token in GetLikeInfo")
		return
	}
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	like, err := h.serviceManager.PostService().GetLikeInfo(ctx, &pb.LikeId{Id: id})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, model.ErrIdNotFound.Error())
		h.log.Error("failed while getting like info", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, like)
}

// GetPostLike ...
// @Summary GetPostLike
// @Description This API for getting post's like
// @Security BearerAuth
// @Tags like
// @Accept json
// @Produce json
// @Param id path string true "Post_id"
// @Success 200 {object} model.Likes
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/like/getPostLike/{id} [get]
func (h *handlerV1) GetPostLike(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	er := CheckClaims(h, c)
	if er == nil {
		newResponse(c, http.StatusUnauthorized, "failed while cheking token in GetPostLike")
		h.log.Error("failed while checking token in GetPostLike")
		return
	}

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	like, err := h.serviceManager.PostService().GetPostLike(ctx, &pb.WithId{Id: id})
	if err != nil {
		newResponse(c,http.StatusInternalServerError,model.ErrIdNotFound.Error())
		h.log.Error("error while getting post likes",logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted,like)
}

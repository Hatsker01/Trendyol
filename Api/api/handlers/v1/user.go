package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/Trendyol/Api/pkg/logger"
	"github.com/gin-gonic/gin"
	pb "github.com/Trendyol/Api/genproto"
	
)

// CreateUser ...
// @Summary CreateUser
// @Description This API for creating a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user request body model.CreateUser true "UserCreate"
// @Success 200 {object} model.User
// @Success 400 {object} response
// @Success 500 {object} response
// @Router   /v1/users/ [POST]
func (h *handlerV1) CreateUser(c *gin.Context) {
	var body pb.CreateUserReq

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":err.Error(),
		})
		h.log.Error("failed while to blind json",logger.Error(err))

	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	user, err := h.serviceManager.UserService().CreateUser(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":err.Error(),
		})
		h.log.Error("failed while creating user json",logger.Error(err))
	}
	c.JSON(http.StatusCreated, user)
}

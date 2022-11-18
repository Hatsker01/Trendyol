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
// @Param user request body CreateUser true "UserCreate"
// @Success 200 {object} CreateUser
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
type CreateUser struct {
	id         string `json:"id"`
	first_name string `json:"first_name" binding:"required"`
	last_name  string `json:"last_name" binding:"required"`
	username   string `json:"username" binding:"required,min=4"`
	phone      string `json:"phone" binding:"required,min=5"`
	email      string `json:"email" binding:"required,email"`
	password   string `json:"password" binding:"required,min=5"`
	address    string `json:"address" binding:"required,min=7"`
	gender     string `json:"gender" binding:"required"`
	role       string `json:"role" binding:"required"`
	postalcode string `json:"postalcode" binding:"required"`
}
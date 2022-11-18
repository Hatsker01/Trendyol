package v1

import (
	"context"
	"net/http"
	"time"

	pb "github.com/Trendyol/api/genproto"
	"github.com/Trendyol/api/internal/transport/rest/v1/model"
	"github.com/Trendyol/api/pkg/logger"
	"github.com/gin-gonic/gin"
)

func (h *handlerV1) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/", h.CreateUser)
	}
}

// CreateUser ...
// @Summary CreateUser
// @Desscription This API for creating a new user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user request body model.CreateUser true "UserCreate"
// @Success 200 {object} structs.User
// @Success 400 {object} response
// @Success 500 {object} response
// @Router   /v1/users/ [POST]
func (h *handlerV1) CreateUser(c *gin.Context) {
	var body pb.CreateUserReq

	err := c.ShouldBindJSON(&body)
	if err != nil {
		newResponse(c, http.StatusBadRequest, model.ErrInputBody.Error())
		h.log.Error("error while parse to json request body CreateUser", logger.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	user, err := h.serviceManager.UserService().CreateUser(ctx, &body)
	if err!=nil{
		newResponse(c, http.StatusInternalServerError,model.ErrWhileCreate.Error())
		h.log.Error("error while creating user CreateUser",logger.Error(err))
		return
	}
	c.JSON(http.StatusCreated,user)
}

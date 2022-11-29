package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pb "github.com/Trendyol/Api/genproto"
	"github.com/Trendyol/Api/pkg/logger"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateUser ...
// @Summary CreateUser
// @Description This API for creating a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body model.CreateUser true "CreateUser"
// @Success 200 {object} model.User!
// @Success 400 {object} Success
// @Success 500 {object} Success
// @Router   /v1/users [post]
func (h *handlerV1) CreateUser(c *gin.Context) {
	var body pb.CreateUserReq

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while to blind json", logger.Error(err))

	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	user, err := h.serviceManager.UserService().CreateUser(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating user json", logger.Error(err))
	}
	c.JSON(http.StatusCreated, user)
}

// UpdateUser ...
// @Summary UpdateUser
// @Description This API for updating a new user
// @Security BearerAuth
// @Tags user
// @Accept json
// @Produce json
// @Param user body model.User true "UpdateUser"
// @Success 200 {object} model.User!
// @Success 400 {object} Success
// @Success 500 {object} Success
// @Router /v1/users [put]
func (h *handlerV1) UpdateUser(c *gin.Context) {
	claims, err := GetClaims(*h, c)
	if err == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized access",
		})
		return
	}

	if claims.Role != "Authorized" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "token not found",
		})
		return
	}
	body := pb.User{}
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while to blind json", logger.Error(err))

	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	user, err := h.serviceManager.UserService().UpdateUser(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating user", logger.Error(err))
	}
	c.JSON(http.StatusAccepted, user)

}

// GetUserById ...
// @Summary GetUserById
// @Description This API for getting user by id
// @Security BearerAuth
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User_ID"
// @Success 200 {object} model.User!
// @Success 400 {object} Success
// @Success 500 {object} Success
// @Router /v1/user/getUserbyId/{id} [get]
func (h *handlerV1) GetUserById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	

	er := CheckClaims(h, c)
	if er == nil {
		fmt.Println(er)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while comparing token",
		})
		h.log.Fatal("failed in token check")
		return

	}
	// if err == nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "Unauthorized access",
	// 	})
	// 	return
	// }

	// fmt.Println("asdfasfa")
	// if claims. != "Authorized" {
	// 	c.JSON(http.StatusForbidden, gin.H{
	// 		"error": "token not found",
	// 	})
	// 	return
	// }
	// fmt.Println("asdfasfa")

	// var jspbMarshal protojson.MarshalOptions
	// jspbMarshal.UseProtoNames = true
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	responce, err := h.serviceManager.UserService().GetUserById(ctx, &pb.WithId{Id: guid})
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting user by id", logger.Error(err))
	}
	c.JSON(http.StatusAccepted, responce)
}

// GetAllUsers
// @Summary GetAllUsers
// @Description This API for getting all Users
// @Security BearerAuth
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} model.Users!
// @Success 400 {object} Success
// @Success 500 {object} Success
// @Router /v1/users/getAll [get]
func (h *handlerV1) GetAllUsers(c *gin.Context) {
	er := CheckClaims(h, c)
	if er == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized access",
		})
		return
	}
	// if er. != "Authorized" || claims.Iss != "admin" {
	// 	c.JSON(http.StatusForbidden, gin.H{
	// 		"error": "user not admin",
	// 	})
	// 	return
	// }
	cxt, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	users, err := h.serviceManager.UserService().GetAllUsers(cxt, &pb.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error while getting all users",
		})
		return
	}
	c.JSON(http.StatusAccepted, users)
}

// DeleteUserById
// @Summary Delete User By Id
// @Description This API for deleting user by ID
// @Security BearerAuth
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User_id"
// @Success 200 {object} model.User!
// @Success 400 {object} Success
// @Success 500 {object} Success
// @Router /v1/user/delete/{id} [delete]
func (h *handlerV1) DeleteUserById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	er := CheckClaims(h, c)
	if er == nil {
		fmt.Println(er)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while comparing token",
		})
		h.log.Fatal("failed in token check")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	user, err := h.serviceManager.UserService().DeleteUserById(ctx, &pb.WithId{Id: guid})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error while deleting user by id",
		})
		return
	}
	c.JSON(http.StatusAccepted, user)
}

type Success struct {
	Message string `json:"message"`
}

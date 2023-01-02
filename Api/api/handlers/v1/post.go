package v1

import (
	"context"
	"net/http"
	"strconv"
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
// @Param post body model.CreatePost true "CreatePost"
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
// @Tags post
// @Accept json
// @Produce json
// @Success 200 {object} model.Posts!
// @Router /v1/posts [get]
func (h *handlerV1) GetAllPosts(c *gin.Context) {
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

// GetPostsSortPrice ...
// @Summary Get posts by seperating by price
// @Description This API for getting posts sorting post by price
// @Security BearerAuth
// @Tags post
// @Accept json
// @Produce json
// @Param high path bool true "High"
// @Success 200 {object} model.Posts
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/post/getSortPrice/{high} [get]
func (h *handlerV1) PriceSep(c *gin.Context){
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	er:=CheckClaims(h,c)
	if er== nil{
		newResponse(c,http.StatusUnauthorized,"failed while checking token")
		h.log.Error("failed while checking token")
		return
	}

	param:=c.Param("high")
	high,err:=strconv.ParseBool(param)
	if err!=nil{
		newResponse(c,http.StatusInternalServerError,"failed while parsing string to bool")
		h.log.Error("failed while parsing string to bool",logger.Error(err))
		return
	}

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	posts,err:=h.serviceManager.PostService().GetPostsSortPrice(ctx,&pb.PriceSep{
		High: high,
	})
	if err!=nil{
		newResponse(c,http.StatusInternalServerError,"failed while getting posts sorting by price")
		h.log.Error("failed while getting posts sorting by price",logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted,posts)
}


// GetPostByPrice ...
// @Summary Sort Post By price
// @Description This API for getting post by price with max and min price
// @Security BearerAuth
// @Tags post
// @Accept json
// @Produce json
// @Param post body model.GetPostByPrice true "GettingPost"
// @Success 200 {object} model.Posts!
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/post/getByPrice [get]
func (h *handlerV1) GetPostByPrice(c *gin.Context){
	var body pb.GetPostPriceReq

	er:=CheckClaims(h,c)
	if er==nil{
		newResponse(c,http.StatusUnauthorized,"failed while checking token")
		h.log.Error("error while checking token")
		return
	}

	err:=c.ShouldBindHeader(&body)
	if err!=nil{
		newResponse(c,http.StatusInternalServerError,"failed while blinding json")
		h.log.Error("failed while blinding json",logger.Error(err))
		return 
	}

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	posts,err:=h.serviceManager.PostService().GetPostByPrice(ctx,&body)
	if err!=nil{
		newResponse(c,http.StatusInternalServerError,"failed while getting post by price")
		h.log.Error("failed while getting post by price",logger.Error(err))
		return
	}

	c.JSON(http.StatusAccepted,posts)
}

// GettingPostsByColor ...
// @Summary Getting Post By Sorting Color
// @Description This API for getting post sorting pb color
// @Security BearerAuth
// @Tags post
// @Accept json
// @Produce json
// @Param color path string true "Color"
// @Success 200 {object} model.Posts
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/post/getByColor/{color} [get]
func (h *handlerV1) GetingPostsByColor(c *gin.Context){
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	er:=CheckClaims(h,c)
	if er==nil{
		newResponse(c,http.StatusUnauthorized,"failed while checking token")
		h.log.Error("failed while checking token")
		return
	}
	color:=c.Param("color")

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	posts,err:=h.serviceManager.PostService().GetingPostsByColor(ctx,&pb.ColorReq{Color: color})
	if err!=nil{
		newResponse(c,http.StatusInternalServerError,"failed while getting posts by color")
		h.log.Error("failed while getting posts by color",logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted,posts)
}


// PutStar ...
// @Summary Putting Star for post
// @Description This API for putting star to post'
// @Security BearerAuth
// @Tags star
// @Accept json
// @Produce json
// @Param star body model.StarReq true "StarReq"
// @Success 200 {object} model.Stars
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/post/star [post]
func (h *handlerV1) StarReq(c *gin.Context){
	er:=CheckClaims(h,c)
	if er==nil{
		newResponse(c,http.StatusUnauthorized,"failed while checking token")
		h.log.Error("failed while checking token")
		return
	}
	var body pb.StarReq
	err:=c.ShouldBindJSON(&body)
	if err!=nil{
		newResponse(c,http.StatusInternalServerError,"failed while blinding json")
		h.log.Error("failed while blinding json",logger.Error(err))
		return
	}
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	star,err:=h.serviceManager.PostService().PutStar(ctx,&body)
	if err!=nil{
		newResponse(c,http.StatusInternalServerError,"failed while putting star to post")
		h.log.Error("failed while putting star to post",logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted,star)
}

// TakeStar ...
// @Summary Taking Star from post
// @Description This API for Taking star from post
// @Security BearerAuth
// @Tags star
// @Accept json
// @Produce json
// @Param id path string true "Post_id"
// @Success 200 {object} model.Stars
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/post/star/{id} [delete]
func (h *handlerV1) TakeStar(c *gin.Context){
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	er:=CheckClaims(h,c)
	if er==nil{
		newResponse(c,http.StatusUnauthorized,"failed while checking token")
		h.log.Error("failed while checking token")
		return
	}
	
	id:=c.Param("id")
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	star,err:=h.serviceManager.PostService().TakeStar(ctx,&pb.WithId{Id: id})
	if err!=nil{
		newResponse(c,http.StatusInternalServerError,"failed while taking star from post")
		h.log.Error("failed while taking star from post",logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted,star)
}

// GetStar ...
// @Summary Getting Avarage Star of Post
// @Description This API for getting avarage star from post
// @Security BearerAuth
// @Tags star
// @Accept json
// @Produce json
// @Param id path string true "Post_Id"
// @Success 200 {object} model.Stars
// @Success 400 {object} response
// @Success 500 {object} response
// @Router /v1/post/stars/{id} [get]
func (h *handlerV1) GetStar(c *gin.Context){
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	er:=CheckClaims(h,c)
	if er==nil{
		newResponse(c,http.StatusUnauthorized,"failed while checking token")
		h.log.Error("failed while checking token")
		return
	}
	id := c.Param("id")

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	star,err:=h.serviceManager.PostService().GetStar(ctx,&pb.WithId{Id: id})
	if err!=nil{
		newResponse(c,http.StatusInternalServerError,"failed while getting avarrage star of post")
		h.log.Error("failed while getting avarage star of post",logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted,star)
}
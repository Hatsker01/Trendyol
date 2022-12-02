package model

import "errors"

type JwtRequestModel struct {
	Token string `json:"token"`
}
type ResponseError struct {
	Error interface{} `json:"error"`
}

type ServerError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type Tokens struct {
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}

type CreateUser struct {
	Id         string `json:"id"`
	First_name string `json:"first_name" binding:"required"`
	Last_name  string `json:"last_name" binding:"required"`
	Username   string `json:"username" binding:"required,min=4"`
	Phone      string `json:"phone" binding:"required,min=5"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=5"`
	Address    string `json:"address" binding:"required,min=7"`
	Gender     string `json:"gender" binding:"required"`
	Role       string `json:"role" binding:"required"`
	Code       string `json:"code"`
	Postalcode string `json:"postalcode" binding:"required"`
}
type Users struct {
	User []User `json:"Users"`
}
type User struct {
	Id         string  `json:"id"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Username   string  `json:"username"`
	Phone      string  `json:"phone"`
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	Address    string  `json:"address"`
	Gender     string  `json:"gender"`
	Role       string  `json:"role"`
	Code       string  `json:"code`
	Postalcode string  `json:"postalcode"`
	Posts      []*Post `json:"Posts"`
	Color      string  `json:"color"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type Post struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Body        string   `json:"body"`
	AuthorId    string   `json:"author_id"`
	Stars       string      `json:"stars" binding:"required,min=0,max=5"`
	Rating      string   `json:"rating"`
	Price       string   `json:"price"`
	ProductType string   `json:"product_type"`
	Size_       []string `json:"size"`
	Color       string   `json:"color"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

type Category struct {
	ID         string `json:"id"`
	Post_ID    string `json:"post_id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Deleted_at string `json:"deleted_at"`
}

type CategiryReq struct {
	Post_id string `json:"post_id"`
	Name    string `json:"name"`
}
type Categories struct {
	Categories []Category `json:"categories"`
}

type Posts struct {
	Posts []Post `json:"Posts"`
}

type Like struct{
	Id string `json:"id"`
	User_id string 	`json:"user_id"`
	Post_id string `json:"post_id"`
	Created_at string `json:"created_at"`
	Deleted_at string `json:"deleted_at"`
}

type Likes struct {
	Likes []Like `json:"likes"`
}

type LikeId struct{
	Id string `json:"id"`
}

type ChengePass struct{
	Id string `json:"id"`
	OldPass string `json:"old_pass"`
	NewPass string `json:"new_pass"`
	VerifyNew string `json:"very_new"`
}

type ChangePassRes struct{
	Id string `json:"id"`
	NewPass string `json:"new_pass"`
}

var (
	ErrInputBody   = errors.New("error input body invalid")
	ErrIdNotFound  = errors.New("error id not found")
	ErrWhileCreate = errors.New("error while create")
	ErrWhileUpdate = errors.New("error while update")
)

package model

type JwtRequestModel struct {
	Token string `json:"token"`
}
type ResponseError struct{
	Error interface{} `json:"error"`
}

type ServerError struct{
	Status string `json:"status"`
	Message string `json:"message"`
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

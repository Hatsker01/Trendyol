package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, image/png, image/jpg, application/pdf, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func GinCorsMiddleware() Options {
	o := cors.Options{
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodPatch,
		},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
			"Content-Length",
			"image/png",
			"accept",
			"Accept-Encoding",
			"origin",
			"Cache-Control",
			"X-Requested-With",
			"application/json",
		},
		OptionsPassthrough: false,
		ExposedHeaders: []string{
			"application/json",
			"multipart/form-data",
			"Authorization",
			"application/pdf",
			"video/mp4",
			"Content-Type",
			"image/png",
			"image/jpg",
		},
		Debug:                true,
		OptionsSuccessStatus: 200,
	}

	return o
}

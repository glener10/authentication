package routes

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	middlewares "github.com/glener10/rotating-pairs-back/src/routes/middlewares"
)

func HandlerRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     getAllowedURLs(),
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		MaxAge:           12 * time.Hour,
	}))
	rateLimiter := middlewares.NewRateLimiter(11, time.Minute)
	r.Use(middlewares.RequestLimitMiddleware(rateLimiter))

	r.GET("/", middlewares.HelloWorld)

	r.Use(middlewares.AuthMiddleware())
	//r.Use(middlewares.HTTPSOnlyMiddleware())

	return r
}

func Listening(r *gin.Engine) {
	err := r.Run()
	if err != nil {
		fmt.Println("Error to up routes")
		os.Exit(-1)
	}
}

func getAllowedURLs() []string {
	if os.Getenv("ENV") == "development" {
		return []string{"*"}
	}
	allowedURLsString := os.Getenv("ALLOW_URLS")
	if allowedURLsString == "" {
		return nil
	}
	allowedUrlsString := strings.Split(allowedURLsString, "|")
	return allowedUrlsString
}

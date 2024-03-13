package routes

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	Middlewares "github.com/glener10/rotating-pairs-back/src/routes/middlewares"
)

func HandlerRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", Middlewares.HelloWorld)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     getAllowedURLs(),
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	rateLimiter := Middlewares.NewRateLimiter(11, time.Minute)
	r.Use(Middlewares.RequestLimitMiddleware(rateLimiter))
	r.Use(Middlewares.AuthMiddleware())
	//r.Use(Middlewares.HTTPSOnlyMiddleware())

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

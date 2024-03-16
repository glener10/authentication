package routes

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/glener10/authentication/docs"
	middlewares "github.com/glener10/authentication/src/routes/middlewares"
	user_controller "github.com/glener10/authentication/src/user/controllers"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandlerRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     getAllowedURLs(),
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	rateLimiter := middlewares.NewRateLimiter(11, time.Minute)
	r.Use(middlewares.RequestLimitMiddleware(rateLimiter))
	r.Use(middlewares.TimeoutMiddleware())

	r.GET("/", middlewares.HelloWorld)

	//r.Use(middlewares.AuthMiddleware())
	//r.Use(middlewares.HTTPSOnlyMiddleware())

	r.POST("/user", user_controller.CreateUser)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}

func Listening(r *gin.Engine) {
	err := r.Run()
	if err != nil {
		log.Fatalf("error to up routes")
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

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
	change_email_controller "github.com/glener10/authentication/src/user/controllers/change_email"
	change_password_controller "github.com/glener10/authentication/src/user/controllers/change_password"
	create_user_controller "github.com/glener10/authentication/src/user/controllers/create_user"
	delete_user_controller "github.com/glener10/authentication/src/user/controllers/delete_user"
	find_user_controller "github.com/glener10/authentication/src/user/controllers/find_user"
	login_controller "github.com/glener10/authentication/src/user/controllers/login"
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

	if os.Getenv("ENV") != "development" {
		rateLimiter := middlewares.NewRateLimiter(11, time.Minute)
		r.Use(middlewares.RequestLimitMiddleware(rateLimiter))
		r.Use(middlewares.TimeoutMiddleware())
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/", middlewares.HelloWorld)
	r.POST("/user", create_user_controller.CreateUser)
	r.POST("/login", login_controller.Login)

	r.Use(middlewares.JwtMiddleware())
	//r.Use(middlewares.HTTPSOnlyMiddleware())

	r.GET("/user/:find", find_user_controller.FindUser)
	r.DELETE("/user/:find", delete_user_controller.DeleteUser)
	r.PUT("/user/changePassword/:find", change_password_controller.ChangePassword)
	r.PUT("/user/changeEmail/:find", change_email_controller.ChangeEmail)
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

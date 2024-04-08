package routes

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/glener10/authentication/docs"
	admin_delete_user_controller "github.com/glener10/authentication/src/admin/controllers/admin_delete_user"
	admin_find_all_logs_controller "github.com/glener10/authentication/src/admin/controllers/admin_find_all_logs"
	admin_find_all_users_controller "github.com/glener10/authentication/src/admin/controllers/admin_find_all_users"
	admin_find_user_controller "github.com/glener10/authentication/src/admin/controllers/admin_find_user"
	admin_find_user_all_logs_controller "github.com/glener10/authentication/src/admin/controllers/admin_find_user_all_logs"
	promote_user_admin_controller "github.com/glener10/authentication/src/admin/controllers/promote_user_admin"
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

	if os.Getenv("ENV") != "production" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		r.GET("/", middlewares.HelloWorld)
	} else {
		rateLimiter := middlewares.NewRateLimiter(11, time.Minute)
		r.Use(middlewares.RequestLimitMiddleware(rateLimiter))
		r.Use(middlewares.TimeoutMiddleware())
		r.Use(middlewares.HTTPSOnlyMiddleware())
	}

	r.POST("/users", create_user_controller.CreateUser)
	r.POST("/login", login_controller.Login)

	r.Use(middlewares.JwtMiddleware())
	r.Use(middlewares.InactiveUserMiddlware())

	r.GET("/users/:find", find_user_controller.FindUser)
	r.DELETE("/users/:find", delete_user_controller.DeleteUser)
	r.PUT("/users/changePassword/:find", change_password_controller.ChangePassword)
	r.PUT("/users/changeEmail/:find", change_email_controller.ChangeEmail)

	r.Use(middlewares.AdminMiddleware())

	r.POST("/admin/users/promote/:find", promote_user_admin_controller.PromoteUserAdmin)
	r.DELETE("/admin/users/:find", admin_delete_user_controller.AdminDeleteUser)
	r.GET("/admin/users/:find", admin_find_user_controller.AdminFindUser)
	r.GET("/admin/users", admin_find_all_users_controller.AdminFindAllUsers)
	r.GET("/admin/logs", admin_find_all_logs_controller.AdminFindAllLogs)
	r.GET("/admin/logs/:find", admin_find_user_all_logs_controller.AdminFindUserAllLogs)
	return r
}

func Listening(r *gin.Engine) {
	var err error
	if os.Getenv("ENV") != "development" {
		err = r.Run()
	} else {
		addr := "localhost:8080"
		err = r.Run(addr)
	}
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

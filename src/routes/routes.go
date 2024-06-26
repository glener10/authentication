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
	admin_inative_user_controller "github.com/glener10/authentication/src/admin/controllers/admin_inative_user"
	promote_user_admin_controller "github.com/glener10/authentication/src/admin/controllers/promote_user_admin"
	middlewares "github.com/glener10/authentication/src/routes/middlewares"
	active_2fa_controller "github.com/glener10/authentication/src/user/controllers/active_2fa"
	change_email_controller "github.com/glener10/authentication/src/user/controllers/change_email"
	change_password_controller "github.com/glener10/authentication/src/user/controllers/change_password"
	change_password_in_recovery_controller "github.com/glener10/authentication/src/user/controllers/change_password_in_recovery"
	create_user_controller "github.com/glener10/authentication/src/user/controllers/create_user"
	delete_user_controller "github.com/glener10/authentication/src/user/controllers/delete_user"
	desactive_2fa_controller "github.com/glener10/authentication/src/user/controllers/desactive_2fa"
	find_user_controller "github.com/glener10/authentication/src/user/controllers/find_user"
	login_controller "github.com/glener10/authentication/src/user/controllers/login"
	login_2fa_controller "github.com/glener10/authentication/src/user/controllers/login_2fa"
	send_change_email_code_controller "github.com/glener10/authentication/src/user/controllers/send_change_email_code"
	send_email_verification_code_controller "github.com/glener10/authentication/src/user/controllers/send_email_verification_code"
	send_password_recovery_code_controller "github.com/glener10/authentication/src/user/controllers/send_password_recovery_code"
	verify_change_email_code_controller "github.com/glener10/authentication/src/user/controllers/verify_change_email_code"
	verify_email_controller "github.com/glener10/authentication/src/user/controllers/verify_email"
	verify_password_recovery_code_controller "github.com/glener10/authentication/src/user/controllers/verify_password_recovery_code"
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
	r.POST("/users/sendPasswordRecoveryCode/:find", send_password_recovery_code_controller.SendPasswordRecoveryCode)
	r.POST("/users/verifyPasswordRecoveryCode/:find", verify_password_recovery_code_controller.VerifyPasswordRecoveryCode)
	r.POST("/users/changePasswordInRecovery/:find", change_password_in_recovery_controller.ChangePasswordInRecovery)

	r.Use(middlewares.JwtSignatureMiddleware())
	r.Use(middlewares.InactiveUserMiddlware())

	r.POST("/login/2fa", login_2fa_controller.Login2FA)

	r.Use(middlewares.TwofaMiddleware())

	r.GET("/users/:find", find_user_controller.FindUser)
	r.DELETE("/users/:find", delete_user_controller.DeleteUser)
	r.PUT("/users/changePassword/:find", change_password_controller.ChangePassword)
	r.PUT("/users/changeEmail/:find", change_email_controller.ChangeEmail)
	r.POST("/users/sendEmailVerificationCode/:find", send_email_verification_code_controller.SendEmailVerificationCode)
	r.POST("/users/sendChangeEmailCode/:find", send_change_email_code_controller.SendChangeEmailCode)
	r.POST("/users/verifyEmail/:find", verify_email_controller.VerifyEmail)
	r.POST("/users/verifyChangeEmailCode/:find", verify_change_email_code_controller.VerifyChangeEmailCode)
	r.POST("/users/2fa/desactive/:find", desactive_2fa_controller.Desactive2FA)
	r.POST("/users/2fa/active/:find", active_2fa_controller.Active2FA)

	r.Use(middlewares.OnlyAdminMiddleware())

	r.POST("/admin/users/promote/:find", promote_user_admin_controller.PromoteUserAdmin)
	r.POST("/admin/users/inative/:find", admin_inative_user_controller.AdminInativeUser)
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

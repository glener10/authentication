package user_usecases

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_entities "github.com/glener10/authentication/src/user/entities"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Repository user_interfaces.IUserRepository
}

func (u *Login) Executar(c *gin.Context, user user_dtos.CreateUserRequest) {
	userInDb, err := u.Repository.FindUser(user.Email)
	if err != nil {
		statusCode := http.StatusUnauthorized
		c.JSON(statusCode, gin.H{"error": "email or password is incorret", "statusCode": statusCode})
		return
	}

	passwordIsValid := bcrypt.CompareHashAndPassword([]byte(userInDb.Password), []byte(user.Password))
	if passwordIsValid == nil {
		statusCode := http.StatusUnauthorized
		c.JSON(statusCode, gin.H{"error": "email or password is incorret", "statusCode": statusCode})
		return
	}

	signedToken, err := GenerateJwt(userInDb)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}

	response := user_dtos.LoginResponse{
		Jwt: *signedToken,
	}
	c.JSON(http.StatusOK, response)
}

func GenerateJwt(user *user_entities.User) (*string, error) {
	claims := jwt.MapClaims{
		"Id":    user.Id,
		"Email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, errors.New("error in token signature")
	}
	return &signedToken, nil
}

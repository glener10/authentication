package jwt_usecases

import (
	"errors"
	"net/http"
	"os"
	"time"

	user_entities "github.com/glener10/authentication/src/user/entities"
	"github.com/golang-jwt/jwt"
)

func CheckSignatureAndReturnClaims(tokenFromHeader string) (jwt.MapClaims, *int, error) {
	token, err := jwt.Parse(tokenFromHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		statusCode := http.StatusUnauthorized
		return nil, &statusCode, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		statusCode := http.StatusBadRequest
		return nil, &statusCode, errors.New("error to map claims")
	}
	return claims, nil, nil
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

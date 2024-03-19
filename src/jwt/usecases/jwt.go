package jwt_usecases

import (
	"errors"
	"net/http"
	"os"

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

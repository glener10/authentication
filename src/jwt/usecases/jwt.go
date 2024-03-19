package jwt_usecases

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

func CheckSignatureAndReturnClaims(tokenFromHeader string) (jwt.MapClaims, error) {
	if tokenFromHeader == "" {
		return nil, errors.New("invalid token format")
	}

	token, err := jwt.Parse(tokenFromHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("error to map claims")
	}
	return claims, nil
}

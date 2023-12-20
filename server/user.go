package server

import (
	"errors"
	"log"
	"time"

	"github.com/OmidRasouli/amuse-park/statics"
	"github.com/dgrijalva/jwt-go/v4"
)

func generateToken(username string, expireTime time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(expireTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(statics.SecretKey))
	if err != nil {
		return "", nil
	}
	return signedToken, nil
}

func validationToken(signedToken string) (string, error) {
	parsedToken, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(statics.SecretKey), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		username := claims["username"].(string)
		return username, nil
	}
	return "", errors.New("Invalid token")
}

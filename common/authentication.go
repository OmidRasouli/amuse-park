package common

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/OmidRasouli/amuse-park/statics"
	"github.com/dgrijalva/jwt-go/v4"
)

type JWTClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userID string, expireTime time.Duration) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(expireTime)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(statics.SecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateJWTToken(signedToken string) (string, error) {
	parsedToken, err := jwt.ParseWithClaims(signedToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(statics.SecretKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := parsedToken.Claims.(*JWTClaims)
	if !ok || !parsedToken.Valid {
		return "", errors.New("invalid token")
	}
	return claims.UserID, nil
}

func Authentication(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		return "", errors.New("invalid token")
	}

	token := authHeaderParts[1]
	if token == "" {
		return "", errors.New("missing token")
	}

	userID, err := ValidateJWTToken(token)
	if err != nil {
		return "", err
	}
	return userID, nil
}

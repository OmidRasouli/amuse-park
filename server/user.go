package server

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/OmidRasouli/amuse-park/statics"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

type JWTClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func generateToken(userID string, expireTime time.Duration) (string, error) {
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

func validateToken(signedToken string) (string, error) {
	parsedToken, err := jwt.ParseWithClaims(signedToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(statics.SecretKey), nil
	})
	if err != nil {
		return "", errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(*JWTClaims)
	if !ok || !parsedToken.Valid {
		return "", errors.New("invalid token")
	}

	return claims.UserID, nil
}

type UserAccount struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func Register(c *gin.Context) {
	var userAccount UserAccount

	if err := c.ShouldBindJSON(&userAccount); err != nil {
		c.String(http.StatusBadRequest, "bad request: %v", err)
		return
	}

	token, err := generateToken(userAccount.UserID, time.Hour*100)
	if err != nil {
		log.Printf("this error occurred while creating jwt: %v", err)
		c.String(http.StatusInternalServerError, "internal error occurred: %v", err)
		return
	}

	userAccount.Token = token
	c.JSON(http.StatusOK, userAccount)
}

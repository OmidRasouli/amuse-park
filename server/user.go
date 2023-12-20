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

type UserAccount struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func Register(c *gin.Context) {
	var userAccount UserAccount

	if err := c.ShouldBind(&userAccount); err != nil {
		c.String(http.StatusBadRequest, "bad request: %v", err)
	}

	token, err := generateToken(userAccount.Username, time.Hour*100)
	if err != nil {
		log.Printf("this error occurred while creating jwt: %v", err)
		c.String(http.StatusInternalServerError, "internal error occurred: %v", err)
	}

	userAccount.Token = token
	c.JSON(http.StatusOK, userAccount)
}

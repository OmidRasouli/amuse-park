package server

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/OmidRasouli/amuse-park/common"
	"github.com/OmidRasouli/amuse-park/models"
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
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var userAccount UserAccount

	err := c.ShouldBindJSON(&userAccount)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request: %v", err)
		return
	}

	account, err := createAccount(&userAccount)

	if err != nil {
		log.Printf("this error occurred while creating an account: %v", err)
		c.String(http.StatusInternalServerError, "internal error occurred: %v", err)
		return
	}

	log.Printf("New account \"%v\" add to database.", account.Username)

	auth, err := CreateAuthentication(account, userAccount.Password)

	if err != nil {
		log.Printf("this error eccourred while creating an authentication: %v", err)
		c.String(http.StatusInternalServerError, "internal error occurred: %v", err)
		return
	}

	log.Printf("New authentication \"%v\" add to database.", auth.Username)

	token, err := generateToken(account.UserID.String(), time.Hour*100)
	if err != nil {
		log.Printf("this error occurred while creating jwt: %v", err)
		c.String(http.StatusInternalServerError, "internal error occurred: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account": account,
		"token":   token,
	})
}

package server

import (
	"log"
	"net/http"
	"time"

	"github.com/OmidRasouli/amuse-park/common"
	"github.com/OmidRasouli/amuse-park/models"
	"github.com/gin-gonic/gin"
)

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

	log.Printf("New account \"%v\" added to database.", account.Username)

	auth, err := CreateAuthentication(account, userAccount.Password)

	if err != nil {
		log.Printf("this error eccourred while creating an authentication: %v", err)
		c.String(http.StatusInternalServerError, "internal error occurred: %v", err)
		return
	}

	log.Printf("New authentication \"%v\" added to database.", auth.Username)

	token, err := common.GenerateToken(account.UserID.String(), time.Hour*100)
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

func UpdateProfile(c *gin.Context) {
	var profile models.Profile

	userID, err := common.Authentication(c.Request.Header)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
	}

	err = c.ShouldBindJSON(&profile)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request: %v", err)
		return
	}

	if userID != profile.ID.String() {
		log.Printf("user id: %v", userID)
		log.Printf("user id: %v", profile.ID.String())
		c.String(http.StatusUnauthorized, "invalid information")
		return
	}

	if !common.IsValidEmail(profile.Email) {
		c.String(http.StatusBadRequest, "email is not valid")
	}

	err = updateProfile(profile)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request: %v", err)
		return
	}

	c.JSON(http.StatusOK, profile)
}

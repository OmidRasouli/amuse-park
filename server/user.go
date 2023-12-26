package server

import (
	"log"
	"net/http"

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

	token := c.GetHeader("token")

	c.JSON(http.StatusOK, gin.H{
		"account": account,
		"token":   token,
	})
}

func UpdateProfile(c *gin.Context) {
	var profile models.Profile

	err := c.ShouldBindJSON(&profile)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request: %v", err)
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

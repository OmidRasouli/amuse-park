package main

import (
	"fmt"

	"github.com/OmidRasouli/amuse-park/database"
	"github.com/OmidRasouli/amuse-park/models"
	"github.com/OmidRasouli/amuse-park/server"
	"github.com/OmidRasouli/amuse-park/statics"
	"github.com/gin-gonic/gin"
)

func main() {
	statics.Read()

	database.Initialize(&models.Account{}, &models.Authentication{}, &models.Profile{})

	fmt.Println("Hello there...")

	router := gin.Default()
	api := router.Group("/")
	{
		api.POST("/register", server.Register)
	}

	router.Run(":8080")
}

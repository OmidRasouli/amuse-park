package routing

import (
	"github.com/OmidRasouli/amuse-park/server"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()
	api := router.Group("/")
	{
		api.POST("/register", refreshToken, server.Register)
		api.POST("/update-profile", authentication, server.UpdateProfile)
	}

	router.Run(":8080")
}

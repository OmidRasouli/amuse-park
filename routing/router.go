package routing

import (
	"github.com/OmidRasouli/amuse-park/server"
	"github.com/gin-gonic/gin"
)

func Initialize() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/account")
	{
		api.POST("/refresh", server.RefreshToken)
		api.POST("/register", server.Register)
		api.POST("/update-profile", server.Authentication, server.UpdateProfile)
	}

	return router
}

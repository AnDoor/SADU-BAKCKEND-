package routes

import (
	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/handlers"
)

func RegisterTeamsRoutes(r *gin.Engine) {
	teams := r.Group("/teams")
	{
		teams.GET("/", handlers.GetTeams)
		teams.POST("/add", handlers.AddTeams)
	}
}



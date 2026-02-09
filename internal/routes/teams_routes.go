package routes

import (
	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/handlers"
)

func RegisterTeamRoutes(r *gin.RouterGroup, teamHandler *handlers.TeamHandler) {
	r.GET("", teamHandler.GetAllTeamHandler)
	r.GET("/:id", teamHandler.GetTeamByIdHandler)
	r.POST("/create", teamHandler.CreateTeamHandler)
	r.PUT("/edit/:id", teamHandler.EditTeamHandler)
	r.DELETE("/delete/:id", teamHandler.DeleteTeam)
}


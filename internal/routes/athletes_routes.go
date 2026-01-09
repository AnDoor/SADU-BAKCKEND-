package routes

import (
	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/handlers"
)

func RegisterAthletesRoutes(r *gin.RouterGroup, athleteHandler *handlers.AthleteHandler) {

	r.GET("", athleteHandler.GetAthletes)
	r.GET("/:id", athleteHandler.GetAthletesByID)
	r.POST("/create", athleteHandler.CreateNewAthlete)
	r.PUT("/edit/:id", athleteHandler.EditAthleteByID)
	r.DELETE("/delete/:id", athleteHandler.DeleteAthleteByID)

}

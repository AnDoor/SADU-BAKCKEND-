package routes

import (
	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/handlers"
)

func RegisterAthletesRoutes(r *gin.RouterGroup) {

	r.GET("", handlers.GetAthletes)
	r.GET("/:id", handlers.GetAthletesByID)
	r.POST("/", handlers.CreateNewAthlete)
	r.PUT("/:id", handlers.EditAthleteByID)
	r.DELETE("/:id", handlers.DeleteAthleteByID)

}

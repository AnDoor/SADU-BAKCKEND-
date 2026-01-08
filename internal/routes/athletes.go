package routes

import (
	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/handlers"
)

func RegisterAthletesRoutes(r *gin.RouterGroup, athleteHandler *handlers.AthleteHandler) {

	r.GET("", athleteHandler.GetAthletes)
	r.GET("/:id")
	r.POST("/:id")
	r.PUT("/:id")
	r.DELETE("/:id")

}

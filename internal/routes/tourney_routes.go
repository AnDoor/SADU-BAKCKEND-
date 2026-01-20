package routes

import (
	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/handlers"
)

func RegisterTourney(r *gin.RouterGroup, TourneyHandler *handlers.TourneyHandler){
	r.GET("", TourneyHandler.GetAllTourneyHandler)
	r.GET("/:id", TourneyHandler.GetTourneyByIdHandler)
	r.POST("/create", TourneyHandler.CreateUniversityHandler)
	r.PUT("/edit/:id", TourneyHandler.UpdateTourneyHandler)
	r.DELETE("/delete/:id", TourneyHandler.DeleteTourneyHandler)

}
package routes

import (
	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/handlers"
)

func RegisterEventsRoutes(r *gin.Engine) {
	events := r.Group("/events")
	{
		events.GET("/", handlers.GetBareEvents)
	}
}

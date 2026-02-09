package routes

import (
	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/handlers"
)

func RegisterEventsRouters(r *gin.RouterGroup, handler *handlers.EventHandler){
	r.GET("", handler.GetAllEventsHandler)
	r.GET("/:id", handler.GetEventByIDHandler)
	r.POST("/create", handler.CreateEventHandler)
	r.PUT("/edit/:id", handler.EditEventHandler)
	r.DELETE("/delete/:id", handler.DeleteEventHandler)
}
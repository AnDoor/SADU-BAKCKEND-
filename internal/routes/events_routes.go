package routes

import (
	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/handlers"
	"uneg.edu.ve/servicio-sadu-back/internal/middlewares"
)

func RegisterEventsRouters(r *gin.RouterGroup, handler *handlers.EventHandler) {
	r.GET("", handler.GetEventsHandler)
	r.GET("/:id", handler.GetEventsHandler)
	r.POST("/create", middlewares.AuthMiddleware(), handler.CreateEventHandler)
	r.PUT("/edit/:id", middlewares.AuthMiddleware(), handler.EditEventHandler)
	r.DELETE("/delete/:id", middlewares.AuthMiddleware(), handler.DeleteEventHandler)
}


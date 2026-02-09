package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/handlers"
	"uneg.edu.ve/servicio-sadu-back/internal/middlewares"
)

func RegisterUserRoutes(r *gin.RouterGroup, userHandler *handlers.UserHandler) {
	r.POST("/login", userHandler.LoginUserHandler)
	// r.POST("/register", userHandler.CreateDisciplineHandler)

	// Protected Route. Example of how to protect a route (Delete later)
	r.GET("/test", middlewares.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Acceso concedido"})
	})
}

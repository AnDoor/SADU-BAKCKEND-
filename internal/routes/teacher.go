package routes

import (
	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/handlers"
)

func RegisterTeachersRoutes(r *gin.Engine) {
	teachers := r.Group("/teachers")
	{
		teachers.GET("/", handlers.GetTeachers)
	}
}

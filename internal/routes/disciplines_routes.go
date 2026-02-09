package routes

import (
	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/handlers"
)

func RegisterDisciplines(r *gin.RouterGroup) {
	r.GET("", handlers.GetAllDisciplineHandler)
	r.GET("/:id", handlers.GetAllDisciplinesByIdHandler)
	r.POST("/", handlers.CreateDisciplineHandler)
	r.PUT("/:id", handlers.EditDisciplineHandler)
	r.DELETE("/:id", handlers.DeleteDisciplineHandler)
}

package routes

import (
	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/handlers"
)

func RegisterDisciplines(r *gin.RouterGroup, disciplineHandler *handlers.DisciplineHandler) {
	r.GET("", disciplineHandler.GetAllDisciplineHandler)
	r.GET("/:id", disciplineHandler.GetAllDisciplinesByIdHandler)
	r.POST("/create", disciplineHandler.CreateDisciplineHandler)
	r.PUT("/edit/:id", disciplineHandler.EditDisciplineHandler)
	r.DELETE("/delete/:id", disciplineHandler.DeleteDisciplineHandler)
}

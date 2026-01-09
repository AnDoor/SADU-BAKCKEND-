package routes

import (
	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/handlers"
)

func RegisterUniversityRoutes(r *gin.RouterGroup, UniversityHandler *handlers.UniversityHandler) {
	r.GET("", UniversityHandler.GetAllUniversities)
	r.GET("/:id", UniversityHandler.GetUniversityByIdHandler)
	r.POST("/create", UniversityHandler.CreateUniversityHandler)
	r.PUT("/edit/:id", UniversityHandler.EditUniversityHandler)
	r.DELETE("/delete/:id", UniversityHandler.DeleteUniversityHandler)
}

package routes

import (
	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/handlers"
)

func RegisterMajorsRoutes(r *gin.RouterGroup, majorHandler *handlers.MajorHandler) {
	r.GET("",majorHandler.GetAllMajorHandler)
	r.GET("/:id", majorHandler.GetAllMajorByiDHandler)
	r.POST("/create",majorHandler.CreateMajorHandler)
	r.PUT("/edit/:id",majorHandler.EditMajorHandler)
	r.DELETE("/delete/:id", majorHandler.DeleteMajorHandler)
}

package routes

import (
	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/handlers"
)

func RegisterTeacherRoutes(r *gin.RouterGroup, teacherHandler *handlers.TeacherHandler) {
	r.GET("", teacherHandler.GetAllTeachersHandler)
	r.GET("/:id", teacherHandler.GetTeacherByIdHandler)
	r.POST("/create", teacherHandler.CreateTeacherHandler)
	r.PUT("/edit/:id", teacherHandler.UpdateTeacherHandler)
	r.DELETE("/delete/:id", teacherHandler.DeleteTeacherHandler)
}

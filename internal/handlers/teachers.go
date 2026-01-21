package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"uneg.edu.ve/servicio-sadu-back/helpers"
	"uneg.edu.ve/servicio-sadu-back/internal/services"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

type TeacherHandler struct {
	service *services.TeacherService
}

func NewTeacherHandler(service *services.TeacherService) *TeacherHandler {
	return &TeacherHandler{service: service}
}
func (h *TeacherHandler) GetAllTeachersHandler(c *gin.Context) {
	teachersDTO, err := h.service.GetTeachers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error listando profesores: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": teachersDTO, "message": "Profesores listados exitosamente"})
}

func (h *TeacherHandler) GetTeacherByIdHandler(c *gin.Context) {
	dto, err := h.service.GetTeacherById(c)
	if err != nil {

		if strings.Contains(err.Error(), "no encontrado") {

			helpers.SendError(c, 400, err.Error())
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo profesor: " + err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dto, "message": "Profesor encontrado exitosamente"})
}

func (h *TeacherHandler) CreateTeacherHandler(c *gin.Context) {
	var teacherDTO schema.TeacherCreateDTO
	if err := c.ShouldBindJSON(&teacherDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := h.service.CreateTeacher(teacherDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": dto, "message": "Profesor creado exitosamente"})
}

func (h *TeacherHandler) UpdateTeacherHandler(c *gin.Context) {
	var teacher schema.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teacher, err := h.service.EditTeacher(c, teacher)
	if err != nil {
		if strings.Contains(err.Error(), "no encontrado") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": teacher})
}

func (h *TeacherHandler) DeleteTeacherHandler(ctx *gin.Context) {
	if err := h.service.DeleteTeacher(ctx); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.SendError(ctx, http.StatusNotFound, "Profesor no encontrada")
		} else {
			helpers.SendError(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helpers.SendSucces(ctx, "Deleting Teacher succesfully", "")
}

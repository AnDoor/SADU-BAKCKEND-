package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/helpers"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

func GetTeachers(ctx *gin.Context) {
	teachers := []schema.Teacher{}

	if err := helpers.DB.Find(&teachers).Error; err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error listing teachers")
		return
	}

	teachersDTO := []schema.TeacherGetDTO{}
	for i, teacher := range teachers {
		var disciplines []schema.Discipline
		teacherID := schema.RegularIDs(teacher.ID)

		if err := helpers.DB.Where("id = ?", teacherID).Find(&disciplines).Error; err != nil {
			helpers.SendError(ctx, http.StatusInternalServerError, "Error finding teachers")
			return
		}
		teachers[i].Disciplines = disciplines

		teacherDTO := schema.TeacherGetDTO{
			ID:         teacherID,
			FirstNames: teacher.FirstNames,
			LastNames:  teacher.LastNames,
			PhoneNum:   teacher.PhoneNum,
			Email:      teacher.Email,
			GovID:      teacher.GovID,
			// Disciplines: disciplines,
		}
		teachersDTO = append(teachersDTO, teacherDTO)
	}
	helpers.SendSucces(ctx, "listing-teachers", teachersDTO)
}

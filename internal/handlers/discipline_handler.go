package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"uneg.edu.ve/servicio-sadu-back/helpers"
	"uneg.edu.ve/servicio-sadu-back/internal/services"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

type DisciplineHandler struct {
	service *services.DisciplineServices
}

func NewDisciplineHandler(service *services.DisciplineServices) *DisciplineHandler {
	return &DisciplineHandler{service: service}
}

func (h *DisciplineHandler) GetAllDisciplineHandler(ctx *gin.Context) {
	name := ctx.Query("name")
	discipline, err := h.service.GetAllDisciplines(name)
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error interno del servidor", "Ocurrió un problema inesperado al procesar la lista de disciplinas.")
		return
	}
	helpers.SendSucces(ctx, "LISTIN-DISCIPLINES-SUCCESFULLY", discipline)
}

func (h *DisciplineHandler) GetAllDisciplinesByIdHandler(ctx *gin.Context) {
	discipline, err := h.service.GetAllDisciplinesByID(ctx)
	if err != nil {
		helpers.SendError(ctx, http.StatusNotFound, "Error de busqueda en la base de datos", "El ID de la disciplina esta mal escrito o no existe")
		return
	}
	helpers.SendSucces(ctx, "LISTIN-DISCIPLINES-BY-ID-SUCCESFULLY", discipline)
}

func (h *DisciplineHandler) CreateDisciplineHandler(ctx *gin.Context) {
	var newDiscipline schema.Discipline
	if err := ctx.ShouldBindJSON(&newDiscipline); err != nil {
		helpers.SendError(ctx, http.StatusNotFound, "Error de busqueda en la base de datos", "Los datos de la disciplina no cargaron o no se encontraron en la Base de datos")
		return
	}

	createdDiscipline, err := h.service.CreateDiscipline(newDiscipline)
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error interno del servidor", "Datos incorrectos ingresado, la disciplina ya fue creado O problema inesperado")
		return
	}
	helpers.SendSucces(ctx, "CREATING-DISCIPLINE-SUCCESFULLY", createdDiscipline)
}

func (h *DisciplineHandler) EditDisciplineHandler(ctx *gin.Context) {
	var updateDiscipline schema.Discipline
	if err := ctx.ShouldBindJSON(&updateDiscipline); err != nil {
		helpers.SendError(ctx, http.StatusNotFound, "Error de busqueda en la base de datos", "Los datos de la disciplina no cargaron o no se encontraron en la Base de datos")
		return
	}

	createdDiscipline, err := h.service.EditDiscipline(updateDiscipline, ctx)
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error interno del servidor", "Datos incorrectos ingresado o la disciplina no fue encontrada")
		return
	}
	helpers.SendSucces(ctx, "EDITING-DISCIPLINE-SUCCESFULLY", createdDiscipline)
}

func (h *DisciplineHandler) DeleteDisciplineHandler(ctx *gin.Context) {
	if err := h.service.DeleteDiscipline(ctx); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.SendError(ctx, http.StatusNotFound, "Error de busqueda en la base de datos", "Los datos de la disciplina no cargaron o no se encontraron en la Base de datos")
				return
		} 
	}

	helpers.SendSucces(ctx, "DISCIPLINE-DELETED-SUCCESFULLY", "")

}

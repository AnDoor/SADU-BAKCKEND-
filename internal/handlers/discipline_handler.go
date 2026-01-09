package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	discipline, err := h.service.GetAllDisciplines()
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "ERROR IN HANDLER\n Error listing Disciplines")
		return
	}
	sendSucces(ctx, "LISTIN DISCIPLINES SUCCESFULLY", discipline)
}

func (h *DisciplineHandler) GetAllDisciplinesByIdHandler(ctx *gin.Context) {
	discipline, err := h.service.GetAllDisciplinesByID(ctx)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "ERROR IN HANDLER\n Error listing Disciplines")
		return
	}
	sendSucces(ctx, "LISTIN DISCIPLINES BY ID SUCCESFULLY", discipline)
}

func (h *DisciplineHandler) CreateDisciplineHandler(ctx *gin.Context) {
	var newDiscipline schema.Discipline
	if err := ctx.ShouldBindJSON(&newDiscipline); err != nil {
		sendError(ctx, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	createdDiscipline, err := h.service.CreateDiscipline(newDiscipline)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	sendSucces(ctx, "CREATING-DISCIPLINE-SUCCESFULLY", createdDiscipline)
}

func (h *DisciplineHandler) EditDisciplineHandler(ctx *gin.Context) {
	var updateDiscipline schema.Discipline
	if err := ctx.ShouldBindJSON(&updateDiscipline); err != nil {
		sendError(ctx, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	createdDiscipline, err := h.service.EditDiscipline(updateDiscipline, ctx)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	sendSucces(ctx, "EDITING-DISCIPLINE-SUCCESFULLY", createdDiscipline)
}

func (h *DisciplineHandler) DeleteDisciplineHandler(ctx *gin.Context) {
	if err := h.service.DeleteDiscipline(ctx); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			sendError(ctx, http.StatusNotFound, "Discipline not found")
		} else {
			sendError(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}

	sendSucces(ctx, "Deleting university succesfully", "")

}

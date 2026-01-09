package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"uneg.edu.ve/servicio-sadu-back/internal/services"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

type UniversityHandler struct {
	service *services.UniversityServices
}

func NewUniversityHandler(service *services.UniversityServices) *UniversityHandler {
	return &UniversityHandler{service: service}
}

func (h *UniversityHandler) GetAllUniversities(ctx *gin.Context) {
	universities, err := h.service.GetAllUniversity()
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "ERROR IN HANDLER\n Error listing Universities")
	}
	sendSucces(ctx, "LISTING-UNIVERSITIES", universities)
}

func (h *UniversityHandler) GetUniversityByIdHandler(ctx *gin.Context) {
	universities, err := h.service.GetUniversityByID(ctx)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "ERROR IN HANDLER\n Error listing by ID Universities")
		return
	}
	sendSucces(ctx, "LISTING-UNIVERSITIES-BY-ID", universities)
}

func (h *UniversityHandler) CreateUniversityHandler(ctx *gin.Context) {
	var newUniversirty schema.University
	if err := ctx.ShouldBindJSON(&newUniversirty); err != nil {
		sendError(ctx, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	createdUniversity, err := h.service.CreateUniversity(newUniversirty)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	sendSucces(ctx, "CREATING-UNIVERSITY-SUCCESFULLY", createdUniversity)

}
func (h *UniversityHandler) EditUniversityHandler(ctx *gin.Context) {
	var university schema.University
	if err := ctx.ShouldBindJSON(&university); err != nil {
		sendError(ctx, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	updatedUniversity, err := h.service.EditUniversity(university, ctx)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, " "+err.Error())
		return
	}
	sendSucces(ctx, "UPDATING-UNIVERSITY-SUCCESFULLY", updatedUniversity)
}

func (h *UniversityHandler) DeleteUniversityHandler(ctx *gin.Context) {
	if err := h.service.DeleteUniversity(ctx); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			sendError(ctx, http.StatusNotFound, "Universidad no encontrada")
		} else {
			sendError(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}

	sendSucces(ctx, "Deleting university succesfully", "")
}

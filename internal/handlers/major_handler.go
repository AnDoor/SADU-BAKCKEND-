package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/services"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

type MajorHandler struct {
	service *services.MajorServices
}

func NewMajorHandler(service *services.MajorServices) *MajorHandler {
	return &MajorHandler{service: service}
}

func (h *MajorHandler) GetAllMajorHandler(ctx *gin.Context) {
	majors, err := h.service.GetAllMajor()
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "ERROR IN HANDLER\nError listin majors")
		return
	}
	sendSucces(ctx, "", majors)
}

func (h *MajorHandler) GetAllMajorByiDHandler(ctx *gin.Context) {
	majors, err := h.service.GetMajorByID(ctx)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "ERROR IN HANDLER\nError listin majors  by ID")
		return
	}
	sendSucces(ctx, "", majors)
}
func (h *MajorHandler) CreateMajorHandler(ctx *gin.Context) {

	var newMajor schema.Major
	if err := ctx.ShouldBindJSON(&newMajor); err != nil {
		sendError(ctx, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	createdMajor, err := h.service.CreateMajor(newMajor)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error creando carreras: "+err.Error())
		return
	}

	sendSucces(ctx, "Atleta creado exitosamente", createdMajor)
}

func (h *MajorHandler) EditMajorHandler(ctx *gin.Context) {

	var newMajor schema.Major
	if err := ctx.ShouldBindJSON(&newMajor); err != nil {
		sendError(ctx, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	updatedMajor, err := h.service.CreateMajor(newMajor)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error creando carreras: "+err.Error())
		return
	}

	sendSucces(ctx, "Atleta creado exitosamente", updatedMajor)
}

func (h *MajorHandler) DeleteMajorHandler(ctx *gin.Context) {

	if err := h.service.DeleteMajor(ctx); err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error eliminando carreras: "+err.Error())
		return
	}

	sendSucces(ctx, "Deleting-major", "")
}

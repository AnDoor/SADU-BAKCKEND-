package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/helpers"
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
	name := ctx.Query("name")
	majors, err := h.service.GetAllMajor(name)
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error interno del servidor", "Ocurrió un problema inesperado al procesar la lista de carreras universitarias.")
		return
	}
	helpers.SendSucces(ctx, "Listing-Majors-Succesfully", majors)
}

func (h *MajorHandler) GetAllMajorByiDHandler(ctx *gin.Context) {
	majors, err := h.service.GetMajorByID(ctx)
	if err != nil {
		helpers.SendError(ctx, http.StatusNotFound, "Error de busqueda en la base de datos", "El ID de la carrera universitaria esta mal escrito o no existe.")
		return
	}
	helpers.SendSucces(ctx, "Listing-Major-By-ID-Succesfully", majors)
}

func (h *MajorHandler) CreateMajorHandler(ctx *gin.Context) {

	var newMajor schema.Major
	if err := ctx.ShouldBindJSON(&newMajor); err != nil {
		helpers.SendError(ctx, http.StatusNotFound, "Error de busqueda en la base de datos", "Los datos de la carrera universitaria no cargaron o no se encontraron en la Base de datos.")
		return
	}

	createdMajor, err := h.service.CreateMajor(newMajor)
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error interno del servidor", "Datos incorrectos ingresados en la creacion de la carrera universitaria.")
		return
	}

	helpers.SendSucces(ctx, "Carrera universitaria creada exitosamente.", createdMajor)
}

func (h *MajorHandler) EditMajorHandler(ctx *gin.Context) {

	var newMajor schema.Major
	if err := ctx.ShouldBindJSON(&newMajor); err != nil {
		helpers.SendError(ctx, http.StatusNotFound, "Error de busqueda en la base de datos", "Los datos de la carrera universitaria no se cargaron o no se encontraron en la Base de datos.")
		return
	}

	updatedMajor, err := h.service.CreateMajor(newMajor)
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error interno del servidor", "Datos incorrectos ingresados en la edicion de la carrera universitaria.")
		return
	}

	helpers.SendSucces(ctx, "Carrera universitaria editada exitosamente.", updatedMajor)
}

func (h *MajorHandler) DeleteMajorHandler(ctx *gin.Context) {

	if err := h.service.DeleteMajor(ctx); err != nil {
		helpers.SendError(ctx, http.StatusNotFound, "Error de busqueda en la base de datos", "La carrera universitaria ya fue eliminada o el ID de la carrera universitaria no se encuentra.")
		return
	}

	helpers.SendSucces(ctx, "Deleting-major", "")
}

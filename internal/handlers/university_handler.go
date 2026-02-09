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

type UniversityHandler struct {
	service *services.UniversityServices
}

func NewUniversityHandler(service *services.UniversityServices) *UniversityHandler {
	return &UniversityHandler{service: service}
}

func (h *UniversityHandler) GetAllUniversities(ctx *gin.Context) {
	name := ctx.Query("name")
	local := ctx.Query("local")
	universities, err := h.service.GetAllUniversity(name,local)
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "ERROR: listing Universities","un error inesperado ha ocurrido dentro del enlistado de universidades")
		return
	}
	helpers.SendSucces(ctx, "LISTING-UNIVERSITIES", universities)
}

func (h *UniversityHandler) GetUniversityByIdHandler(ctx *gin.Context) {
	universities, err := h.service.GetUniversityByID(ctx)
	if err != nil {
		helpers.SendError(ctx, http.StatusNotFound, "ERROR: listing by ID Universities","el ID de la universidad ingresado esta mal escrito o fue eliminado")
		return
	}
	helpers.SendSucces(ctx, "LISTING-UNIVERSITIES-BY-ID", universities)
}

func (h *UniversityHandler) CreateUniversityHandler(ctx *gin.Context) {
	var newUniversirty schema.University
	if err := ctx.ShouldBindJSON(&newUniversirty); err != nil {
		helpers.SendError(ctx, http.StatusNotFound, "JSON inválido: "+ err.Error(), "La universidad no se encuentra o un dato incorrecto fue ingresado a la hora de crear una universidad")
		return
	}

	createdUniversity, err := h.service.CreateUniversity(newUniversirty)
	if err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, err.Error(),"Error en la creacion de universidad inesperado")
		return
	}
	helpers.SendSucces(ctx, "CREATING-UNIVERSITY-SUCCESFULLY", createdUniversity)

}
func (h *UniversityHandler) EditUniversityHandler(ctx *gin.Context) {
	var university schema.University
	if err := ctx.ShouldBindJSON(&university); err != nil {
		helpers.SendError(ctx, http.StatusNotFound, "JSON inválido: " +err.Error(),"La universidad a editar fue eliminada o no se en encuentra creada en la base de datos")
		return
	}
	updatedUniversity, err := h.service.EditUniversity(university, ctx)
	if err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "Error editing"+ err.Error(),"Error en la edicion de universidad inesperado")
		return
	}
	helpers.SendSucces(ctx, "UPDATING-UNIVERSITY-SUCCESFULLY", updatedUniversity)
}

func (h *UniversityHandler) DeleteUniversityHandler(ctx *gin.Context) {
	if err := h.service.DeleteUniversity(ctx); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.SendError(ctx, http.StatusNotFound, "Universidad no encontrada","El ID ingresado para buscar la universidad no existe o fue ingresado de manera incorrecto")
		} else {
			helpers.SendError(ctx, http.StatusInternalServerError, err.Error(),"Error inesperado a la hora de eliminar la universidad")
		}
		return
	}

	helpers.SendSucces(ctx, "Deleting university succesfully", "")
}

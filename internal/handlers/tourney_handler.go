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

type TourneyHandler struct {
	service *services.TourneyServices
}

func NewTourneyHandler(service *services.TourneyServices) *TourneyHandler {
	return &TourneyHandler{service: service}
}

func (h *TourneyHandler) GetAllTourneyHandler(ctx *gin.Context) {
	name := ctx.Query("name")
	status := ctx.Query("status")

	dtos, err := h.service.GetAllTourney(name, status)
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error interno del servidor", "Ocurrió un problema inesperado al procesar la lista de torneos.")
		return
	}

	helpers.SendSucces(ctx, "Listing-Tourneys-Succesfully", dtos)
}

func (h *TourneyHandler) GetTourneyByIdHandler(ctx *gin.Context) {
	tourney, err := h.service.GetTourneyByID(ctx)
	if err != nil {
		helpers.SendError(ctx, http.StatusNotFound, "Error de busqueda", "El ID del equipo esta mal escrito o no se encuentra en la base de datos.")
		return
	}
	helpers.SendSucces(ctx, "Listing-Tourneys-By-ID-Succesfully", tourney)
}

func (h *TourneyHandler) CreateUniversityHandler(ctx *gin.Context) {
	var newTourney schema.Tourney
	if err := ctx.ShouldBindJSON(&newTourney); err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "JSON inválido: "+ err.Error(),"El torneo ya fue creado anteriormente o no fue encontrado.")
		return
	}

	createdTourney, err := h.service.CreateTourney(newTourney)
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error interno del servidor", "El inesperado a la hora de crear torneo.")
		return
	}
	helpers.SendSucces(ctx, "CREATING-TOURNEY-SUCCESFULLY", createdTourney)

}

func (h *TourneyHandler) UpdateTourneyHandler(ctx *gin.Context) {

	var t schema.Tourney
	if err := ctx.ShouldBindJSON(&t); err != nil {
	helpers.SendError(ctx, http.StatusNotFound, "Error interno del servidor", "El torneo no fue encontrado en la base de datos.")
		return
	}

	dto, err := h.service.UpdateTourney(t, ctx)
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error interno del servidor", "El inesperado a la hora de editar torneo.")
		return
	
	}
	helpers.SendSucces(ctx, "EDITING-TOURNEY-SUCCESFULLY", dto)

}

func (h *TourneyHandler) DeleteTourneyHandler(ctx *gin.Context) {
	if err := h.service.DeleteTourney(ctx); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.SendError(ctx, http.StatusNotFound, "Torneo no encontrada", "El ID del torneo esta mal escrito o es invalido para buscar el torneo en la base de datos")
		} else {
			helpers.SendError(ctx, http.StatusInternalServerError, err.Error(),"Error interno en el servidor: no se encuentra torneo para eliminar")
		}
		return
	}

	helpers.SendSucces(ctx, "Deleting Tourney succesfully", "")
}

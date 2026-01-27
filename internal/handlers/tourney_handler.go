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

func (h *TourneyHandler) GetAllTourneyHandler(c *gin.Context) {
		name := c.Query("name")
		status := c.Query("status")

	dtos, err := h.service.GetAllTourney(name,status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener torneos"})
		return
	}

	c.JSON(http.StatusOK, dtos)
}

func (h *TourneyHandler) GetTourneyByIdHandler(ctx *gin.Context) {
	tourney, err := h.service.GetTourneyByID(ctx)
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "ERROR IN HANDLER\n Error listing by ID TOURNEYS")
		return
	}
	helpers.SendSucces(ctx, "LISTING-TOURNEY-BY-ID", tourney)
}

func (h *TourneyHandler) CreateUniversityHandler(ctx *gin.Context) {
	var newTourney schema.Tourney
	if err := ctx.ShouldBindJSON(&newTourney); err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	createdTourney, err := h.service.CreateTourney(newTourney)
	if err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	helpers.SendSucces(ctx, "CREATING-TOURNEY-SUCCESFULLY", createdTourney)

}

func (h *TourneyHandler) UpdateTourneyHandler(c *gin.Context) {

	var t schema.Tourney
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := h.service.UpdateTourney(t, c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dto})

}

func (h *TourneyHandler) DeleteTourneyHandler(ctx *gin.Context) {
	if err := h.service.DeleteTourney(ctx); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.SendError(ctx, http.StatusNotFound, "Torneo no encontrada")
		} else {
			helpers.SendError(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helpers.SendSucces(ctx, "Deleting Tourney succesfully", "")
}

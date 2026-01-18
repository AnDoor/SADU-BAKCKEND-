package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/services"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

type TourneyHandler struct {
	service *services.TourneyServices
}

func NewTourneyHandler(service *services.TourneyServices) *TourneyHandler {
	return &TourneyHandler{service: service}
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

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/services"
)

type AthleteHandler struct {
    service *services.AthleteService
}

func NewAthleteHandler(service *services.AthleteService) *AthleteHandler {
    return &AthleteHandler{service: service}
}

func (h *AthleteHandler) GetAthletes(ctx *gin.Context) {
	athletes, err := h.service.GetAllAthletes()
	if err != nil{
		sendError(ctx, http.StatusInternalServerError, "ERROR IN HANDLER\nError listin athletes")
	}
	sendSucces(ctx, "listing-athletes", athletes)
}

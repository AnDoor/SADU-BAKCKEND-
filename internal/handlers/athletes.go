package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/services"
	"uneg.edu.ve/servicio-sadu-back/schema"
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

func (h *AthleteHandler) GetAthletesByID(ctx *gin.Context) {
	athletes, err := h.service.GetAthletesByID(ctx)
	if err != nil{
		sendError(ctx, http.StatusInternalServerError, "ERROR IN HANDLER\nError listin athletes by ID")
	}
	sendSucces(ctx, "listing-athletes-by-id", athletes)
}
func (h *AthleteHandler) CreateNewAthlete(ctx *gin.Context) {
    // 1. BIND JSON → OBTENER DATOS del request
    var newAthlete schema.Athlete
    if err := ctx.ShouldBindJSON(&newAthlete); err != nil {
        sendError(ctx, http.StatusBadRequest, "JSON inválido: "+err.Error())
        return
    }

    createdAthlete, err := h.service.CreateAthlete(newAthlete)
    if err != nil {
        sendError(ctx, http.StatusInternalServerError, "Error creando atleta: "+err.Error())
        return
    }

    sendSucces(ctx, "Atleta creado exitosamente", createdAthlete)
}
func (h *AthleteHandler) EditAthleteByID(ctx *gin.Context) {

	var athlete schema.Athlete
	if err:= ctx.ShouldBindJSON(&athlete); err != nil {
		sendError(ctx,http.StatusBadRequest, "JSON INVALIDO" + err.Error())
		return
	}
	updateAthlete,err := h.service.EditAthlete(athlete,ctx)
	if err != nil {
		sendError(ctx, http.StatusNotFound,err.Error())
		return
	}
	sendSucces(ctx, "updated Athletes succesfully",updateAthlete)

}

func (h *AthleteHandler) DeleteAthleteByID(ctx *gin.Context) {

	if err := h.service.DeleteAthlete(ctx); err != nil{
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	
	sendSucces(ctx, "Deleting-athlete", "")
}
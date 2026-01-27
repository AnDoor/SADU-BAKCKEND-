package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/helpers"
	"uneg.edu.ve/servicio-sadu-back/internal/services"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

type EventHandler struct {
	service *services.EventService
}

func NewEventHandler(service *services.EventService) *EventHandler {
	return &EventHandler{service: service}
}

func (h *EventHandler) GetAllEventsHandler(c *gin.Context) {
	name := c.Query("name")
	status := c.Query("status")
	events, err := h.service.GetAllEvents(name, status)
	if err != nil {
		helpers.SendError(c, http.StatusInternalServerError, "ERROR IN HANDLER: Error listing Events")
	}
	helpers.SendSucces(c, "LISTIN EVENTS SUCCESFULLY", events)
}

func (h *EventHandler) GetEventByIDHandler(ctx *gin.Context) {
	events, err := h.service.GetEventByID(ctx)
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "ERROR IN HANDLER: Error listin Event  by ID")
		return
	}
	helpers.SendSucces(ctx, "", events)
}
func (h *EventHandler) CreateEventHandler(ctx *gin.Context) {

	var newEvent schema.Event
	if err := ctx.ShouldBindJSON(&newEvent); err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	createdEvent, err := h.service.CreateEvent(newEvent)
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error creando Evento: "+err.Error())
		return
	}

	helpers.SendSucces(ctx, "Evento creado exitosamente", createdEvent)
}

func (h *EventHandler) EditEventHandler(ctx *gin.Context) {

	var newEvent schema.Event
	if err := ctx.ShouldBindJSON(&newEvent); err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	updatedEvent, err := h.service.CreateEvent(newEvent)
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error creando Eventos: "+err.Error())
		return
	}

	helpers.SendSucces(ctx, "Evento creado exitosamente", updatedEvent)
}

func (h *EventHandler) DeleteEventHandler(ctx *gin.Context) {

	if err := h.service.DeleteEvent(ctx); err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error eliminando Evento: "+err.Error())
		return
	}

	helpers.SendSucces(ctx, "Deleting-Event", "")
}

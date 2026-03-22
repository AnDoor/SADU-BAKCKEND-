package handlers

import (
	"net/http"
	"strconv"

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

func (h *EventHandler) GetEventsHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id uint
	if idParam != "" {
		val, _ := strconv.Atoi(idParam)
		id = uint(val)
	}

	name := ctx.Query("name")
	status := ctx.Query("status")

	events, err := h.service.GetEvents(id,name, status)

	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error interno del servidor", "Ocurrió un problema inesperado al procesar la lista de Eventos.")
		return
	}

	helpers.SendSucces(ctx, "LISTIN-EVENTS-SUCCESFULLY", events)
}

func (h *EventHandler) CreateEventHandler(ctx *gin.Context) {

	var newEvent schema.Event
	if err := ctx.ShouldBindJSON(&newEvent); err != nil {
		helpers.SendError(ctx, http.StatusNotFound, "Error de busqueda en la base de datos", "Los datos del evento no cargaron o no se encontraron en la Base de datos")
		return
	}

	createdEvent, err := h.service.CreateEvent(newEvent)
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error interno del servidor", "Datos incorrectos ingresados en la creacion de atletas")
		return
	}

	helpers.SendSucces(ctx, "Evento creado exitosamente", createdEvent)
}

func (h *EventHandler) EditEventHandler(ctx *gin.Context) {

	var newEvent schema.Event
	if err := ctx.ShouldBindJSON(&newEvent); err != nil {
		helpers.SendError(ctx, http.StatusNotFound, "Error de busqueda en la base de datos", "Los datos del evento no cargaron o no se encontraron en la Base de datos")
		return
	}

	updatedEvent, err := h.service.CreateEvent(newEvent)
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error interno del servidor", "Datos incorrectos ingresados en la edicion de atletas")
		return
	}

	helpers.SendSucces(ctx, "Evento creado exitosamente", updatedEvent)
}

func (h *EventHandler) DeleteEventHandler(ctx *gin.Context) {

	if err := h.service.DeleteEvent(ctx); err != nil {
		helpers.SendError(ctx, http.StatusNotFound, "Error de busqueda en la base de datos", "Los datos del atleta no cargaron y no se encontraron en la Base de datos para eliminarlo")
		return
	}

	helpers.SendSucces(ctx, "Deleting-Event", "")
}

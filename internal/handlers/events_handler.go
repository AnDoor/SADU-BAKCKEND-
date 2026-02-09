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

func (h *EventHandler) GetAllEventsHandler(ctx *gin.Context) {
	name := ctx.Query("name")
	status := ctx.Query("status")

	events, err := h.service.GetAllEvents(name, status)

	if err != nil {
	helpers.SendError(ctx, http.StatusInternalServerError, "Error interno del servidor", "Ocurrió un problema inesperado al procesar la lista de Eventos.")
	return
}

	helpers.SendSucces(ctx, "LISTIN-EVENTS-SUCCESFULLY", events)
}

func (h *EventHandler) GetEventByIDHandler(ctx *gin.Context) {
	events, err := h.service.GetEventByID(ctx)
	if err != nil {
		helpers.SendError(ctx, http.StatusNotFound, "Error de busqueda en la base de datos", "El ID del evento es incorrecto, no se encontro Evento.")
		return
	}
	helpers.SendSucces(ctx, "LISTIN-EVENTS-BY-ID-SUCCESFULLY", events)
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

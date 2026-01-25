package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/internal/services"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

type TeamHandler struct {
	service *services.TeamServices
}

func NewTeamHandler(service *services.TeamServices) *TeamHandler {
	return &TeamHandler{service: service}
}

func (h *TeamHandler) GetAllTeamHandler(c *gin.Context) {
	team, err := h.service.GetAllTeam()
	if err != nil {

		if strings.Contains(err.Error(), "team") && strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo equipos: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": team})
}


func (h *TeamHandler) GetTeamByIdHandler(c *gin.Context) {
	team, err := h.service.GetAllTeamByID(c)
	if err != nil {

		if strings.Contains(err.Error(), "invalid team ID") {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if strings.Contains(err.Error(), "team") && strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo equipo: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": team})
}

func (h *TeamHandler) CreateTeamHandler(c *gin.Context) {
	var input schema.TeamPostDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	team, err := h.service.CreateTeam(input)
	if err != nil {
		
		if strings.Contains(err.Error(), "discipline") && strings.Contains(err.Error(), "not found") ||
			strings.Contains(err.Error(), "university") && strings.Contains(err.Error(), "not found") ||
			strings.Contains(err.Error(), "athlete") && strings.Contains(err.Error(), "not found") ||
			strings.Contains(err.Error(), "creating team") {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creando equipo: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": team})
}


func (h *TeamHandler) DeleteTeam(c *gin.Context) {
    err := h.service.DeleteTeam(c)
    if err != nil {

        if strings.Contains(err.Error(), "invalid team ID") {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if strings.Contains(err.Error(), "team") && strings.Contains(err.Error(), "not found") ||
           strings.Contains(err.Error(), "already deleted") {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        if strings.Contains(err.Error(), "deleting team") {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error eliminando equipo: " + err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}

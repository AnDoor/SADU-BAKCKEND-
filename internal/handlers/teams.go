package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

func GetTeams(ctx *gin.Context) {
	teams := []schema.Team{}

	if err := DB.Preload("Athletes").Find(&teams).Error; err != nil {
		println(err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error listing teams")
		return
	}

	teamsDTO := []schema.TeamGetDTO{}
	for _, team := range teams {
		teamsDTO = append(teamsDTO, schema.TeamGetDTO{
			ID:      schema.RegularIDs(team.ID),
			Name:    team.Name,
			Regular: team.Regular,
			// Category:     team.Category,
			DisciplineID: team.DisciplineID,
			UniversityID: team.UniversityID,
			Athletes:     team.Athletes,
		})
	}

	sendSucces(ctx, "listing-teams", teamsDTO)
	return
}

func AddTeams(ctx *gin.Context) {
	teamDTO := schema.TeamPostDTO{}
	println(ctx.Request.Header)

	if err := ctx.BindJSON(&teamDTO); err != nil {
		println(err.Error())
		sendError(ctx, http.StatusBadRequest, "Error binding json")
		return
	}

	athletes := []schema.Athlete{}
	for _, athleteID := range teamDTO.AthleteIDs {
		athlete := schema.Athlete{}
		if err := DB.Where("id = ?", athleteID).First(&athlete).Error; err != nil {
			sendError(ctx, http.StatusInternalServerError, "Error finding athletes")
			return
		}
		athletes = append(athletes, athlete)
	}

	team := schema.Team{
		Name:    teamDTO.Name,
		Regular: teamDTO.Regular,
		// Category:     teamDTO.Category,
		DisciplineID: teamDTO.DisciplineID,
		UniversityID: teamDTO.UniversityID,
		Athletes:     athletes,
	}

	if err := DB.Create(&team).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error inserting team")
		return
	}
	sendSucces(ctx, "adding-team", team)
}

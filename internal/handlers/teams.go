package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/helpers"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

func GetTeams(ctx *gin.Context) {
	teams := []schema.Team{}

	if err := helpers.DB.Preload("Athletes").Find(&teams).Error; err != nil {
		println(err.Error())
		helpers.SendError(ctx, http.StatusInternalServerError, "Error listing teams")
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

	helpers.SendSucces(ctx, "listing-teams", teamsDTO)
	return
}

func AddTeams(ctx *gin.Context) {
	teamDTO := schema.TeamPostDTO{}
	println(ctx.Request.Header)

	if err := ctx.BindJSON(&teamDTO); err != nil {
		println(err.Error())
		helpers.SendError(ctx, http.StatusBadRequest, "Error binding json")
		return
	}

	athletes := []schema.Athlete{}
	for _, athleteID := range teamDTO.AthleteIDs {
		athlete := schema.Athlete{}
		if err := helpers.DB.Where("id = ?", athleteID).First(&athlete).Error; err != nil {
			helpers.SendError(ctx, http.StatusInternalServerError, "Error finding athletes")
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

	if err := helpers.DB.Create(&team).Error; err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error inserting team")
		return
	}
	helpers.SendSucces(ctx, "adding-team", team)
}

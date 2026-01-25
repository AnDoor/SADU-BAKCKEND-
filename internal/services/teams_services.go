package services

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"uneg.edu.ve/servicio-sadu-back/config"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

type TeamServices struct {
	DB *gorm.DB
}

func NewTeamServices() *TeamServices {
	return &TeamServices{DB: config.DB}
}

func (s *TeamServices) GetAllTeam() ([]schema.TeamGetDTO, error) {
	var teams []schema.Team
	if err := s.DB.Preload("Athletes").Find(&teams).Error; err != nil {
		return nil, fmt.Errorf("listing teams: %w", err)
	}

	// Mapear a DTOs
	dtos := make([]schema.TeamGetDTO, len(teams))
	for i, team := range teams {
		dtos[i] = schema.TeamGetDTO{
			ID:           schema.RegularIDs(team.ID),
			Name:         team.Name,
			Regular:      team.Regular,
			Category:     string(team.Category),
			DisciplineID: schema.RegularIDs(team.DisciplineID),
			UniversityID: schema.RegularIDs(team.UniversityID),
			Athletes:     team.Athletes,
		}
	}
	return dtos, nil
}

func (s *TeamServices) GetAllTeamByID(ctx *gin.Context) (schema.TeamGetBareDTO, error) {
	id := ctx.Param("id")
	teamID, err := strconv.Atoi(id)
	if err != nil {
		return schema.TeamGetBareDTO{}, fmt.Errorf("invalid team ID: %w", err)
	}

	var team schema.Team
	if err := s.DB.Preload("University").Preload("Athletes").First(&team, teamID).Error; err != nil {
		return schema.TeamGetBareDTO{}, fmt.Errorf("team %d not found: %w", teamID, err)
	}

	return schema.TeamGetBareDTO{
		ID:       schema.RegularIDs(team.ID),
		Name:     team.Name,
		Regular:  team.Regular,
		Category: string(team.Category),
		University: schema.UniversityGetBareDTO{
			ID:   schema.RegularIDs(team.University.ID),
			Name: team.University.Name,
		},
		Athletes: team.Athletes,
	}, nil
}

func (s *TeamServices) CreateTeam(t schema.TeamPostDTO) (schema.TeamGetBareDTO, error) {
	// 1. Validar referencias FK
	var discipline schema.Discipline
	if err := s.DB.First(&discipline, uint(t.DisciplineID)).Error; err != nil {
		return schema.TeamGetBareDTO{}, fmt.Errorf("discipline %d not found: %w", t.DisciplineID, err)
	}

	var university schema.University
	if err := s.DB.First(&university, uint(t.UniversityID)).Error; err != nil {
		return schema.TeamGetBareDTO{}, fmt.Errorf("university %d not found: %w", t.UniversityID, err)
	}

	team := schema.Team{
		Name:         t.Name,
		Regular:      t.Regular,
		Category:     schema.Gender(t.Category),
		DisciplineID: t.DisciplineID,
		UniversityID: t.UniversityID,
	}

	if err := s.DB.Create(&team).Error; err != nil {
		return schema.TeamGetBareDTO{}, fmt.Errorf("creating team: %w", err)
	}

	// 4. Asignar atletas (many2many)
	for _, athleteID := range t.AthleteIDs {
		var athlete schema.Athlete
        if err := s.DB.First(&athlete, uint(athleteID)).Error; err != nil {
            return schema.TeamGetBareDTO{}, fmt.Errorf("athlete %d not found: %w", athleteID, err)
        }
    }

	//  Retornar DTO completo con relaciones
	var fullTeam schema.Team
	if err := s.DB.Preload("University").Preload("Athletes").First(&fullTeam, team.ID).Error; err != nil {
		return schema.TeamGetBareDTO{}, fmt.Errorf("loading created team: %w", err)
	}

	dto := schema.TeamGetBareDTO{
		ID:       schema.RegularIDs(fullTeam.ID),
		Name:     fullTeam.Name,
		Regular:  fullTeam.Regular,
		Category: string(fullTeam.Category),
		University: schema.UniversityGetBareDTO{
			ID:   schema.RegularIDs(fullTeam.University.ID),
			Name: fullTeam.University.Name,
		},
		Athletes: fullTeam.Athletes,
	}

	return dto, nil
}
//func (s *TeamServices) EditTeam(t schema.TeamUpdateDTO, ctx *gin.Context) (schema.TeamGetBareDTO, error)

func (s *TeamServices) DeleteTeam(ctx *gin.Context) error {
	id := ctx.Param("id")
	teamID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid team ID %q: %w", id, err)
	}

	var team schema.Team
	if err := s.DB.First(&team, teamID).Error; err != nil {
		return fmt.Errorf("team %d not found: %w", teamID, err)
	}

	result := s.DB.Delete(&team, teamID)
	if result.Error != nil {
		return fmt.Errorf("deleting team %d: %w", teamID, result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("team %d not found or already deleted", teamID)
	}

	return nil
}

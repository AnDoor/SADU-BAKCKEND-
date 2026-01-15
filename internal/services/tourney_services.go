package services

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"uneg.edu.ve/servicio-sadu-back/config"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

type TourneyServices struct {
	DB *gorm.DB
}

func NewTourneyServices() *TourneyServices {
	return &TourneyServices{DB: config.DB}
}

func mapEventsBare(events []schema.Event) []schema.EventGetBareDTO {
    dtos := make([]schema.EventGetBareDTO, len(events))
    for i, event := range events {
        dtos[i] = schema.EventGetBareDTO{
            ID:             schema.RegularIDs(event.ID),
            Name:           event.Name,
            Date:           event.Date,
            Status:         event.Status,
            HomePoints:     event.HomePoints,
            OppositePoints: event.OppositePoints,

        }
    }
    return dtos
}

func (s *TourneyServices) GetAllTourney() ([]schema.TourneyGetBareDTO, error) {
	var tourneys []schema.Tourney
	if err := s.DB.Preload("Events", nil).Find(&tourneys).Error; err != nil {
		return nil, err
	}
	var dtos []schema.TourneyGetBareDTO

	for _, t := range tourneys {
		dto := schema.TourneyGetBareDTO{
			ID:   schema.RegularIDs(t.ID),
			Name: t.Name,
		}
		dtos = append(dtos, dto)
	}
	return dtos, nil
}
func (s *TourneyServices) GetTourneyByID(c *gin.Context) (schema.TourneyGetFullDTO, error) {
	var id = c.Param("id")
	var tourney schema.Tourney
	tourneyID, err := strconv.Atoi(id)
	if err != nil {
		return schema.TourneyGetFullDTO{}, fmt.Errorf("ID INVALID:%w", err)

	}
	result := s.DB.Preload("Events").First(&tourney, tourneyID).Error
	if result.Error != nil {
		return schema.TourneyGetFullDTO{}, result
	}

	return schema.TourneyGetFullDTO{
		ID:     schema.RegularIDs(tourneyID),
		Name:   tourney.Name,
		 Events: mapEventsBare(tourney.Events),
	}, nil
}
func (s *TourneyServices) CreateTourney(t schema.Tourney) (schema.TourneyGetFullDTO, error)
func (s *TourneyServices) UpdateTourney(t schema.Tourney, c *gin.Context) (schema.TourneyGetFullDTO, error)

func (s *TourneyServices) DeleteTourney(c *gin.Context) error {
	var id = c.Param("id")
	tourneyID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("Invalid ID")
	}
	var result = s.DB.Delete(&schema.Tourney{}, tourneyID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("No se ha encontrado Torneo")
	}

	return nil
}

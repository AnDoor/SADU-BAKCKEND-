package services

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"uneg.edu.ve/servicio-sadu-back/config"
	"uneg.edu.ve/servicio-sadu-back/helpers"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

type TourneyServices struct {
	DB *gorm.DB
}

func NewTourneyServices() *TourneyServices {
	return &TourneyServices{DB: config.DB}
}

func (s *TourneyServices) GetAllTourney(name, status string) ([]schema.TourneyGetBareDTO, error) {
	var tourneys []schema.Tourney
	query := s.DB.Preload("Events")
   
    if name != "" {
        query = query.Where("name LIKE ?", "%"+name+"%")
    } 
	if status != "" {
        query = query.Where("status LIKE ?", "%"+status+"%")
    } 

    if err := query.Find(&tourneys).Error; err != nil {
        return nil, err
    }
	var dtos []schema.TourneyGetBareDTO

	for _, t := range tourneys {
		dto := schema.TourneyGetBareDTO{
			ID:   schema.RegularIDs(t.ID),
			Name: t.Name,
			Status: t.Status,
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
	if result != nil {
		return schema.TourneyGetFullDTO{}, result
	}

	return schema.TourneyGetFullDTO{
		ID:     schema.RegularIDs(tourneyID),
		Name:   tourney.Name,
		Status: tourney.Status,
		Events: helpers.MapEventsBare(tourney.Events),
	}, nil
}
func (s *TourneyServices) CreateTourney(t schema.Tourney) (schema.TourneyGetFullDTO, error) {

	result := s.DB.Create(&t)
	if result.Error != nil || result.RowsAffected == 0 {
		return schema.TourneyGetFullDTO{}, result.Error
	}

	return schema.TourneyGetFullDTO{
		ID:     schema.RegularIDs(t.ID),
		Name:   t.Name,
		Events: helpers.MapEventsBare(t.Events),
	}, nil
}
func (s *TourneyServices) UpdateTourney(t schema.Tourney, c *gin.Context) (schema.TourneyGetFullDTO, error) {
	var id = c.Param("id")
	tourneyID, err := strconv.Atoi(id)

	if err != nil {
		return schema.TourneyGetFullDTO{}, fmt.Errorf("invalid team ID: %w", err)
	}

result := s.DB.Model(&schema.Tourney{}).Where("id = ?", tourneyID).Updates(&t)

	if result.Error != nil || result.RowsAffected == 0{
		return schema.TourneyGetFullDTO{}, fmt.Errorf("tourney not found or update failed: %w", result.Error)
	}

	var updateTourney schema.Tourney
	if err := s.DB.Preload("Events").First(&updateTourney, tourneyID).Error; err != nil {
		return schema.TourneyGetFullDTO{}, err
	}
		dto := schema.TourneyGetFullDTO{
			ID: schema.RegularIDs(updateTourney.ID),
			Name: updateTourney.Name,
			Events: helpers.MapEventsBare(updateTourney.Events),
		}
	return dto, nil	
	
}

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

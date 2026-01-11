package services

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"uneg.edu.ve/servicio-sadu-back/config"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

// conexión a la DB
type AthleteService struct {
	DB *gorm.DB 
}

func NewAthleteService() *AthleteService {
	return &AthleteService{DB: config.DB}
}
//GET  METHOD 
func (s *AthleteService) GetAllAthletes() ([]schema.Athlete, error) {
	var athletes []schema.Athlete
	return athletes, config.DB.Preload("Teams", nil).Find(&athletes).Error
}

//GET BY ID
func (s *AthleteService) GetAthletesByID(c *gin.Context) (schema.Athlete, error) {
	var id = c.Param("id")
	athleteID, err := strconv.Atoi(id)
	if err != nil {
		return schema.Athlete{}, fmt.Errorf("ID invalido: %v", err)
	}
	var athlete schema.Athlete

	result := s.DB.First(&athlete, athleteID)
	if result.Error != nil {
		return schema.Athlete{}, result.Error
	}

	
	return athlete, nil
}

//POST METHOD
func (s *AthleteService) CreateAthlete(a schema.Athlete) (schema.Athlete, error) {
	return a, s.DB.Create(&a).Error
}

//PUT METHOD
func (s *AthleteService) EditAthlete(a schema.Athlete, c *gin.Context) (schema.Athlete, error){
	var id = c.Param("id")
	athleteID, err := strconv.Atoi(id)

	if err != nil{
		return schema.Athlete{}, fmt.Errorf("ID invalido: %v",err)
	}
	    result := s.DB.Model(&schema.Athlete{}).Where("id = ?", athleteID).Updates(&a)

    if result.Error != nil {
        return schema.Athlete{}, result.Error
    }
    
    if result.RowsAffected == 0 {
        return schema.Athlete{}, fmt.Errorf("atleta no encontrado: %d", athleteID)
    }

    var updatedAthlete schema.Athlete
    if err := s.DB.First(&updatedAthlete, athleteID).Error; err != nil {
        return schema.Athlete{}, err
    }
    
    return updatedAthlete, nil
}
//DELETE METHOD
func (s *AthleteService) DeleteAthlete( c *gin.Context) (error) {
	var id = c.Param("id")
	athleteID,err := strconv.Atoi(id)

	if err != nil { 
		return  err
	}
	result := s.DB.Delete(&schema.Athlete{},athleteID)
	if result != nil{
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("Athlete not found: %d", athleteID)

	}
	return nil
}
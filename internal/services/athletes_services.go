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
	

func (s *AthleteService) GetAllAthletes() ([]schema.AthleteDTO, error) {
	var athletes []schema.Athlete
	if err := s.DB.Omit("Discipline").Find(&athletes).Error; err != nil {
		return nil,err 
	}
	 athleteDTO := make([]schema.AthleteDTO, len(athletes))
	for i, value := range athletes {
		athleteDTO[i] = schema.AthleteDTO{
			ID: schema.RegularIDs(value.ID),
			GovID: value.GovID,
			FirstNames: value.FirstNames,
			LastNames: value.LastNames,
			PhoneNum: value.PhoneNum,
			Email: value.Email,
			Inscripted: value.Inscripted,
			Regular: value.Regular,

		}
	}
	return athleteDTO,nil
}

//GET BY ID
func (s *AthleteService) GetAthletesByID(c *gin.Context) (schema.Athlete, error) {
	var id = c.Param("id")
	athleteID, err := strconv.Atoi(id)
	if err != nil {
		return schema.Athlete{}, fmt.Errorf("ID invalido: %v", err)
	}
	var athlete schema.Athlete

	result := s.DB.Preload("Teams", nil).Preload("Disciplines").Preload("Events").First(&athlete, athleteID)
	if result.Error != nil {
		return schema.Athlete{}, result.Error
	}
	
	return athlete, nil
}

//POST METHOD
func (s *AthleteService) CreateAthlete(a schema.Athlete) (schema.Athlete, error) {
	if err := s.DB.Omit("Teams", "Events", "Disciplines").Create(&a).Error; err != nil {
		return a, err
	}

	if len(a.Disciplines) > 0 {
		s.DB.Model(&a).Association("Disciplines").Append(a.Disciplines)
	}

	if len(a.Events) > 0 {
		s.DB.Model(&a).Association("Events").Append(a.Events)
	}

	if len(a.Teams) > 0 {
		s.DB.Model(&a).Association("Teams").Append(a.Teams)
	}
	return a, nil
}

//PUT METHOD
func (s *AthleteService) EditAthlete(a schema.Athlete, c *gin.Context) (schema.Athlete, error){
	var id = c.Param("id")
	athleteID, err := strconv.Atoi(id)

	if err != nil{
		return schema.Athlete{}, fmt.Errorf("ID invalido: %v",err)
	}
	   

    var athlete schema.Athlete
    if err := s.DB.First(&athlete, athleteID).Error; err != nil {
        return schema.Athlete{}, fmt.Errorf("atleta no encontrado: %d", athleteID)
    }

	//Actualizar campos escalares
    s.DB.Model(&athlete).Select("FirstNames", "LastNames", "Email").Updates(&a)

	if len(a.Teams) > 0 {
        s.DB.Model(&athlete).Association("Teams").Replace(a.Teams)
    }

	if len(a.Disciplines) > 0 {
        s.DB.Model(&athlete).Association("Disciplines").Replace(a.Disciplines)
    }
	if len(a.Events) > 0 {
        s.DB.Model(&athlete).Association("Events").Replace(a.Events)
    }

    return athlete, s.DB.Preload("Teams").Preload("Disciplines").Preload("Events").First(&athlete, athleteID).Error
}
//DELETE METHOD
func (s *AthleteService) DeleteAthlete( c *gin.Context) (error) {
	var id = c.Param("id")
	athleteID,err := strconv.Atoi(id)

	if err != nil { 
		return  fmt.Errorf("ID inválido: %w", err)
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
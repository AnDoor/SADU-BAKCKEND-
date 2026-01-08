package services

import (
	"gorm.io/gorm"
	"uneg.edu.ve/servicio-sadu-back/config"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

type AthleteService struct {
	DB *gorm.DB // ← Tu conexión DB
}

func NewAthleteService() *AthleteService {
	return &AthleteService{DB: config.DB}
}
func (s *AthleteService) GetAllAthletes() ([]schema.Athlete, error) {
	var athletes []schema.Athlete
	return athletes, config.DB.Preload("Teams",nil).Find(&athletes).Error
}

func (s *AthleteService) CreateAthlete(a schema.Athlete) (schema.Athlete, error) {
	return a, s.DB.Create(&a).Error
}

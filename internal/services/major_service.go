package services

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"uneg.edu.ve/servicio-sadu-back/config"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

type MajorServices struct {
	DB *gorm.DB
}

func newMajorServices() *MajorServices {
	return &MajorServices{DB: config.DB}
}

func (s *MajorServices) GetAllMajor() ([]schema.Major, error) {
	var major []schema.Major
	return major, s.DB.Preload("Athletes", nil).Find(&major).Error
}

func (s *MajorServices) GetMajorByID(ctx *gin.Context) (schema.Major, error) {
	var id = ctx.Param("id")
	majorID, err := strconv.Atoi(id)
	if err != nil {
		return schema.Major{}, fmt.Errorf("ERROR: ID INVALID %w", err)
	}
	var major schema.Major

	result := s.DB.First(&major, majorID)
	if result.Error != nil {
		return schema.Major{}, result.Error
	}

	if result.RowsAffected == 0 {
		return schema.Major{}, fmt.Errorf("major %d not found", majorID)
	}
	return major, nil
}

func (s *MajorServices) CreateMajor(m schema.MajorCreateDTO) (schema.MajorGetDTO, error) {

	major := schema.Major{
		Name: m.Name,
	}

	result := s.DB.Create(&major)
	if result.Error != nil {
		return schema.MajorGetDTO{}, fmt.Errorf("creando carrera: %w", result.Error)
	}
	if len(m.AthleteIDs) > 0 {
		athletes := make([]schema.Athlete, len(m.AthleteIDs))
		for i, athleteID := range m.AthleteIDs {
			athletes[i].ID = uint(athleteID)
		}
		s.DB.Model(&major).Association("Athletes").Append(athletes)
	}

	return schema.MajorGetDTO{
		ID:   schema.RegularIDs(major.ID),
		Name: major.Name,
	}, nil
}

func (s *MajorServices) EditMajor(m schema.Major, ctx *gin.Context) (schema.Major, error) {
	var id = ctx.Param("id")
	majorID, err := strconv.Atoi(id)

	if err != nil {
		return schema.Major{}, fmt.Errorf("ID invalido: %v", err)
	}
	result := s.DB.Model(&schema.Major{}).Where("id = ?", majorID).Updates(&m)

	if result.Error != nil {
		return schema.Major{}, result.Error
	}

	if result.RowsAffected == 0 {
		return schema.Major{}, fmt.Errorf("atleta no encontrado: %d", majorID)
	}

	var updatedMajor schema.Major
	if err := s.DB.First(&updatedMajor, majorID).Error; err != nil {
		return schema.Major{}, err
	}

	return updatedMajor, nil
}

func (s *MajorServices) DeleteMajor(ctx *gin.Context) error {
	var id = ctx.Param("id")
	majorID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	result := s.DB.Delete(&schema.Major{}, majorID)
	if result != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("Athlete not found: ID:%d", majorID)
	}
	return nil

}

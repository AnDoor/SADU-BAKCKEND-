package services

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"uneg.edu.ve/servicio-sadu-back/config"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

type UniversityServices struct {
	DB *gorm.DB
}

func NewUniversityService() *UniversityServices {
	return &UniversityServices{DB: config.DB}
}

func (s *UniversityServices) GetAllUniversity(name string, local string) ([]schema.UniversityGetBareDTO, error) {
	var universities []schema.University
	query := s.DB.Model(&schema.University{}).Preload("Teams", nil)

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if local == "true" || local == "false" {
		isLocal := local == "true"
		query = query.Where("local = ?", isLocal)
	}
	if err := query.Find(&universities).Error; err != nil {
		return nil, err
	}

	universitiesBareDTO := []schema.UniversityGetBareDTO{}
	for _, university := range universities {
		universitiesBareDTO = append(universitiesBareDTO, schema.UniversityGetBareDTO{
			ID:   schema.RegularIDs(university.ID),
			Name: university.Name,
		})
	}

	return universitiesBareDTO, nil
}

func (s *UniversityServices) GetUniversityByID(ctx *gin.Context) (schema.University, error) {
	var id = ctx.Param("id")

	universityID, err := strconv.Atoi(id)

	if err != nil {
		return schema.University{}, fmt.Errorf("ID invalid: %v\n ERROR:%v", universityID, err)
	}

	var universities schema.University
	if result := s.DB.Preload("Teams").First(&universities, universityID); result.Error != nil {
		return schema.University{}, fmt.Errorf("Universidad no encontrada: %v", result.Error)
	}

	return universities, nil
}

func (s *UniversityServices) CreateUniversity(u schema.University) (schema.University, error) {

	if err := s.DB.Omit("Teams").Create(&u).Error; err != nil {
		return schema.University{}, fmt.Errorf("error creando universidad: %v", err)
	}
	if len(u.Teams) > 0 {
		s.DB.Model(&u).Preload("Teams").Association("Teams").Append(u.Teams)
	}

	return u, nil

}

func (s *UniversityServices) EditUniversity(u schema.University, ctx *gin.Context) (schema.University, error) {
	var id = ctx.Param("id")
	var uni schema.University

	
	if err := s.DB.Preload("Teams").First(&uni, id).Error; err != nil {
		return schema.University{}, fmt.Errorf("universidad no encontrada: %w", err)
	}

	if err := s.DB.Model(&uni).Select("Name").Updates(u).Error; err != nil {
		return schema.University{}, err
	}

	if u.Teams != nil {
		
		if err := s.DB.Model(&uni).Association("Teams").Replace(u.Teams); err != nil {
			return schema.University{}, err
		}
	}

	err := s.DB.Preload("Teams").First(&uni, id).Error
	return uni, err
}

func (s *UniversityServices) DeleteUniversity(ctx *gin.Context) error {
	var id = ctx.Param("id")
	universityID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("ERROR: ID INVALID %v: %w", universityID, err)
	}
	result := s.DB.Delete(&schema.University{}, universityID)
	if result.Error != nil {
		return fmt.Errorf("error deleting university %d: %w", universityID, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("University with ID:%d is not FOUND", universityID)
	}
	return nil
}

package services

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"uneg.edu.ve/servicio-sadu-back/config"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

type DisciplineServices struct {
	DB *gorm.DB
}

func NewDisciplineServices() *DisciplineServices {
	return &DisciplineServices{DB: config.DB}
}

func (d *DisciplineServices) GetAllDisciplines() ([]schema.DisciplineGetBareDTO, error) {
	var disciplines []schema.Discipline
	if err := d.DB.Preload("Teams", nil).Find(&disciplines).Error; err != nil {
		return nil, err
	}
	var disciplinesDTO []schema.DisciplineGetBareDTO

	for _, value := range disciplines {
		disciplinesDTO = append(disciplinesDTO, schema.DisciplineGetBareDTO{
			ID:   schema.RegularIDs(value.ID),
			Name: value.Name,
		})
	}
	return disciplinesDTO, nil
}

func (d *DisciplineServices) GetAllDisciplinesByID(c *gin.Context) (schema.DisciplineGetBareDTO, error) {
	var id = c.Param("id")
	disciplineID, err := strconv.Atoi(id)
	if err != nil {
		return schema.DisciplineGetBareDTO{}, fmt.Errorf("ID invalido | %v", err)
	}
	var discipline schema.Discipline
	if result := d.DB.First(&discipline, disciplineID).Error; result != nil {
		return schema.DisciplineGetBareDTO{}, fmt.Errorf("Discipline not found | %w", result)
	}

	return schema.DisciplineGetBareDTO{
		ID:   schema.RegularIDs(discipline.ID),
		Name: discipline.Name,
	}, nil
}

func (d *DisciplineServices) CreateDiscipline(dis schema.Discipline) (schema.DisciplineGetBareDTO, error) {

	if err := d.DB.Create(&dis).Error; err != nil {
		return schema.DisciplineGetBareDTO{}, fmt.Errorf("ERROR creating a new discipline: %w", err)
	}

	return schema.DisciplineGetBareDTO{
		ID:   schema.RegularIDs(dis.ID),
		Name: dis.Name,
	}, nil
}

func (d *DisciplineServices) EditDiscipline(dis schema.Discipline, ctx *gin.Context) (schema.DisciplineGetBareDTO, error) {

	var id = ctx.Param("id")
	disciplineID, err := strconv.Atoi(id)
	if err != nil {
		return schema.DisciplineGetBareDTO{}, fmt.Errorf("ID invalido | %w", err)
	}
	result := d.DB.Model(&schema.Discipline{}).Where("id = ?", disciplineID).Updates(&dis)
	if result.Error != nil {
		return schema.DisciplineGetBareDTO{}, fmt.Errorf("error actualizando disciplina %d: %w", disciplineID, result.Error)

	}
	if result.RowsAffected == 0 {
		return schema.DisciplineGetBareDTO{}, fmt.Errorf("disciplina %d no encontrada", disciplineID)
	}

	return schema.DisciplineGetBareDTO{
		ID:   schema.RegularIDs(dis.ID),
		Name: dis.Name,
	}, nil
}

func (d *DisciplineServices) DeleteDiscipline(ctx *gin.Context) error {
	var id = ctx.Param("id")
	disciplineID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("ID invalid: %w", err)
	}
	result := d.DB.Delete(&schema.Discipline{}, disciplineID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("Column not affected")
	}
	return nil
}

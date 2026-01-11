package services

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"uneg.edu.ve/servicio-sadu-back/config"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

type TeacherService struct {
	DB *gorm.DB
}

func NewTeacherService() *TeacherService {
	return &TeacherService{DB: config.DB}
}

func mapDisciplines(disciplines []schema.Discipline) []schema.DisciplineGetBareDTO {
	dtos := make([]schema.DisciplineGetBareDTO, len(disciplines))
	for i, disc := range disciplines {
		dtos[i] = schema.DisciplineGetBareDTO{
			ID:   schema.RegularIDs(disc.ID),
			Name: disc.Name,
		}
	}
	return dtos
}

func (s *TeacherService) GetTeachers() ([]schema.TeacherGetDTO, error) {
	var teachers []schema.Teacher

	if err := s.DB.Preload("Disciplines").Find(&teachers).Error; err != nil {

		return nil, err
	}

	teachersDTO := make([]schema.TeacherGetDTO, len(teachers))

	for i, teacher := range teachers {
		teachersDTO[i] = schema.TeacherGetDTO{
			ID:          schema.RegularIDs(teacher.ID),
			FirstNames:  teacher.FirstNames,
			LastNames:   teacher.LastNames,
			PhoneNum:    teacher.PhoneNum,
			Email:       teacher.Email,
			GovID:       teacher.GovID,
			Disciplines: mapDisciplines(teacher.Disciplines),
		}
	}
	return teachersDTO, nil

}
func (s *TeacherService) GetTeacherById(ctx *gin.Context) (schema.TeacherGetBareDTO, error) {
	var id = ctx.Param("id")
	teacherId, err := strconv.Atoi(id)

	if err != nil {
		return schema.TeacherGetBareDTO{}, fmt.Errorf("ID INVALID: %w", err)
	}
	var teacher schema.Teacher
	result := s.DB.Preload("Disciplines").First(&teacher, teacherId)
	if result.Error != nil {
		return schema.TeacherGetBareDTO{}, fmt.Errorf("profesor %d no encontrado: %w", teacherId, result.Error)
	}
	return schema.TeacherGetBareDTO{
		ID:          schema.RegularIDs(teacherId),
		FirstNames:  teacher.FirstNames,
		LastNames:   teacher.LastNames,
		GovID:       teacher.GovID,
		Disciplines: mapDisciplines(teacher.Disciplines),
	}, nil
}
func (s *TeacherService) CreateTeacher(t schema.TeacherCreateDTO) (schema.TeacherGetDTO, error) {
	// 1. Validar disciplinas
	for _, discID := range t.DisciplineIDs {
		var disc schema.Discipline
		if err := s.DB.First(&disc, uint(discID)).Error; err != nil {
			return schema.TeacherGetDTO{}, fmt.Errorf("disciplina %d no existe", discID)
		}
	}

	// 2. Crear modelo desde DTO
	teacher := schema.Teacher{
		FirstNames: t.FirstNames,
		LastNames:  t.LastNames,
		PhoneNum:   t.PhoneNum,
		Email:      t.Email,
		GovID:      t.GovID,
	}

	// 3. Guardar
	result := s.DB.Create(&teacher)
	if result.Error != nil {
		return schema.TeacherGetDTO{}, fmt.Errorf("creando profesor: %w", result.Error)
	}

	// 4. Asignar disciplinas (many2many)
	if len(t.DisciplineIDs) > 0 {
		disciplines := make([]schema.Discipline, len(t.DisciplineIDs))
		for i, discID := range t.DisciplineIDs {
			disciplines[i].ID = uint(discID)
		}
		s.DB.Model(&teacher).Association("Disciplines").Append(disciplines)
	}

	// 5. Retornar DTO completo CON Disciplines
	var fullTeacher schema.Teacher
	s.DB.Preload("Disciplines").First(&fullTeacher, teacher.ID)

	return schema.TeacherGetDTO{
		ID:          schema.RegularIDs(fullTeacher.ID),
		FirstNames:  fullTeacher.FirstNames,
		LastNames:   fullTeacher.LastNames,
		PhoneNum:    fullTeacher.PhoneNum,
		Email:       fullTeacher.Email,
		GovID:       fullTeacher.GovID,
		Disciplines: mapDisciplines(fullTeacher.Disciplines),
	}, nil
}
func (s *TeacherService) EditTeacher(ctx *gin.Context) (schema.TeacherGetDTO, error){
	id := ctx.Param("id")
    teacherID, err := strconv.Atoi(id)
    if err != nil {
        return schema.TeacherGetDTO{}, fmt.Errorf("ID inválido: %w", err)
    }

	var input schema.TeacherGetDTO
    if err := ctx.ShouldBindJSON(&input); err != nil {
        return schema.TeacherGetDTO{}, fmt.Errorf("datos inválidos: %w", err)
    }
	var teacher schema.Teacher
    result := s.DB.Preload("Disciplines").First(&teacher, teacherID)
    if result.Error != nil {
        return schema.TeacherGetDTO{}, fmt.Errorf("profesor %d no encontrado", teacherID)
    }

 
    if input.FirstNames != "" {
        teacher.FirstNames = input.FirstNames
    }
    if input.LastNames != "" {
        teacher.LastNames = input.LastNames
    }
    if input.PhoneNum != "" {
        teacher.PhoneNum = input.PhoneNum
    }
    if input.Email != "" {
        teacher.Email = input.Email
    }
    if input.GovID != "" {
        teacher.GovID = input.GovID
    }
	if len(input.Disciplines) > 0 {
        disciplineIDs := make([]uint, len(input.Disciplines))
        for i, disc := range input.Disciplines {
            disciplineIDs[i] = uint(disc.ID)
        }
        
 
        s.DB.Model(&teacher).Association("Disciplines").Clear()
        disciplines := make([]schema.Discipline, len(disciplineIDs))
        for i, discID := range disciplineIDs {
            disciplines[i].ID = discID
        }
        s.DB.Model(&teacher).Association("Disciplines").Append(disciplines)
    }

    if err := s.DB.Save(&teacher).Error; err != nil {
        return schema.TeacherGetDTO{}, fmt.Errorf("actualizando profesor: %w", err)
    }

    return schema.TeacherGetDTO{
        ID:          schema.RegularIDs(teacher.ID),
        FirstNames:  teacher.FirstNames,
        LastNames:   teacher.LastNames,
        PhoneNum:    teacher.PhoneNum,
        Email:       teacher.Email,
        GovID:       teacher.GovID,
        Disciplines: mapDisciplines(teacher.Disciplines),
    }, nil
}
func (s *TeacherService) DeleteTeacher(ctx *gin.Context) error{
	id := ctx.Param("id")
    teacherID, err := strconv.Atoi(id)
    if err != nil {
        return fmt.Errorf("ID inválido: %w", err)
    }

	 var teacher schema.Teacher
    result := s.DB.First(&teacher, teacherID)
    if result.Error != nil {
        return fmt.Errorf("profesor %d no encontrado", teacherID)
    }

    
    result = s.DB.Delete(&schema.Teacher{}, teacherID)
    if result.Error != nil {
        return fmt.Errorf("error eliminando profesor %d: %w", teacherID, result.Error)
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("profesor %d no encontrado", teacherID)
    }
	return nil
}

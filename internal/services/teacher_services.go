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

type TeacherService struct {
	DB *gorm.DB
}

func NewTeacherService() *TeacherService {
	return &TeacherService{DB: config.DB}
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
			Disciplines: helpers.MapDisciplines(teacher.Disciplines),
		}
	}
	return teachersDTO, nil

}
func (s *TeacherService) GetTeacherById(ctx *gin.Context) (schema.TeacherGetDTO, error) {
	var id = ctx.Param("id")
	teacherId, err := strconv.Atoi(id)

	if err != nil {
		return schema.TeacherGetDTO{}, fmt.Errorf("ID INVALID: %w", err)
	}
	var teacher schema.Teacher
	result := s.DB.Preload("Disciplines").First(&teacher, teacherId)
	if result.Error != nil {
		return schema.TeacherGetDTO{}, fmt.Errorf("profesor %d no encontrado: %w", teacherId, result.Error)
	}
	return schema.TeacherGetDTO{
		ID:          schema.RegularIDs(teacherId),
		FirstNames:  teacher.FirstNames,
		LastNames:   teacher.LastNames,
		PhoneNum:    teacher.PhoneNum,
		Email:       teacher.Email,
		GovID:       teacher.GovID,
		Disciplines: helpers.MapDisciplines(teacher.Disciplines),
	}, nil
}
func (s *TeacherService) CreateTeacher(t schema.TeacherCreateDTO) (schema.TeacherGetDTO, error) {

	for _, discID := range t.DisciplineIDs {
		var disc schema.Discipline
		if err := s.DB.First(&disc, uint(discID)).Error; err != nil {
			return schema.TeacherGetDTO{}, fmt.Errorf("disciplina %d no existe", discID)
		}
	}

	//  Crear modelo desde DTO
	teacher := schema.Teacher{
		FirstNames: t.FirstNames,
		LastNames:  t.LastNames,
		PhoneNum:   t.PhoneNum,
		Email:      t.Email,
		GovID:      t.GovID,
	}

	if err := s.DB.Create(&teacher).Error; err != nil {
        return schema.TeacherGetDTO{}, fmt.Errorf("creando profesor: %w", err)
    }

	//  Asignar disciplinas many2many
	if len(t.DisciplineIDs) > 0 {
		disciplineIDs := make([]uint, len(t.DisciplineIDs))
        for i, discID := range t.DisciplineIDs {
            disciplineIDs[i] = uint(discID)
        }
		if err := s.DB.Model(&teacher).Association("Disciplines").Replace(disciplineIDs).Error; err != nil {
			return schema.TeacherGetDTO{}, fmt.Errorf("asignando disciplinas: %w", err)
		}
	}

	if err := s.DB.Preload("Disciplines").First(&teacher, teacher.ID).Error; err != nil {
        return schema.TeacherGetDTO{}, fmt.Errorf("cargando profesor completo: %w", err)
    }

	// Retornar DTO completo
	return schema.TeacherGetDTO{
        ID:          schema.RegularIDs(teacher.ID),
        FirstNames:  teacher.FirstNames,
        LastNames:   teacher.LastNames,
        PhoneNum:    teacher.PhoneNum,
        Email:       teacher.Email,
        GovID:       teacher.GovID,
        Disciplines: helpers.MapDisciplines(teacher.Disciplines),
    }, nil
}

func (s *TeacherService) EditTeacher(ctx *gin.Context, t schema.Teacher) (schema.Teacher, error) {
 
    id := ctx.Param("id")
    teacherID, err := strconv.Atoi(id)
    if err != nil {
        return schema.Teacher{}, fmt.Errorf("ID inválido: %w", err)
    }

   
    var teacher schema.Teacher
    result := s.DB.Preload("Disciplines").First(&teacher, teacherID)
    if result.Error != nil {
        return schema.Teacher{}, fmt.Errorf("profesor %d no encontrado: %w", teacherID, result.Error)
    }

  
    for _, discipline := range t.Disciplines {
        var disc schema.Discipline
        if err := s.DB.First(&disc, discipline.ID).Error; err != nil {
            return schema.Teacher{}, fmt.Errorf("disciplina %d no existe: %w", discipline.ID, err)
        }
    }

   
    teacher.FirstNames = t.FirstNames
    teacher.LastNames  = t.LastNames
    teacher.PhoneNum   = t.PhoneNum
    teacher.Email      = t.Email
    teacher.GovID      = t.GovID

    // 5. Guardar + reemplazar disciplinas
    if err := s.DB.Save(&teacher).Error; err != nil {
        return schema.Teacher{}, fmt.Errorf("guardando profesor: %w", err)
    }

    if len(t.Disciplines) > 0 {
        disciplineIDs := make([]uint, len(t.Disciplines))
        for i, disc := range t.Disciplines {
            disciplineIDs[i] = disc.ID
        }
        if err := s.DB.Model(&teacher).Association("Disciplines").Replace(disciplineIDs); err != nil {
            return schema.Teacher{}, fmt.Errorf("actualizando disciplinas: %w", err)
        }
    }


    if err := s.DB.Preload("Events").Preload("Disciplines").First(&teacher, teacherID).Error; err != nil {
        return schema.Teacher{}, fmt.Errorf("recargando profesor: %w", err)
    }

    return teacher, nil  
}


func (s *TeacherService) DeleteTeacher(ctx *gin.Context) error {
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

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
		Disciplines: helpers.MapDisciplines(teacher.Disciplines),
	}, nil
}
func (s *TeacherService) CreateTeacher(t schema.TeacherCreateDTO) (schema.TeacherGetDTO, error) {
	//  Validar disciplinas
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

	//  Guardar
	 if err := s.DB.Omit("Events", "Disciplines").Create(&teacher).Error; err != nil {
        return schema.TeacherGetDTO{}, fmt.Errorf("creando profesor: %w", err)
    }

	 //  Asignar disciplinas many2many
    if len(t.DisciplineIDs) > 0 {
        disciplines := make([]schema.Discipline, len(t.DisciplineIDs))
        for i, discID := range t.DisciplineIDs {
            disciplines[i].ID = uint(discID)
        }
        if err := s.DB.Model(&teacher).Association("Disciplines").Append(disciplines).Error; err != nil {
            return schema.TeacherGetDTO{}, fmt.Errorf("asignando disciplinas: %w", err)
        }
    }

    // Recargar con Preload para DTO completo
    var fullTeacher schema.Teacher
    if err := s.DB.Preload("Disciplines").First(&fullTeacher, teacher.ID).Error; err != nil {
        return schema.TeacherGetDTO{}, fmt.Errorf("cargando profesor completo: %w", err)
    }

    // Retornar DTO completo
    return schema.TeacherGetDTO{
        ID:          schema.RegularIDs(fullTeacher.ID),
        FirstNames:  fullTeacher.FirstNames,
        LastNames:   fullTeacher.LastNames,
        PhoneNum:    fullTeacher.PhoneNum,
        Email:       fullTeacher.Email,
        GovID:       fullTeacher.GovID,
        Disciplines: helpers.MapDisciplines(fullTeacher.Disciplines),
    }, nil

}
func (s *TeacherService) EditTeacher(ctx *gin.Context) (schema.TeacherGetDTO, error) {
	id := ctx.Param("id")
	teacherID, err := strconv.Atoi(id)
	if err != nil {
		return schema.TeacherGetDTO{}, fmt.Errorf("ID inválido: %w", err)
	}

	var input schema.TeacherCreateDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return schema.TeacherGetDTO{}, fmt.Errorf("datos inválidos: %w", err)
	}


	var teacher schema.Teacher
	result := s.DB.Preload("Disciplines").First(&teacher, teacherID)
	if result.Error != nil {
		return schema.TeacherGetDTO{}, fmt.Errorf("profesor %d no encontrado", teacherID)
	}

	//validar disciplinas 
	for _, discID := range input.DisciplineIDs {
        var disc schema.Discipline
        if err := s.DB.First(&disc, uint(discID)).Error; err != nil {
            return schema.TeacherGetDTO{}, fmt.Errorf("disciplina %d no existe", discID)
        }
    }

	//Actualizar campos (solo si vienen en request)
    updates := make(map[string]interface{})
    if input.FirstNames != "" { updates["first_names"] = input.FirstNames }
    if input.LastNames != "" { updates["last_names"] = input.LastNames }
    if input.PhoneNum != "" { updates["phone_num"] = input.PhoneNum }
    if input.Email != "" { updates["email"] = input.Email }
    if input.GovID != "" { updates["gov_id"] = input.GovID }

	if len(updates) > 0 {
        if err := s.DB.Model(&teacher).Updates(updates).Error; err != nil {
            return schema.TeacherGetDTO{}, fmt.Errorf("actualizando datos profesor: %w", err)
        }
    }

    //  Reemplazar disciplinas (Clear + Append = Replace)
    if len(input.DisciplineIDs) > 0 {
        disciplines := make([]schema.Discipline, len(input.DisciplineIDs))
        for i, discID := range input.DisciplineIDs {
            disciplines[i].ID = uint(discID)
        }
   
        if err := s.DB.Model(&teacher).Association("Disciplines").Replace(disciplines).Error; err != nil {
            return schema.TeacherGetDTO{}, fmt.Errorf("actualizando disciplinas: %w", err)
        }
    }

    //  Recargar con Preload para respuesta completa
    if err := s.DB.Preload("Disciplines").First(&teacher, teacherID).Error; err != nil {
        return schema.TeacherGetDTO{}, fmt.Errorf("recargando profesor: %w", err)
    }

    //  Retornar DTO completo
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

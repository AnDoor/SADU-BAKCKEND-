package services

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"uneg.edu.ve/servicio-sadu-back/config"
	"uneg.edu.ve/servicio-sadu-back/helpers"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

type EventService struct {
	DB *gorm.DB
}

func NewEventService() *EventService {
	return &EventService{DB: config.DB}
}

func (s *EventService) GetEvents(id uint, name, status string) ([]schema.EventGetDTO, error) {

	var event []schema.Event

	query := s.DB.Preload("HomeTeam").
		Preload("HomeTeam.University").
		Preload("OppositeTeam.University").
		Preload("Tourney").
		Preload("ResponsableTeacher.Disciplines").
		Preload("ResponsableTeacher").
		Preload("Discipline")

	if id != 0 {
		query = query.Where("events.id = ?", id)
	}
	if name != " " {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if status != "" {
		query = query.Where("status LIKE ?", "%"+status+"%")
	}
	if err := query.Find(&event).Error; err != nil {
		return nil, nil
	}

	dto := helpers.MapEventsGetDTO(event)

	return dto, nil
}

func (s *EventService) CreateEvent(event schema.Event) (schema.Event, error) {
	var tourney schema.Tourney
	tx := s.DB.Begin()

	if event.DisciplineID != tourney.DisciplineID {
		return schema.Event{}, fmt.Errorf("la disciplina del evento no coincide con la del torneo")
	}

	if err := tx.Create(&event).Error; err != nil {
		tx.Rollback()
		return schema.Event{}, err
	}

	if len(event.Athletes) > 0 {
		if err := tx.Model(&event).Association("Athletes").Replace(event.Athletes); err != nil {
			tx.Rollback()
			return schema.Event{}, err
		}
	}

	tx.Commit()

	s.DB.Preload("HomeTeam").
		Preload("OppositeTeam").
		Preload("Tourney").
		Preload("ResponsableTeacher").
		Preload("Discipline").
		Preload("Athletes").
		First(&event, event.ID)

	return event, nil
}
func (s *EventService) EditEvent(ctx *gin.Context) (schema.Event, error) {

	id := ctx.Param("id")
	var input schema.Event

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return schema.Event{}, err
	}

	var existingEvent schema.Event
	if err := s.DB.First(&existingEvent, id).Error; err != nil {
		return schema.Event{}, err // Retornará gorm.ErrRecordNotFound si no existe
	}

	tx := s.DB.Begin()

	err := tx.Model(&existingEvent).Omit("HomeTeam", "OppositeTeam", "Tourney", "ResponsableTeacher", "Discipline", "Athletes").
		Updates(input).Error

	if err != nil {
		tx.Rollback()
		return schema.Event{}, err
	}

	if input.Athletes != nil {
		if err := tx.Model(&existingEvent).Association("Athletes").Replace(input.Athletes); err != nil {
			tx.Rollback()
			return schema.Event{}, err
		}
	}

	tx.Commit()

	s.DB.Preload("HomeTeam").
		Preload("OppositeTeam").
		Preload("Tourney").
		Preload("ResponsableTeacher").
		Preload("Discipline").
		Preload("Athletes").
		First(&existingEvent, id)

	return existingEvent, nil

}
func (s *EventService) DeleteEvent(ctx *gin.Context) error {
	id := ctx.Param("id")

	result := s.DB.Delete(&schema.Event{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

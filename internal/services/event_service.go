package services

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"uneg.edu.ve/servicio-sadu-back/config"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

type EventService struct {
	DB *gorm.DB
}

func NewEventService() *EventService {
	return &EventService{DB: config.DB}
}

func (s *EventService) GetAllEvents(name, status string) ([]schema.EventGetBareDTO, error) {
	query := s.DB.Preload("Athletes")
	var event []schema.Event

	if name != " " {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if status != "" {
		query = query.Where("status LIKE ?", "%"+status+"%")
	}
	if err := query.Find(&event).Error; err != nil {
		return nil, nil
	}
	eventDTO := []schema.EventGetBareDTO{}
	for _, value := range event {
		eventDTO = append(eventDTO, schema.EventGetBareDTO{
			ID:             schema.RegularIDs(value.ID),
			Name:           value.Name,
			Date:           value.Date,
			Status:         value.Status,
			HomePoints:     value.HomePoints,
			OppositePoints: value.OppositePoints,
			// ResponsableTeacher: value.ResponsableTeacher,
			// Discipline: value.Discipline,
		})

	}

	return eventDTO, nil
}
func (s *EventService) GetEventByID(c *gin.Context) (schema.EventGetDTO, error) {
	id := c.Param("id")

	var event schema.Event

	err := s.DB.Preload("HomeTeam").
		Preload("OppositeTeam").
		Preload("Tourney").
		Preload("ResponsableTeacher").
		Preload("Discipline").
		First(&event, id).Error

	if err != nil {
		return schema.EventGetDTO{}, err
	}

	dto := schema.EventGetDTO{
		ID:             schema.RegularIDs(event.ID),
		Name:           event.Name,
		Date:           event.Date,
		Status:         event.Status,
		Observation:    event.Observation,
		Ubication:      event.Ubication,
		HomePoints:     event.HomePoints,
		OppositePoints: event.OppositePoints,

		HomeTeam: schema.TeamGetBareDTO{
			ID:   schema.RegularIDs(event.HomeTeam.ID),
			Name: event.HomeTeam.Name,
		},
		OppositeTeam: schema.TeamGetBareDTO{
			ID:   schema.RegularIDs(event.OppositeTeam.ID),
			Name: event.OppositeTeam.Name,
		},
		Tourney: schema.TourneyGetBareDTO{
			ID:   schema.RegularIDs(event.Tourney.ID),
			Name: event.Tourney.Name,
		},
		ResponsableTeacher: schema.TeacherGetBareDTO{
			ID:        schema.RegularIDs(event.ResponsableTeacher.ID),
			FirstNames: event.ResponsableTeacher.FirstNames,
			LastNames:  event.ResponsableTeacher.LastNames,
		},
		Discipline: schema.DisciplineGetBareDTO{
			ID:   schema.RegularIDs(event.Discipline.ID),
			Name: event.Discipline.Name,
		},
	}

	return dto, nil
}

func (s *EventService) CreateEvent(event schema.Event) (schema.Event, error) {
   
    tx := s.DB.Begin()

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
 func (s *EventService) EditEvent (c  *gin.Context) (schema.Event,error){
	
    id := c.Param("id")
    var input schema.Event

    if err := c.ShouldBindJSON(&input); err != nil {
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
 func (s *EventService) DeleteEvent (c *gin.Context) error{
    id := c.Param("id")
	
    result := s.DB.Delete(&schema.Event{}, id)

    if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return gorm.ErrRecordNotFound
    }

    return nil
 }

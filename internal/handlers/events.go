package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"uneg.edu.ve/servicio-sadu-back/schema"
)

func _GetBareEvents(ctx *gin.Context) {
	var events []schema.Event
	if err := DB.Select("id, name").Find(&events).Error; err != nil {
		log.Printf("Error listing events: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Error listing events")
		return
	}

	// Más eficiente: usa make con capacidad conocida
	eventsDTO := make([]schema.EventGetBareDTO, 0, len(events))
	for _, event := range events {
		eventsDTO = append(eventsDTO, schema.EventGetBareDTO{
			ID:   schema.RegularIDs(event.ID),
			Name: event.Name,
		})
	}

	//return sendSuccess(ctx, "listing-events", eventsDTO)
}

func GetBareEvents(ctx *gin.Context) {
	events := []schema.Event{}

	// Preload all the necessary relationships for EventGetBareDTO
	if err := DB.Preload("HomeTeam").
		Preload("HomeTeam.University").
		Preload("HomeTeam.Athletes").
		Preload("OppositeTeam").
		Preload("OppositeTeam.University").
		Preload("OppositeTeam.Athletes").
		Preload("ResponsableTeacher").
		Preload("ResponsableTeacher.Disciplines").
		Preload("Discipline").
		Find(&events).Error; err != nil {
		println(err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error listing events")
		return
	}

	eventsDTO := []schema.EventGetBareDTO{}
	for _, event := range events {
		// Map HomeTeam
		homeTeam := schema.TeamGetBareDTO{
			ID:       schema.RegularIDs(event.HomeTeam.ID),
			Name:     event.HomeTeam.Name,
			Regular:  event.HomeTeam.Regular,
			Category: string(event.HomeTeam.Category),
			University: schema.UniversityGetBareDTO{
				ID:    schema.RegularIDs(event.HomeTeam.UniversityID),
				Name:  event.HomeTeam.University.Name,
				Local: event.HomeTeam.University.Local,
			},
			Athletes: event.HomeTeam.Athletes,
		}

		// Map OppositeTeam
		oppositeTeam := schema.TeamGetBareDTO{
			ID:       schema.RegularIDs(event.OppositeTeam.ID),
			Name:     event.OppositeTeam.Name,
			Regular:  event.OppositeTeam.Regular,
			Category: string(event.OppositeTeam.Category),
			University: schema.UniversityGetBareDTO{
				ID:    schema.RegularIDs(event.OppositeTeam.University.ID),
				Name:  event.OppositeTeam.University.Name,
				Local: event.OppositeTeam.University.Local,
			},
			Athletes: event.OppositeTeam.Athletes,
		}

		// Map ResponsableTeacher
		responsableTeacher := schema.TeacherGetBareDTO{
			ID:         schema.RegularIDs(event.ResponsableTeacher.ID),
			FirstNames: event.ResponsableTeacher.FirstNames,
			LastNames:  event.ResponsableTeacher.LastNames,
			GovID:      event.ResponsableTeacher.GovID,
		}

		// Map Disciplines for the teacher
		disciplines := []schema.DisciplineGetBareDTO{}
		for _, discipline := range event.ResponsableTeacher.Disciplines {
			disciplines = append(disciplines, schema.DisciplineGetBareDTO{
				ID:   schema.RegularIDs(discipline.ID),
				Name: discipline.Name,
			})
		}
		responsableTeacher.Disciplines = disciplines

		// Map Discipline
		discipline := schema.DisciplineGetBareDTO{
			ID:   schema.RegularIDs(event.Discipline.ID),
			Name: event.Discipline.Name,
		}

		// Create EventGetBareDTO
		eventDTO := schema.EventGetBareDTO{
			ID:                 schema.RegularIDs(event.ID),
			Name:               event.Name,
			Date:               event.Date,
			Status:             event.Status,
			HomePoints:         event.HomePoints,
			OppositePoints:     event.OppositePoints,
			HomeTeam:           homeTeam,
			OppositeTeam:       oppositeTeam,
			ResponsableTeacher: responsableTeacher,
			Discipline:         discipline,
		}

		eventsDTO = append(eventsDTO, eventDTO)
	}

	sendSucces(ctx, "listing-events", eventsDTO)
	return
}

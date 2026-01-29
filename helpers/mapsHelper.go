package helpers

import "uneg.edu.ve/servicio-sadu-back/schema"

func MapEventsBare(events []schema.Event) []schema.EventGetBareDTO {
	dtos := make([]schema.EventGetBareDTO, len(events))
	for i, event := range events {
		dtos[i] = schema.EventGetBareDTO{
			ID:             schema.RegularIDs(event.ID),
			Name:           event.Name,
			Date:           event.Date,
			Status:         event.Status,
			HomePoints:     event.HomePoints,
			OppositePoints: event.OppositePoints,

			HomeTeam: schema.TeamGetBareDTO{
				ID:       schema.RegularIDs(event.HomeTeam.ID),
				Name:     event.HomeTeam.Name,
				Category: string(event.HomeTeam.Category),
				University: schema.UniversityGetBareDTO{
					ID:    schema.RegularIDs(event.HomeTeam.University.ID),
					Name:  event.HomeTeam.University.Name,
					Local: event.HomeTeam.University.Local,
				},
				Athletes: MapAthletes(event.HomeTeam.Athletes),
			},
			OppositeTeam: schema.TeamGetBareDTO{
				ID:       schema.RegularIDs(event.OppositeTeam.ID),
				Name:     event.OppositeTeam.Name,
				Category: string(event.OppositeTeam.Category),
				University: schema.UniversityGetBareDTO{
					ID:    schema.RegularIDs(event.HomeTeam.University.ID),
					Name:  event.HomeTeam.University.Name,
					Local: event.HomeTeam.University.Local,
				},
				Athletes: MapAthletes(event.OppositeTeam.Athletes),
			},
			ResponsableTeacher: schema.TeacherGetBareDTO{
				ID:         schema.RegularIDs(event.ResponsableTeacher.ID),
				FirstNames: event.ResponsableTeacher.FirstNames,
				LastNames:  event.ResponsableTeacher.LastNames,
				GovID: event.ResponsableTeacher.GovID,
				Disciplines: MapDisciplines(event.ResponsableTeacher.Disciplines),
			},
			Discipline: schema.DisciplineGetBareDTO{
				ID:   schema.RegularIDs(event.Discipline.ID),
				Name: event.Discipline.Name,
			},
		}
	}
	return dtos
}

func MapEventsGetDTO(events []schema.Event) []schema.EventGetDTO {
	dtos := make([]schema.EventGetDTO, len(events))
	for i, event := range events {
		dtos[i] = schema.EventGetDTO{
		ID:             schema.RegularIDs(event.ID),
		Name:           event.Name,
		Date:           event.Date,
		Status:         event.Status,
		Observation:    event.Observation,
		Ubication:      event.Ubication,
		HomePoints:     event.HomePoints,
		OppositePoints: event.OppositePoints,

		HomeTeam: schema.TeamGetBareDTO{
			ID:      schema.RegularIDs(event.HomeTeam.ID),
			Name:    event.HomeTeam.Name,
			Regular: event.HomeTeam.Regular,
			Category: string(event.OppositeTeam.Category),
			University: schema.UniversityGetBareDTO{
				ID:    schema.RegularIDs(event.HomeTeam.University.ID),
				Name:  event.HomeTeam.University.Name,
				Local: event.HomeTeam.University.Local,
			},
			Athletes: MapAthletes(event.HomeTeam.Athletes),
		},
		OppositeTeam: schema.TeamGetBareDTO{
			ID:      schema.RegularIDs(event.HomeTeam.ID),
			Name:    event.OppositeTeam.Name,
			Regular: event.OppositeTeam.Regular,
			Category: string(event.OppositeTeam.Category),
			
			University: schema.UniversityGetBareDTO{
				ID:    schema.RegularIDs(event.OppositeTeam.University.ID),
				Name:  event.OppositeTeam.University.Name,
				Local: event.OppositeTeam.University.Local,
			},
			Athletes: MapAthletes(event.OppositeTeam.Athletes),
		},
		Tourney: schema.TourneyGetBareDTO{
			ID:   schema.RegularIDs(event.Tourney.ID),
			Name: event.Tourney.Name,
			Status: event.Tourney.Status,
		},
		ResponsableTeacher: schema.TeacherGetBareDTO{
			ID:         schema.RegularIDs(event.ResponsableTeacher.ID),
			FirstNames: event.ResponsableTeacher.FirstNames,
			LastNames:  event.ResponsableTeacher.LastNames,
			GovID: event.ResponsableTeacher.GovID,
			Disciplines: MapDisciplines(event.ResponsableTeacher.Disciplines),
		},
		Discipline: schema.DisciplineGetBareDTO{
			ID:   schema.RegularIDs(event.Discipline.ID),
			Name: event.Discipline.Name,
		},
	}


}	
return dtos
}

func MapDisciplines(disciplines []schema.Discipline) []schema.DisciplineGetBareDTO {
	dtos := make([]schema.DisciplineGetBareDTO, len(disciplines))
	for i, disc := range disciplines {
		dtos[i] = schema.DisciplineGetBareDTO{
			ID:   schema.RegularIDs(disc.ID),
			Name: disc.Name,
		}
	}
	return dtos
}

func MapAthletes(athletes []schema.Athlete) []schema.AthleteDTO {

	dtos := make([]schema.AthleteDTO, len(athletes))
	for i, disc := range athletes {
		dtos[i] = schema.AthleteDTO{
			ID:         schema.RegularIDs(disc.ID),
			FirstNames: disc.FirstNames,
			LastNames:  disc.LastNames,
			GovID:      disc.GovID,
			PhoneNumber:   disc.PhoneNumber,
			Email:      disc.Email,
			Enrolled: disc.Enrolled,
			Regular:    disc.Regular,
		}
	}
	return dtos
}

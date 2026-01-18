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
		}
	}
	return dtos
}

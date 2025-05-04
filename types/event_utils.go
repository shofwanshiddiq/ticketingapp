package types

import "ticketingapp/entity"

func ToEventResponse(e entity.Event) EventResponse {
	return EventResponse{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		Location:    e.Location,
		StartTime:   e.StartTime,
		EndTime:     e.EndTime,
		Capacity:    e.Capacity,
		Price:       e.Price,
		Status:      e.Status,
	}
}

func ToEventResponseList(events []entity.Event) []EventResponse {
	res := make([]EventResponse, len(events))
	for i, e := range events {
		res[i] = ToEventResponse(e)
	}
	return res
}

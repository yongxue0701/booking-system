package dto

import (
	"time"

	"booking-system/internal/data"
)

type FlightDto struct {
	ID                int       `json:"id"`
	Departure         string    `json:"departure"`
	Destination       string    `json:"destination"`
	DepartureTerminal string    `json:"departure_terminal"`
	ArrivalTerminal   string    `json:"arrival_terminal"`
	DepartureDate     time.Time `json:"departure_date"`
	ArrivalDate       time.Time `json:"arrival_date"`
}

func BuildFlightDto(ticketEntity *data.Flight) *FlightDto {
	return &FlightDto{
		ID:                ticketEntity.ID,
		Departure:         ticketEntity.Departure,
		Destination:       ticketEntity.Destination,
		DepartureTerminal: ticketEntity.DepartureTerminal,
		ArrivalTerminal:   ticketEntity.ArrivalTerminal,
		DepartureDate:     ticketEntity.DepartureDate,
		ArrivalDate:       ticketEntity.ArrivalDate,
	}
}

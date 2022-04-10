package dto

import (
	"time"

	"booking-system/internal/data"
)

type BookingDetailsDto struct {
	PassengerName     string    `json:"passenger_name"`
	PassengerPassport string    `json:"passenger_passport"`
	PurchaseDate      time.Time `json:"purchase_date"`
	Departure         string    `json:"departure"`
	Destination       string    `json:"destination"`
	DepartureTerminal string    `json:"departure_terminal"`
	ArrivalTerminal   string    `json:"arrival_terminal"`
	DepartureDate     time.Time `json:"departure_date"`
	ArrivalDate       time.Time `json:"arrival_date"`
}

func BuildBookingDetailsDto(bookingEntity *data.Booking, flightEntity *data.Flight) *BookingDetailsDto {
	return &BookingDetailsDto{
		PassengerName:     bookingEntity.PassengerName,
		PassengerPassport: bookingEntity.PassengerPassport,
		PurchaseDate:      bookingEntity.PurchaseDate,
		Departure:         flightEntity.Departure,
		Destination:       flightEntity.Destination,
		DepartureTerminal: flightEntity.DepartureTerminal,
		ArrivalTerminal:   flightEntity.ArrivalTerminal,
		DepartureDate:     flightEntity.DepartureDate,
		ArrivalDate:       flightEntity.ArrivalDate,
	}
}

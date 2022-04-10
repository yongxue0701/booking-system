package dto

import (
	"time"

	"booking-system/internal/data"
)

type BookingDto struct {
	ID                int       `json:"id"`
	FlightTicketID    int       `json:"flight_ticket_id"`
	PassengerName     string    `json:"passenger_name"`
	PassengerPassport string    `json:"passenger_passport"`
	PurchaseDate      time.Time `json:"purchase_date"`
}

func BuildBookingDto(bookingEntity *data.Booking) *BookingDto {
	return &BookingDto{
		ID:                bookingEntity.ID,
		FlightTicketID:    bookingEntity.FlightTicketID,
		PassengerName:     bookingEntity.PassengerName,
		PassengerPassport: bookingEntity.PassengerPassport,
		PurchaseDate:      bookingEntity.PurchaseDate,
	}
}

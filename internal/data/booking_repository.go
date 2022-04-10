package data

import (
	"fmt"
	"time"
)

type Booking struct {
	ID                int
	FlightTicketID    int
	PassengerName     string
	PassengerPassport string
	PurchaseDate      time.Time
}

type BookingRepository struct {
}

func NewBookingRepository() *BookingRepository {
	return &BookingRepository{}
}

func (b *BookingRepository) Create() (*Booking, error) {
	return &Booking{
		ID:                1,
		FlightTicketID:    1,
		PassengerName:     "YX",
		PassengerPassport: "G1234567H",
		PurchaseDate:      time.Date(2022, time.April, 10, 16, 0, 0, 0, time.UTC),
	}, nil
}

func (b *BookingRepository) GetByID(id int) (*Booking, error) {
	if id != 1 {
		return nil, fmt.Errorf("%s: id %d", "booking is not found", id)
	}

	return &Booking{
		ID:                id,
		FlightTicketID:    1,
		PassengerName:     "YX",
		PassengerPassport: "G1234567H",
		PurchaseDate:      time.Date(2022, time.April, 10, 16, 0, 0, 0, time.UTC),
	}, nil
}

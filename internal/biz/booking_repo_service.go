package biz

import "flight-booking/internal/data"

type BookingRepoService interface {
	Create() (*data.Booking, error)
	GetByID(id int) (*data.Booking, error)
}

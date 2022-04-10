package biz

import "booking-system/internal/data"

type BookingRepoService interface {
	Create() (*data.Booking, error)
	GetByID(id int) (*data.Booking, error)
}

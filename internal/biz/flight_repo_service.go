package biz

import "booking-system/internal/data"

type FlightRepoService interface {
	Create() (*data.Flight, error)
	GetByID(id int) (*data.Flight, error)
}

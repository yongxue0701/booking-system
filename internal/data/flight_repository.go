package data

import (
	"fmt"
	"time"
)

type Flight struct {
	ID                int
	Departure         string
	Destination       string
	DepartureTerminal string
	ArrivalTerminal   string
	DepartureDate     time.Time
	ArrivalDate       time.Time
}

type FlightRepository struct {
}

func NewFlightRepository() *FlightRepository {
	return &FlightRepository{}
}

func (f *FlightRepository) Create() (*Flight, error) {
	return &Flight{
		ID:                1,
		Departure:         "Wuhan",
		Destination:       "Shenzhen",
		DepartureTerminal: "Terminal 3",
		ArrivalTerminal:   "Terminal 1",
		DepartureDate:     time.Date(2022, time.December, 01, 8, 0, 0, 0, time.UTC),
		ArrivalDate:       time.Date(2022, time.December, 01, 10, 0, 0, 0, time.UTC),
	}, nil
}

func (f *FlightRepository) GetByID(id int) (*Flight, error) {
	if id != 1 {
		return nil, fmt.Errorf("%s: id %d", "flight is not found", id)
	}

	return &Flight{
		ID:                id,
		Departure:         "Wuhan",
		Destination:       "Shenzhen",
		DepartureTerminal: "Terminal 3",
		ArrivalTerminal:   "Terminal 1",
		DepartureDate:     time.Date(2022, time.December, 01, 8, 0, 0, 0, time.UTC),
		ArrivalDate:       time.Date(2022, time.December, 01, 10, 0, 0, 0, time.UTC),
	}, nil
}

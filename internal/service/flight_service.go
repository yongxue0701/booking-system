package service

import (
	"booking-system/internal/biz"
	"booking-system/internal/dto"
	"net/http"
)

type FlightService struct {
	flightRepoService biz.FlightRepoService
}

func NewFlightService(flightRepoService biz.FlightRepoService) *FlightService {
	return &FlightService{
		flightRepoService: flightRepoService,
	}
}

func (f *FlightService) Create() (*dto.FlightDto, int) {
	flightEntity, err := f.flightRepoService.Create()
	if err != nil {
		return nil, http.StatusBadRequest
	}
	return dto.BuildFlightDto(flightEntity), http.StatusOK
}

func (f *FlightService) GetByID(id int) (*dto.FlightDto, int) {
	flightEntity, err := f.flightRepoService.GetByID(id)
	if err != nil {
		return nil, http.StatusNotFound
	}
	return dto.BuildFlightDto(flightEntity), http.StatusOK
}

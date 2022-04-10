package service

import (
	"net/http"

	"booking-system/internal/biz"
	"booking-system/internal/dto"
)

type BookingService struct {
	bookingRepoService biz.BookingRepoService
	flightRepoService  biz.FlightRepoService
}

func NewBookingService(bookingRepoService biz.BookingRepoService, flightRepoService biz.FlightRepoService) *BookingService {
	return &BookingService{
		bookingRepoService: bookingRepoService,
		flightRepoService:  flightRepoService,
	}
}

func (b *BookingService) Create() (*dto.BookingDto, int) {
	bookingEntity, err := b.bookingRepoService.Create()
	if err != nil {
		return nil, http.StatusBadRequest
	}
	return dto.BuildBookingDto(bookingEntity), http.StatusOK
}

func (b *BookingService) GetByID(id int) (*dto.BookingDetailsDto, int) {
	bookingEntity, err := b.bookingRepoService.GetByID(id)
	if err != nil {
		return nil, http.StatusNotFound
	}
	flightTicketID := bookingEntity.FlightTicketID

	ticketEntity, err := b.flightRepoService.GetByID(flightTicketID)
	if err != nil {
		return nil, http.StatusNotFound
	}

	return dto.BuildBookingDetailsDto(bookingEntity, ticketEntity), http.StatusOK
}

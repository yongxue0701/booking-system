package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"flight-booking/internal/data"
	"flight-booking/internal/service"
)

func createBooking(w http.ResponseWriter, r *http.Request) {
	bookingRepo := data.NewBookingRepository()
	flightRepo := data.NewFlightRepository()

	bookingService := service.NewBookingService(bookingRepo, flightRepo)
	result, statusCode := bookingService.Create()
	message := getMessage(statusCode)

	response := &Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       result,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		println("failed to marshal http response")
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func getBookingByID(w http.ResponseWriter, r *http.Request) {
	idFromQuery := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idFromQuery)
	if err != nil {
		println("failed to convert string to int")
	}

	bookingRepo := data.NewBookingRepository()
	flightRepo := data.NewFlightRepository()

	bookingService := service.NewBookingService(bookingRepo, flightRepo)
	result, statusCode := bookingService.GetByID(id)
	message := getMessage(statusCode)

	response := &Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       result,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		println("failed to marshal http response")
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

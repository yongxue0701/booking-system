package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"booking-system/internal/data"
	"booking-system/internal/service"
)

func createFlight(w http.ResponseWriter, r *http.Request) {
	flightRepo := data.NewFlightRepository()
	flightService := service.NewFlightService(flightRepo)

	result, statusCode := flightService.Create()
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

func getFlightByID(w http.ResponseWriter, r *http.Request) {
	idFromQuery := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idFromQuery)
	if err != nil {
		println("failed to convert string to int")
	}

	flightRepo := data.NewFlightRepository()
	flightService := service.NewFlightService(flightRepo)

	result, statusCode := flightService.GetByID(id)
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

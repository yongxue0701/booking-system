package api

import "net/http"

func HandleRequests() {
	http.HandleFunc("/booking/create", createBooking)
	http.HandleFunc("/booking/get", getBookingByID)
	http.HandleFunc("/flight/create", createFlight)
	http.HandleFunc("/flight/get", getFlightByID)
	http.ListenAndServe(":8080", nil)
}

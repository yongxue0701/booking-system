package api

import (
	"fmt"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome!")
	})
	http.HandleFunc("/booking/create", createBooking)
	http.HandleFunc("/booking/get", getBookingByID)
	http.HandleFunc("/flight/create", createFlight)
	http.HandleFunc("/flight/get", getFlightByID)
	http.ListenAndServe(":8080", nil)
}

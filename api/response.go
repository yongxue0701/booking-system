package api

import "net/http"

type Response struct {
	StatusCode int
	Message    string
	Data       interface{}
}

func getMessage(statusCode int) string {
	var message string
	if statusCode == http.StatusOK {
		message = "Successful"
	} else if statusCode == http.StatusNotFound {
		message = "Data is not found"
	} else {
		message = "Failed to create booking"
	}

	return message
}

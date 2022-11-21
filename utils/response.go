package utils

import "github.com/theHinneh/labs-service-booking/models"

type AppointmentResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    *models.Appointment
}

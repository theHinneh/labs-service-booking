package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/theHinneh/labs-service-booking/db"
	"github.com/theHinneh/labs-service-booking/models"
	"github.com/theHinneh/labs-service-booking/utils"
	"net/http"
	"regexp"
	"strconv"
)

func GetAllAppointments(context *gin.Context) {
	var appointments []models.Appointment
	var errorResponse utils.AppointmentErrorResponse

	result := db.DB.Raw("SELECT * FROM Appointment;").Scan(&appointments)

	if &result.Error != nil {
		errorResponse.Status = "error"
		errorResponse.Data = nil
		errorResponse.Message = "Not found"

		context.JSONP(http.StatusNotFound, &errorResponse)
	} else {
		context.JSONP(http.StatusOK, gin.H{
			"message": "Appointments",
			"status":  "Success",
			"data":    appointments,
		})
	}
}

func GetAnAppointment(context *gin.Context) {
	var appointment models.Appointment
	appointmentId, _ := strconv.Atoi(context.Param("id"))
	var errorResponse utils.AppointmentErrorResponse

	query := "SELECT * FROM Appointment WHERE appointment_id = ?"
	result := db.DB.Raw(query, appointmentId).Scan(&appointment)

	if &result.Error != nil {
		errorResponse.Status = "error"
		errorResponse.Data = nil
		errorResponse.Message = "Not found"

		context.JSONP(http.StatusNotFound, &errorResponse)
	} else {
		context.JSONP(http.StatusOK, gin.H{
			"message": "Appointments",
			"status":  "Success",
			"data":    appointment,
		})
	}
}

func GetPendingAppointments(context *gin.Context) {
	var appointments []models.Appointment
	var errorResponse utils.AppointmentErrorResponse

	query := "SELECT * FROM Appointment WHERE completed = ?"
	result := db.DB.Raw(query, 0).Scan(&appointments)

	if &result.Error != nil {
		errorResponse.Status = "error"
		errorResponse.Data = nil
		errorResponse.Message = "Not found"

		context.JSONP(http.StatusNotFound, &errorResponse)
	} else {
		context.JSONP(http.StatusOK, gin.H{
			"message": "Appointments",
			"status":  "Success",
			"data":    appointments,
		})
	}
}

func GetCompletedAppointments(context *gin.Context) {
	var appointments []models.Appointment
	var errorResponse utils.AppointmentErrorResponse

	query := "SELECT * FROM Appointment WHERE completed = ?"
	result := db.DB.Raw(query, 1).Scan(&appointments)

	if &result.Error != nil {
		errorResponse.Status = "error"
		errorResponse.Data = nil
		errorResponse.Message = "Not found"

		context.JSONP(http.StatusNotFound, &errorResponse)
	} else {
		context.JSONP(http.StatusOK, gin.H{
			"message": "Appointments",
			"status":  "Success",
			"data":    appointments,
		})
	}
}

func AddAppointment(context *gin.Context) {
	var apt models.Appointment
	var errorResponse utils.AppointmentErrorResponse

	if err := context.ShouldBindJSON(&apt); err != nil {
		errorResponse.Status = "error"
		errorResponse.Message = "An error occurred while binding json"
		errorResponse.Data = nil
		context.JSONP(http.StatusUnprocessableEntity, &errorResponse)
		return
	}

	match, _ := regexp.MatchString("(\\d{4})-(\\d{2})-(\\d{2}) (\\d{2}):(\\d{2}):(\\d{2})", apt.AppointmentDate)

	if !match {
		errorResponse.Status = "error"
		errorResponse.Message = "Date format is invalid!"
		errorResponse.Data = nil
		context.JSONP(http.StatusBadRequest, &errorResponse)
		return
	}

	query := "INSERT INTO Appointment(client_name, appointment_date, client_email, note, completed, service_id, shop_id, staff_id, contact_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result := db.DB.Raw(query, apt.ClientName, apt.AppointmentDate, apt.ClientEmail, apt.Note, apt.Completed, apt.ServiceId, apt.ShopId, apt.StaffId, apt.ContactId).Error

	if result != nil {
		errorResponse.Status = "error"
		errorResponse.Message = "An error occurred"
		errorResponse.Data = nil
		context.JSON(http.StatusBadRequest, &errorResponse)
		return
	}

	context.JSONP(http.StatusOK, gin.H{
		"message": "Appointments",
		"status":  "Success",
		"data":    apt,
	})
}

func DeleteAppointment(context *gin.Context) {
	appointmentId, _ := strconv.Atoi(context.Param("id"))
	var errorResponse utils.AppointmentErrorResponse

	query := "DELETE FROM Appointment WHERE appointment_id = ?"
	result := db.DB.Raw(query, appointmentId).Error

	if result != nil {
		errorResponse.Status = "error"
		errorResponse.Message = "An error occurred"
		errorResponse.Data = nil
		context.JSON(http.StatusBadRequest, &errorResponse)
		return
	}

	context.JSONP(http.StatusOK, gin.H{
		"message": "Appointment Deleted",
		"status":  "Success",
		"data":    appointmentId,
	})
}

func UpdateAppointment(context *gin.Context) {
	var apt models.Appointment
	var errorResponse utils.AppointmentErrorResponse
	appointmentId, _ := strconv.Atoi(context.Param("id"))

	if err := context.ShouldBindJSON(&apt); err != nil {
		errorResponse.Status = "error"
		errorResponse.Message = "An error occurred while binding json"
		errorResponse.Data = nil
		context.JSONP(http.StatusUnprocessableEntity, &errorResponse)
		return
	}

	match, _ := regexp.MatchString("(\\d{4})-(\\d{2})-(\\d{2}) (\\d{2}):(\\d{2}):(\\d{2})", apt.AppointmentDate)

	if !match {
		errorResponse.Status = "error"
		errorResponse.Message = "Date format is invalid!"
		errorResponse.Data = nil
		context.JSONP(http.StatusBadRequest, &errorResponse)
		return
	}

	query := "UPDATE appointment SET client_name = ?, appointment_date = ?, client_email = ?, note = ?, completed = ?, service_id = ?, shop_id = ?, staff_id = ?, contact_id = ? WHERE appointment_id = ?"
	result := db.DB.Raw(query, apt.ClientName, apt.AppointmentDate, apt.ClientEmail, apt.Note, apt.Completed, apt.ServiceId, apt.ShopId, apt.StaffId, apt.ContactId, appointmentId)

	apt.AppointmentId = appointmentId

	if result.Error != nil {
		errorResponse.Status = "error"
		errorResponse.Message = "An error occurred"
		errorResponse.Data = nil
		context.JSON(http.StatusBadRequest, &errorResponse)
		return
	}

	context.JSONP(http.StatusOK, gin.H{
		"message": "Appointment Updated",
		"status":  "Success",
		"data":    apt,
	})
}

func GetAppointmentStaffDetails(context *gin.Context) {
	var appointment models.AppointmentStaffData
	appointmentId, _ := strconv.Atoi(context.Param("id"))
	var errorResponse utils.AppointmentErrorResponse

	query := "SELECT shop.shop_name Shop, CONCAT(s.last_name, ' ', s.first_name) Staff, s2.service_name Service, c.contact FROM shop LEFT JOIN appointment a on shop.shop_id = a.shop_id LEFT JOIN staff s on shop.shop_id = s.shop_id LEFT JOIN services s2 on s.staff_id = s2.staff_id LEFT OUTER JOIN contacts c on c.contact_id = s.contact_id WHERE appointment_id = ?;"
	result := db.DB.Debug().Raw(query, appointmentId).Scan(&appointment)

	if result.Error != nil {
		errorResponse.Status = "error"
		errorResponse.Data = nil
		errorResponse.Message = "Not found"

		context.JSONP(http.StatusNotFound, &errorResponse)
	} else {
		context.JSONP(http.StatusOK, gin.H{
			"message": "Appointments",
			"status":  "Success",
			"data":    appointment,
		})
	}
}

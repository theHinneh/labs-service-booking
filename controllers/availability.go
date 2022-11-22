package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/theHinneh/labs-service-booking/db"
	"github.com/theHinneh/labs-service-booking/models"
	"github.com/theHinneh/labs-service-booking/utils"
	"net/http"
	"regexp"
	"strconv"
)

func GetAllAvailability(context *gin.Context) {
	var availability []models.Availability
	var errorResponse utils.AppointmentErrorResponse

	result := db.DB.Raw("SELECT availability_id, date availability_date, staff_id FROM Availability;").Scan(&availability)

	if result.Error != nil {
		errorResponse.Status = "error"
		errorResponse.Data = nil
		errorResponse.Message = "Not found"

		context.JSONP(http.StatusNotFound, &errorResponse)
	} else {
		context.JSONP(http.StatusOK, gin.H{
			"message": "Availability",
			"status":  "Success",
			"data":    availability,
		})
	}
}

func GetAnAvailability(context *gin.Context) {
	var availability models.Availability
	var errorResponse utils.AppointmentErrorResponse
	availabilityId, _ := strconv.Atoi(context.Param("id"))

	query := "SELECT availability_id, date availability_date, staff_id FROM Availability WHERE availability_id = ?"
	result := db.DB.Raw(query, availabilityId).Scan(&availability)

	availability.AvailabilityId = availabilityId

	if result.Error != nil {
		errorResponse.Status = "error"
		errorResponse.Data = nil
		errorResponse.Message = "Not found"

		context.JSONP(http.StatusNotFound, &errorResponse)
	} else {
		context.JSONP(http.StatusOK, gin.H{
			"message": "Availability",
			"status":  "Success",
			"data":    availability,
		})
	}
}

func AddAvailability(context *gin.Context) {

	var availability models.Availability
	var errorResponse utils.AppointmentErrorResponse

	if err := context.ShouldBindJSON(&availability); err != nil {
		errorResponse.Status = "error"
		errorResponse.Message = "An error occurred while binding json"
		errorResponse.Data = nil
		context.JSONP(http.StatusUnprocessableEntity, &errorResponse)
		return
	}

	match, _ := regexp.MatchString("(\\d{4})-(\\d{2})-(\\d{2}) (\\d{2}):(\\d{2}):(\\d{2})", availability.AvailabilityDate)

	if !match {
		errorResponse.Status = "error"
		errorResponse.Message = "Date format is invalid!"
		errorResponse.Data = nil
		context.JSONP(http.StatusBadRequest, &errorResponse)
		return
	}

	query := "INSERT INTO Availability(date, staff_id) VALUES (?, ?)"
	result := db.DB.Debug().Raw(query, availability.AvailabilityDate, availability.StaffId).Scan(&availability)

	if result.Error != nil {
		errorResponse.Status = "error"
		errorResponse.Message = "An error occurred"
		errorResponse.Data = nil
		context.JSON(http.StatusBadRequest, &errorResponse)
		return
	}

	type data struct {
		AvailabilityDate string `json:"availability_date" binding:"required"`
		StaffId          int    `json:"staff_id" binding:"required"`
	}

	d := data{StaffId: availability.StaffId, AvailabilityDate: availability.AvailabilityDate}

	context.JSONP(http.StatusOK, gin.H{
		"message": "Added Availability",
		"status":  "Success",
		"data":    &d,
	})
}

func DeleteAvailability(context *gin.Context) {
	availabilityId, _ := strconv.Atoi(context.Param("id"))
	var errorResponse utils.AppointmentErrorResponse

	query := "DELETE FROM Availability WHERE availability_id = ?"
	result := db.DB.Raw(query, availabilityId).Error

	if result != nil {
		errorResponse.Status = "error"
		errorResponse.Message = "An error occurred"
		errorResponse.Data = nil
		context.JSON(http.StatusBadRequest, &errorResponse)
		return
	}

	context.JSONP(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Availability %d Deleted", availabilityId),
		"status":  "Success",
		"data":    availabilityId,
	})
}

func UpdateAvailability(context *gin.Context) {
	var availability models.Availability
	var errorResponse utils.AppointmentErrorResponse
	appointmentId, _ := strconv.Atoi(context.Param("id"))

	if err := context.ShouldBindJSON(&availability); err != nil {
		errorResponse.Status = "error"
		errorResponse.Message = "An error occurred while binding json"
		errorResponse.Data = nil
		context.JSONP(http.StatusUnprocessableEntity, &errorResponse)
		return
	}

	match, _ := regexp.MatchString("(\\d{4})-(\\d{2})-(\\d{2}) (\\d{2}):(\\d{2}):(\\d{2})", availability.AvailabilityDate)

	if !match {
		errorResponse.Status = "error"
		errorResponse.Message = "Date format is invalid!"
		errorResponse.Data = nil
		context.JSONP(http.StatusBadRequest, &errorResponse)
		return
	}

	query := "UPDATE Availability SET date = ?, staff_id = ? WHERE availability_id = ?"
	result := db.DB.Raw(query, availability.AvailabilityDate, availability.StaffId)

	availability.AvailabilityId = appointmentId

	if result.Error != nil {
		errorResponse.Status = "error"
		errorResponse.Message = "An error occurred"
		errorResponse.Data = nil
		context.JSON(http.StatusBadRequest, &errorResponse)
		return
	}

	context.JSONP(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Availability %d Updated", appointmentId),
		"status":  "Success",
		"data":    availability,
	})
}

func GetAvailabilityDetails(context *gin.Context) {
	var availability models.AvailabilityDetailsResponse
	availabilityId, _ := strconv.Atoi(context.Param("id"))
	var errorResponse utils.AppointmentErrorResponse

	query := `
			SELECT s.email StaffEmail, CONCAT(s.first_name, ' ', s.last_name) StaffName, a.date AvailabilityDate
			FROM staff s
					 LEFT JOIN availability a on s.staff_id = a.staff_id
			WHERE a.availability_id = ?;
`
	result := db.DB.Raw(query, availabilityId).Scan(&availability)

	if result.Error != nil {
		errorResponse.Status = "error"
		errorResponse.Data = nil
		errorResponse.Message = "Not found"

		context.JSONP(http.StatusNotFound, &errorResponse)
	} else {
		context.JSONP(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Availability %d Details", availabilityId),
			"status":  "Success",
			"data":    availability,
		})
	}
}

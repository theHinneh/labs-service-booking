package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/theHinneh/labs-service-booking/db"
	"github.com/theHinneh/labs-service-booking/models"
	"github.com/theHinneh/labs-service-booking/utils"
	"net/http"
	"strconv"
)

func GetServices(context *gin.Context) {
	var services []models.Service
	var errorResponse utils.AppointmentErrorResponse

	result := db.DB.Raw("SELECT * FROM Services;").Scan(&services)

	if result.Error != nil {
		errorResponse.Status = "error"
		errorResponse.Data = nil
		errorResponse.Message = "Not found"

		context.JSONP(http.StatusNotFound, &errorResponse)
	} else {
		context.JSONP(http.StatusOK, gin.H{
			"message": "Services",
			"status":  "Success",
			"data":    services,
		})
	}
}

func GetAService(context *gin.Context) {
	var service models.Service
	var errorResponse utils.AppointmentErrorResponse
	serviceId, _ := strconv.Atoi(context.Param("id"))

	query := "SELECT * FROM Services WHERE service_id = ?;"
	result := db.DB.Raw(query, serviceId).Scan(&service)

	if result.Error != nil {
		errorResponse.Status = "error"
		errorResponse.Data = nil
		errorResponse.Message = "Not found"

		context.JSONP(http.StatusNotFound, &errorResponse)
	} else {
		context.JSONP(http.StatusOK, gin.H{
			"message": "Service",
			"status":  "Success",
			"data":    service,
		})
	}
}

func AddService(context *gin.Context) {
	var service models.Service
	var errorResponse utils.AppointmentErrorResponse

	if err := context.ShouldBindJSON(&service); err != nil {
		errorResponse.Status = "error"
		errorResponse.Message = "An error occurred while binding json"
		errorResponse.Data = nil
		context.JSONP(http.StatusUnprocessableEntity, &errorResponse)
		return
	}

	query := "INSERT INTO Services(start_price, end_price, staff_id, service_name, shop_id) VALUES (?, ?, ?, ?, ?)"
	result := db.DB.Raw(query, service.StartPrice, service.EndPrice, service.StaffId, service.ServiceName, service.ShopId).Scan(&service)

	if result.Error != nil {
		errorResponse.Status = "error"
		errorResponse.Data = nil
		errorResponse.Message = "Not found"
		context.JSONP(http.StatusNotFound, &errorResponse)
		return
	}

	type data struct {
		StartPrice  float32 `json:"start_price"`
		EndPrice    float32 `json:"end_price"`
		StaffId     int     `json:"staff_id"`
		ServiceName string  `json:"service_name"`
		ShopId      int     `json:"shop_id"`
	}

	d := data{StartPrice: service.StartPrice, EndPrice: service.EndPrice, StaffId: service.StaffId, ServiceName: service.ServiceName, ShopId: service.ShopId}

	context.JSONP(http.StatusOK, gin.H{
		"message": "Contact Added",
		"status":  "Success",
		"data":    &d,
	})
}

func DeleteService(context *gin.Context) {
	serviceId, _ := strconv.Atoi(context.Param("id"))
	var errorResponse utils.AppointmentErrorResponse

	query := "DELETE FROM Services WHERE service_id = ?"
	result := db.DB.Raw(query, serviceId).Error

	if result != nil {
		errorResponse.Status = "error"
		errorResponse.Message = "An error occurred"
		errorResponse.Data = nil
		context.JSON(http.StatusBadRequest, &errorResponse)
		return
	}

	context.JSONP(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Service %d Deleted", serviceId),
		"status":  "Success",
		"data":    serviceId,
	})
}

func UpdateService(context *gin.Context) {
	var service models.Service
	var errorResponse utils.AppointmentErrorResponse
	serviceId, _ := strconv.Atoi(context.Param("id"))

	if err := context.ShouldBindJSON(&service); err != nil {
		errorResponse.Status = "error"
		errorResponse.Message = "An error occurred while binding json"
		errorResponse.Data = nil
		context.JSONP(http.StatusUnprocessableEntity, &errorResponse)
		return
	}

	query := "UPDATE Services SET end_price = ?, start_price = ?, staff_id = ?, service_name = ?, shop_id = ? WHERE service_id = ?"
	result := db.DB.Raw(query, service.EndPrice, service.StartPrice, service.StaffId, service.ServiceName, service.ShopId, serviceId)

	service.ServiceId = serviceId

	if result.Error != nil {
		errorResponse.Status = "error"
		errorResponse.Message = "An error occurred"
		errorResponse.Data = nil
		context.JSON(http.StatusBadRequest, &errorResponse)
		return
	}

	context.JSONP(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Service %d Updated", serviceId),
		"status":  "Success",
		"data":    service,
	})
}

func GetServiceDetails(context *gin.Context) {
	var service models.ServiceDetails
	appointmentId, _ := strconv.Atoi(context.Param("id"))
	var errorResponse utils.AppointmentErrorResponse

	query := `
			SELECT s.service_id,
				   s.start_price,
				   s.end_price,
				   s.staff_id,
				   s.service_name,
				   a.appointment_date,
				   s2.shop_name,
				   s2.country,
				   CONCAT(s3.first_name, ' ', s3.last_name) StaffName
			FROM Services s
					 LEFT JOIN appointment a on s.service_id = a.service_id
					 LEFT JOIN shop s2 on s2.shop_id = a.shop_id
					 LEFT JOIN staff s3 on s3.staff_id = a.staff_id
			WHERE s.shop_id = ?;
`
	result := db.DB.Raw(query, appointmentId).Scan(&service)

	if result.Error != nil {
		errorResponse.Status = "error"
		errorResponse.Data = nil
		errorResponse.Message = "Not found"

		context.JSONP(http.StatusNotFound, &errorResponse)
	} else {
		context.JSONP(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Service %d Details", appointmentId),
			"status":  "Success",
			"data":    service,
		})
	}
}

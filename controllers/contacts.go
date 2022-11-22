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

func GetContacts(context *gin.Context) {
	var contacts []models.Contacts
	var errorResponse utils.AppointmentErrorResponse

	result := db.DB.Raw("SELECT * FROM Contacts;").Scan(&contacts)

	if result.Error != nil {
		errorResponse.Status = "error"
		errorResponse.Data = nil
		errorResponse.Message = "Not found"

		context.JSONP(http.StatusNotFound, &errorResponse)
	} else {
		context.JSONP(http.StatusOK, gin.H{
			"message": "Contacts",
			"status":  "Success",
			"data":    contacts,
		})
	}
}

func GetAContact(context *gin.Context) {
	var contact models.Contacts
	var errorResponse utils.AppointmentErrorResponse
	contactId, _ := strconv.Atoi(context.Param("id"))

	query := "SELECT * FROM Contacts WHERE contact_id = ?;"
	result := db.DB.Raw(query, contactId).Scan(&contact)

	if result.Error != nil {
		errorResponse.Status = "error"
		errorResponse.Data = nil
		errorResponse.Message = "Not found"

		context.JSONP(http.StatusNotFound, &errorResponse)
	} else {
		context.JSONP(http.StatusOK, gin.H{
			"message": "Contact",
			"status":  "Success",
			"data":    contact,
		})
	}
}

func AddContact(context *gin.Context) {
	var contact models.Contacts
	var errorResponse utils.AppointmentErrorResponse

	if err := context.ShouldBindJSON(&contact); err != nil {
		errorResponse.Status = "error"
		errorResponse.Message = "An error occurred while binding json"
		errorResponse.Data = nil
		context.JSONP(http.StatusUnprocessableEntity, &errorResponse)
		return
	}

	query := "INSERT INTO Contacts(contact) VALUES (?)"
	result := db.DB.Raw(query, contact.Contact).Scan(&contact)

	if result.Error != nil {
		errorResponse.Status = "error"
		errorResponse.Data = nil
		errorResponse.Message = "Not found"
		context.JSONP(http.StatusNotFound, &errorResponse)
		return
	}

	type data struct {
		Contact string `json:"contact" binding:"required"`
	}

	d := data{Contact: contact.Contact}

	context.JSONP(http.StatusOK, gin.H{
		"message": "Contact Added",
		"status":  "Success",
		"data":    &d,
	})
}

func DeleteContact(context *gin.Context) {
	contactId, _ := strconv.Atoi(context.Param("id"))
	var errorResponse utils.AppointmentErrorResponse

	query := "DELETE FROM Contacts WHERE contact_id = ?"
	result := db.DB.Raw(query, contactId).Error

	if result != nil {
		errorResponse.Status = "error"
		errorResponse.Message = "An error occurred"
		errorResponse.Data = nil
		context.JSON(http.StatusBadRequest, &errorResponse)
		return
	}

	context.JSONP(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Contact %d Deleted", contactId),
		"status":  "Success",
		"data":    contactId,
	})
}

func UpdateContact(context *gin.Context) {
	var contact models.Contacts
	var errorResponse utils.AppointmentErrorResponse
	contactId, _ := strconv.Atoi(context.Param("id"))

	if err := context.ShouldBindJSON(&contact); err != nil {
		errorResponse.Status = "error"
		errorResponse.Message = "An error occurred while binding json"
		errorResponse.Data = nil
		context.JSONP(http.StatusUnprocessableEntity, &errorResponse)
		return
	}

	query := "UPDATE Contacts SET contact = ? WHERE contact_id = ?"
	result := db.DB.Raw(query, contact.Contact, contactId)

	contact.ContactId = contactId

	if result.Error != nil {
		errorResponse.Status = "error"
		errorResponse.Message = "An error occurred"
		errorResponse.Data = nil
		context.JSON(http.StatusBadRequest, &errorResponse)
		return
	}

	context.JSONP(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Contact %d Updated", contactId),
		"status":  "Success",
		"data":    contact,
	})
}

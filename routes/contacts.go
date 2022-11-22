package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/theHinneh/labs-service-booking/controllers"
)

func Contacts(router *gin.Engine) {
	v1 := router.Group("/v1/contacts")
	{
		v1.GET("/all", controllers.GetContacts)
		v1.GET("/:id", controllers.GetAContact)
		v1.POST("/add", controllers.AddContact)
		v1.DELETE("/:id", controllers.DeleteContact)
		v1.PATCH("/:id", controllers.UpdateContact)
	}
}

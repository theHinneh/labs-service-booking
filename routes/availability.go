package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/theHinneh/labs-service-booking/controllers"
)

func AvailabilityRoutes(router *gin.Engine) {
	v1 := router.Group("/v1/availability")
	{
		v1.GET("/all", controllers.GetAllAvailability)
		v1.GET("/:id", controllers.GetAnAvailability)
		v1.POST("/add", controllers.AddAvailability)
		v1.DELETE("/:id", controllers.DeleteAvailability)
		v1.PATCH("/:id", controllers.UpdateAvailability)
		v1.GET("/availability/:id", controllers.GetAvailabilityDetails)
	}
}

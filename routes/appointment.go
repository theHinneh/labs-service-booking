package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/theHinneh/labs-service-booking/controllers"
)

func AppointmentRoutes(router *gin.Engine) {
	v1 := router.Group("/v1/appointments")
	{
		v1.GET("/all", controllers.GetAllAppointments)
		v1.GET("/:id", controllers.GetAnAppointment)
		v1.GET("/all/pending", controllers.GetPendingAppointments)
		v1.GET("/all/completed", controllers.GetCompletedAppointments)
		v1.POST("/add", controllers.AddAppointment)
		v1.DELETE("/:id", controllers.DeleteAppointment)
		v1.PATCH("/:id", controllers.UpdateAppointment)
	}
}

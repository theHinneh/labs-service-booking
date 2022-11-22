package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/theHinneh/labs-service-booking/controllers"
)

func ServicesRoutes(router *gin.Engine) {
	v1 := router.Group("/v1/services")
	{
		v1.GET("/all", controllers.GetServices)
		v1.GET("/:id", controllers.GetAService)
		v1.POST("/add", controllers.AddService)
		v1.DELETE("/:id", controllers.DeleteService)
		v1.PATCH("/:id", controllers.UpdateService)
		v1.GET("/service/:id", controllers.GetServiceDetails)
	}
}

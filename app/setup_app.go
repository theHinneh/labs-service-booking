package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/theHinneh/labs-service-booking/db"
	"github.com/theHinneh/labs-service-booking/middlewares"
	"github.com/theHinneh/labs-service-booking/routes"
)

func SetupApp() *gin.Engine {
	log.Info().Msg("Initializing service")

	// setup database
	db.ConnectDb()

	//create engine
	app := gin.New()
	//add default recovery middleware
	app.Use(gin.Recovery())

	// disabling the trusted proxy feature
	app.SetTrustedProxies(nil)

	// Add cors, request ID and request logging middleware
	log.Info().Msg("Adding cors, request id and request logging middleware")
	app.Use(middlewares.CORSMiddleware(), middlewares.RequestID(), middlewares.RequestLogger())

	// Setup routers
	log.Info().Msg("Setting up routers")
	routes.AppointmentRoutes(app)
	routes.AvailabilityRoutes(app)
	routes.ContactsRoutes(app)
	routes.ServicesRoutes(app)

	return app
}

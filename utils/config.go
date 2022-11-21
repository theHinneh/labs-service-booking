package utils

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/theHinneh/labs-service-booking/constants"
)

func init() {
	viper.AddConfigPath(constants.EnvFileDirectory)
	viper.SetConfigName("service_booking")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Debug().Err(err).
			Msg("Error occurred while reading env file, might fallback to OS env config")
	}
	viper.AutomaticEnv()
}

// GetEnvVar This function can be used to get ENV Var in our App
// Modify this if you change the library to read ENV
func GetEnvVar(name string) string {
	if !viper.IsSet(name) {
		log.Debug().Msgf("Environment variable %s is not set", name)
		return ""
	}
	value := viper.GetString(name)
	return value
}

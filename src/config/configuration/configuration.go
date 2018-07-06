package configuration

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfiguration() {
	log.Printf("reading configuration")
	viper.AddConfigPath("/home/andres/prueba/a/")
	viper.AddConfigPath("../")

	viper.SetConfigFile("../configFiles.json")
	viper.AutomaticEnv()
	error := viper.ReadInConfig()

	if error != nil {
		panic(error)
	}
	value := viper.Get("port")

	log.Printf("el puerto leido de viper es %v", value)

}

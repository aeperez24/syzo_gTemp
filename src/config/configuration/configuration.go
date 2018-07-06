package configuration

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfiguration() {
	log.Printf("reading configuration")
	viper.AddConfigPath("/home/andres/prueba/a/")
	viper.AddConfigPath("../")

	viper.SetConfigFile("/home/andres/Documents/golanggit/syzo_gTemp/configFiles.json")
	viper.AutomaticEnv()
	error := viper.ReadInConfig()

	if error != nil {
		panic(error)
	}
	value := viper.Get("port")
	ReadQuerysFile()
	log.Printf("el puerto leido de viper es %v", value)

}

func ReadQuerysFile() {
	log.Printf("reading queries")

	viper.SetConfigFile("/home/andres/Documents/golanggit/syzo_gTemp/queryFiles.json")
	viper.AddConfigPath("../")
	viper.SetConfigType("json")
	error := viper.MergeInConfig()

	if error != nil {
		panic(error)
	}
	query := viper.Get("query.mongo.temperature.dateStartDateEnd")

	log.Printf(" query leida es %v", query)

}

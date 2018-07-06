package main

import (
	dlvr "temperature/delivery"
	usrcase "temperature/useCase"

	tempRepository "temperature/repository"

	"log"

	config "config/configuration"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func init() {

}
func main() {

	config.InitConfiguration()
	e := echo.New()
	repository := tempRepository.NewTemperatureRepository()

	scase := usrcase.NewTemperatureUseCase(repository)

	//handler init injecting useCase "scase"
	dlvr.NewMesajeHttpHandler(e, scase)
	e.Start(viper.GetString("port"))
	log.Print("started")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
}

package repository

import (
	"config/configuration"
	"fmt"
	"log"
	"testing"
	"time"
)

const (
	myLoc = "America/Argentina/Buenos_Aires"
	ny    = "America/New_York"
)

func TestTemperature(t *testing.T) {
	t.Log("accaaaa")
	configuration.InitConfiguration()
	repo := NewTemperatureRepository()
	// timeAux, _ := time.Parse("2006-01-02 15:04 -0700", "2011-01-19 16:00 -0400")
	// repo.SaveTemperature(temperature.Temperature{Date: timeAux, Measurement: 25.1})
	begin, _ := time.Parse("2006-01-02 15:04 -0700", "2010-01-10 16:00 -0400")
	end := time.Now()

	x := repo.GetTemperature(begin, end)

	result := fmt.Sprintf(" value %v", x)
	log.Printf("%v", result)
	for _, element := range x {
		tzone, _ := time.LoadLocation(myLoc)
		// fmt.Printf("local tzone is %v", tzone)
		auxTime := element.Date.In(tzone)
		log.Printf("recovered :%v ", auxTime)
	}

	// t.Log(result)
	t.Log("finalizado")

}

package repository

import (
	temperature "temperature"
	"time"
)

type TemperatureRepository interface {
	SaveTemperature(temperature.Temperature)
	GetTemperature(start time.Time, end time.Time) []temperature.Temperature
}

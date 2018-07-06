package useCase

import (
	"temperature/repository"
	"testing"
	"time"
)

const (
	myLoc = "America/Argentina/Buenos_Aires"
	ny    = "America/New_York"
)

func TestTemperature(t *testing.T) {
	t.Log("probando caso de uso")
	scase := tempertureUseCase{repository: repository.NewTemperatureRepository()}
	scase.StartReading()
	time.Sleep(time.Duration(5) * time.Second)
	scase.StopReading()
	time.Sleep(time.Duration(1) * time.Second)

}

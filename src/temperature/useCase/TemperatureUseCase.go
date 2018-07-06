package useCase

import (
	"encoding/json"
	"fmt"
	"log"
	temp "temperature"
	repo "temperature/repository"
	"time"

	"github.com/spf13/viper"
)

type TemperatureUseCase interface {
	StartReading()
	GetTemperatureRecords(startAt string, endAt string) []temp.Temperature
	GetTemperatureByDay(t time.Time) []temp.Temperature

	StopReading()
}

type tempertureUseCase struct {
	repository repo.TemperatureRepository
	reading    bool
}

func NewTemperatureUseCase(repository repo.TemperatureRepository) TemperatureUseCase {
	return &tempertureUseCase{repository: repository}

}

func (tcase *tempertureUseCase) StartReading() {
	if !tcase.reading {
		tcase.reading = true

		aux_func := func() {
			for tcase.reading {
				readedTemp := readSerial()
				readedTemp.Date = time.Now()

				// tcase.repository.SaveTemperature(readedTemp)
				time.Sleep(time.Duration(1) * time.Second)

			}
			fmt.Printf("outside the for loop, stoping reading")
		}
		go aux_func()

	} else {
		fmt.Printf("cant start reading, because it's already reading")
	}
}

func (tcase tempertureUseCase) GetTemperatureRecords(startAt string, endAt string) []temp.Temperature {
	timeFormat := viper.GetString("timeFormat")
	begin, error1 := time.Parse(timeFormat, startAt)
	end, error2 := time.Parse(timeFormat, endAt)
	if error1 != nil {
		log.Printf("error al parsear en fecha de inicio %v", error1)
	}
	if error2 != nil {
		log.Printf("error al parsear en fecha de fin %v", error2)

	}
	log.Printf(" en el caso de uso begin %v end %v", begin, end)
	return tcase.repository.GetTemperature(begin, end)
}

func (tcase *tempertureUseCase) StopReading() {
	tcase.reading = false

}

func readSerial() temp.Temperature {
	serialPort := viper.GetString("seriaPort")
	baudRate := viper.GetInt("baudrate")
	readValue := readSerialService(serialPort, baudRate, "\n")
	result := temp.Temperature{}
	fmt.Printf("el json es:%v", string(readValue))

	fmt.Println("")
	err := json.Unmarshal(readValue, &result)
	if err != nil {
		fmt.Printf("error %v", err)
	}
	//	json.Unmarshal(readValue, &result)
	return result
}

type prueba struct {
	Measurement int
}

func (tcase tempertureUseCase) GetTemperatureByDay(day time.Time) []temp.Temperature {
	end := time.Date(day.Year(), day.Month(), day.Day(), 23, 59, 59, 59, day.Location())
	timeFormat := viper.GetString("timeFormat")

	startString := day.Format(timeFormat)
	endString := end.Format(timeFormat)
	return tcase.GetTemperatureRecords(startString, endString)
}

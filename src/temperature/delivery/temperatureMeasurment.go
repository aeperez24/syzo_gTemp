package http

import (
	"net/http"
	usecase "temperature/useCase"
	"time"

	"github.com/spf13/viper"

	"github.com/labstack/echo"
)

//handler definition
type HttpMessageHandler struct {
	UseCase usecase.TemperatureUseCase
}

// func (a *HttpMessageHandler) getTemperatures(c echo.Context) error {

// 	start := c.QueryParam("start")
// 	end := c.QueryParam("end")

// 	log.Printf("los datos de quiery son %v y %v", start, end)
// 	result := a.UseCase.GetTemperatureRecords(start, end)

// 	return c.JSON(http.StatusOK, result)
// }
func (a *HttpMessageHandler) getTemperaturesByDay(c echo.Context) error {

	day := c.QueryParam("day")
	dayFormat := viper.GetString("dayformat")
	dayTime, error := time.Parse(dayFormat, day)
	if error != nil {
		panic("error al formatear fecha de consulta")
	}
	dayTime = time.Date(dayTime.Year(), dayTime.Month(), dayTime.Day(), 0, 0, 0, 0, time.Now().Location())
	result := a.UseCase.GetTemperatureByDay(dayTime)

	return c.JSON(http.StatusOK, result)
}
func (a *HttpMessageHandler) startReading(c echo.Context) error {

	a.UseCase.StartReading()

	return c.JSON(http.StatusOK, true)
}

func (a *HttpMessageHandler) stopReadiing(c echo.Context) error {

	a.UseCase.StopReading()

	return c.JSON(http.StatusOK, true)
}

//routing
func NewMesajeHttpHandler(e *echo.Echo, us usecase.TemperatureUseCase) {
	handler := &HttpMessageHandler{
		UseCase: us,
	}

	public := e.Group("/temperature")
	// public.GET("/getValues", handler.getTemperatures)
	public.GET("/start", handler.startReading)
	public.GET("/end", handler.stopReadiing)
	public.GET("/getValues", handler.getTemperaturesByDay)

}

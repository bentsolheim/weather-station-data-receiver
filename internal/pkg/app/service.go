package app

import "fmt"

type TemperatureData struct {
	TempCelsius float32
}

type DataLogService struct {
}

func (s *DataLogService) RegisterData(loggerId string, sensorName string, value float32) error {

	fmt.Printf("%s - %s - %f\n", loggerId, sensorName, value)
	return nil
}

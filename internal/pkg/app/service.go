package app

import "fmt"

type TemperatureData struct {
	TempCelsius float32
}

type Debug struct {
	SignalStrength   string
	TimeSpent        int32
	Iteration        int32
	Errors           int32
	MillisSinceStart int64
}

type DataLogService struct {
}

func (s *DataLogService) RegisterData(loggerId string, sensorName string, value float32) error {

	fmt.Printf("%s - %s - %f\n", loggerId, sensorName, value)
	return nil
}

func (s *DataLogService) RegisterDebug(loggerId string, debug Debug) error {

	fmt.Printf("%s - %+v\n", loggerId, debug)
	return nil
}

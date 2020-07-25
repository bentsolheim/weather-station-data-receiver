package app

import (
	"log"
)

type SensorReading struct {
	SensorName string
	Value      float32
	LocalTime  int32
}

type Debug struct {
	SignalStrength   string
	TimeSpent        int32
	Iteration        int32
	Errors           int32
	MillisSinceStart int64
}

func NewDataLogService() *DataLogService {
	return &DataLogService{entries: make(map[string]map[string]SensorReading)}
}

type DataLogService struct {
	entries      map[string]map[string]SensorReading
	debugEntries map[string]Debug
}

func (s *DataLogService) RegisterData(loggerId string, reading SensorReading) error {

	loggerEntries := s.entries[loggerId]
	if loggerEntries == nil {
		loggerEntries = make(map[string]SensorReading)
		s.entries[loggerId] = loggerEntries
	}
	loggerEntries[reading.SensorName] = reading
	log.Printf("%s - %+v\n", loggerId, reading)
	return nil
}

func (s *DataLogService) RegisterDebug(loggerId string, debug Debug) error {

	s.debugEntries[loggerId] = debug
	log.Printf("%s - %+v\n", loggerId, debug)
	return nil
}

func (s *DataLogService) FindLatestReadings(loggerId string) ([]SensorReading, error) {

	entries := s.entries[loggerId]
	values := make([]SensorReading, 0, len(entries))
	if entries != nil {
		for _, v := range entries {
			values = append(values, v)
		}
	}
	return values, nil
}

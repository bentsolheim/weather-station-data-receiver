package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SensorReading struct {
	SensorName string
	Value      float32
	LocalTime  int32
	UnixTime   int32
}

type Debug struct {
	SignalStrength   string
	TimeSpent        int32
	Iteration        int32
	ConnectionErrors int32
	SensorErrors     int32
	MillisSinceStart int64
	Battery          BatteryLevel
}

type BatteryLevel struct {
	AnalogReading int32
	Voltage       float32
	Level         int32
}

func NewDataLogService() *DataLogService {
	return &DataLogService{
		entries:      make(map[string]map[string]SensorReading),
		debugEntries: make(map[string]Debug),
	}
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

func (s *DataLogService) FindLatestDebug(loggerId string) (*Debug, error) {
	if debug, found := s.debugEntries[loggerId]; found {
		return &debug, nil
	}
	return nil, errors.New(fmt.Sprintf("No debug info found for logger %s", loggerId))
}

type LocationForeCast20Response struct {
	Type string
}

type MetService struct {
	BaseUrl string
}

func (s MetService) LoadForecast(lat string, lon string) {
	// https://api.met.no/weatherapi/locationforecast/2.0/complete?lat=58.55288&lon=8.97572
	res, err := http.Get(fmt.Sprintf("%s/weatherapi/locationforecast/2.0/complete?lat=%s&lon=%s", s.BaseUrl, lat, lon))
	if err != nil {
		log.Fatal(err)
	}
	bb, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	response := LocationForeCast20Response{}
	json.Unmarshal(bb, &response)

	fmt.Printf("%+v\n", response)
	println(string(bb))
}

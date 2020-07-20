package app

type TemperatureData struct {
	TempCelsius float32
}

type DataPayloadService struct {
}

func (s *DataPayloadService) GetMostRecentTemp() (TemperatureData, error) {

	return TemperatureData{TempCelsius: 23.4}, nil
}

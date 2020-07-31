package app_test

import (
	"github.com/bentsolheim/weather-station-data-receiver/internal/pkg/app"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAbs(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println(r.RequestURI)
		file, err := ioutil.ReadFile("testdata/locationforecast_2_0_complete_lat_58.55288_lon_8.97572.json")
		if err != nil {
			panic(err)
		}
		if _, err := w.Write(file); err != nil {
			panic(err)
		}
	}))
	defer ts.Close()

	metService := app.MetService{ts.URL}
	metService.LoadForecast("58.55288", "8.97572")
}

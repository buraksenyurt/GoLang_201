package weather

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetWeather(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockWeatherService := NewMockWeatherService(mockCtrl)
	mockWeatherService.EXPECT().GetWeather("Istanbul").Return("Sunny", nil)

	w, err := mockWeatherService.GetWeather("Istanbul")
	if err != nil {
		t.Fatal("Expected no error, got", err)
	}

	if w != "Sunny" {
		t.Fatalf("Expected Sunny, got %s", w)
	}
}

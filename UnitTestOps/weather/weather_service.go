package weather

type WeatherService interface {
	GetWeatherStatus(city string) (string, error)
}

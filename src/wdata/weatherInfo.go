package wdata

// WeatherDetail stores temperature, status and raining rate
type WeatherDetail struct {
	Temperature string
	Status      string
	RainingRate string
}

// WeatherDetailArr represents the next 7 days temperature
type WeatherDetailArr []*WeatherDetail

// WeeklyWeatherInfo Store weather data by week
type WeeklyWeatherInfo struct {
	City          string
	DayWeathers   WeatherDetailArr
	NightWeathers WeatherDetailArr
}

// SetCity set the city name.
func (info *WeeklyWeatherInfo) SetCity(c string) {
	info.City = c
}

// SetData set the weather information from queried data.
func (info *WeeklyWeatherInfo) SetData(d WeatherDetailArr, n WeatherDetailArr) {
	for index, val := range d {
		info.DayWeathers[index] = val
	}
	for index, val := range n {
		info.NightWeathers[index] = val
	}
}

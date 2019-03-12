package wdata

// WeatherDetail stores temperature, status and raining rate
type WeatherDetail struct {
	Temperature string
	Status      string
	RainingRate string
}

// WeatherDetailArr represents the next 7 days temperature
type WeatherDetailArr []*WeatherDetail

// WeatherInfo Store weather data by week
type WeatherInfo struct {
	City          string
	DayWeathers   WeatherDetailArr
	NightWeathers WeatherDetailArr
}

// IsGood return true if all data is filled
func (detail *WeatherDetail) IsGood() bool {
	return len([]rune(detail.Temperature)) > 0 && len([]rune(detail.Status)) > 0 && len([]rune(detail.RainingRate)) > 0
}

// SetCity set the city name.
func (info *WeatherInfo) SetCity(c string) {
	info.City = c
}

// SetData set the weather information from queried data.
func (info *WeatherInfo) SetData(d WeatherDetailArr, n WeatherDetailArr) {
	for index, val := range d {
		info.DayWeathers[index] = val
	}
	for index, val := range n {
		info.NightWeathers[index] = val
	}
}

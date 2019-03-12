package wdata

// Temperatures represents the next 7 days temperature
type Temperatures []string

// WeeklyWeatherInfo Store weather data by week
type WeeklyWeatherInfo struct {
	City      string
	DayData   Temperatures
	NightData Temperatures
}

// SetCity set the city name.
func (info *WeeklyWeatherInfo) SetCity(c string) {
	info.City = c
}

// SetData set the weather information from queried data.
func (info *WeeklyWeatherInfo) SetData(d Temperatures, n Temperatures) {
	for index, val := range d {
		info.DayData[index] = val
	}
	for index, val := range n {
		info.NightData[index] = val
	}
}

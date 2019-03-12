package wdata

import (
	"fmt"
)

// Dates represents the array of next 7 dates
type Dates []string

// WeatherInfoCollection store a dictionary for collecting all data from website.
type WeatherInfoCollection struct {
	Date     Dates
	Weathers map[string]*WeeklyWeatherInfo
}

// HasDate return true if dates already inserted.
func (collect *WeatherInfoCollection) HasDate() bool {
	return len(collect.Date) > 0
}

// SetDate set the date array from queried data.
func (collect *WeatherInfoCollection) SetDate(d Dates) {
	collect.Date = make([]string, len(d))
	copy(collect.Date, d)
}

// Query get weather data for given city name. Return empty data if given city is not found.
func (collect *WeatherInfoCollection) Query(city string) WeeklyWeatherInfo {
	info, exist := collect.Weathers[city]
	if exist {
		return *info
	}
	return WeeklyWeatherInfo{}
}

// PrettyPrint show data in the console with pretty print.
func (collect *WeatherInfoCollection) PrettyPrint() {
	fmt.Print("City\t")
	for _, date := range collect.Date {
		fmt.Printf("%s\t", date)
	}
	fmt.Println("")
	for key, info := range collect.Weathers {
		fmt.Printf("%s\t", key)
		for _, temp := range info.DayWeathers {
			fmt.Printf("%s:%s\t", temp.Temperature, temp.Status)
		}
		fmt.Println("")
	}
}

package wdata

import (
	"common"
	"fmt"
)

const (
	day   string = "白天"
	night string = "晚上"
)

// Dates represents the array of next 7 dates
type Dates []string

// WeatherInfoCollection store a dictionary for collecting all data from website.
type WeatherInfoCollection struct {
	Date     Dates
	Weathers map[string]*WeatherInfo
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
func (collect *WeatherInfoCollection) Query(city string) WeatherInfo {
	info, exist := collect.Weathers[city]
	if exist {
		return *info
	}
	return WeatherInfo{}
}

// Print print the given city weather data to console.
func (collect *WeatherInfoCollection) Print(city string, days int32, verbose bool) {
	if info, exist := collect.Weathers[city]; exist {
		days = common.Min(common.Max(1, days), 7)
		text := fmt.Sprintf("============ %s ============\n\n", city)

		for index := int32(0); index < days; index++ {
			date := collect.Date[index]
			dayDetail, nightDetail := info.DayWeathers[index], info.NightWeathers[index]
			dayTemp, nightTemp := fmt.Sprintf("%s:%s°C", day, dayDetail.Temperature), fmt.Sprintf("%s:%s°C", night, nightDetail.Temperature)
			if verbose {
				dayTemp += ":" + dayDetail.Status
				nightTemp += ":" + nightDetail.Status
			}
			text += fmt.Sprintf("%-15s%-30s\n", date, dayTemp)
			text += fmt.Sprintf("%-15s%-30s\n\n", date, nightTemp)
		}
		text += fmt.Sprintf("============ %s ============\n", city)
		fmt.Print(text)
	} else {
		if days <= 0 {
			fmt.Println("Please input correct days")
		} else {
			fmt.Println("Cannot find any data.")
		}
	}
}

// PrintAll show data in the console with pretty print.
func (collect *WeatherInfoCollection) PrintAll(days int32, verbose bool) {
	for city := range collect.Weathers {
		collect.Print(city, days, verbose)
	}
}

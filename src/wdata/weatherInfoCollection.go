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
		fmt.Printf("============ %s ============\n", city)
		fmt.Print("\t")
		for index := int32(0); index < days; index++ {
			fmt.Printf("%s\t", collect.Date[index])
		}
		fmt.Println()
		fmt.Printf("%s\t", day)
		for index := int32(0); index < days; index++ {
			detail := info.DayWeathers[index]
			fmt.Printf("%s", detail.Temperature)
			if verbose {
				fmt.Printf(":%s", detail.Status)
			}
			fmt.Print("\t")
		}
		fmt.Println()
		fmt.Printf("%s\t", night)
		for index := int32(0); index < days; index++ {
			detail := info.NightWeathers[index]
			fmt.Printf("%s", detail.Temperature)
			if verbose {
				fmt.Printf(":%s", detail.Status)
			}
			fmt.Print("\t")
		}
		fmt.Println()
	} else {
		if days <= 0 {
			fmt.Println("Please input correct days")
		} else {
			fmt.Println("Cannot find any data.")
		}
	}
}

// PrintAll show data in the console with pretty print.
func (collect *WeatherInfoCollection) PrintAll(verbose bool) {
	fmt.Print("\t")
	for _, date := range collect.Date {
		fmt.Printf("%s\t", date)
	}
	fmt.Println()
	fmt.Printf("%s\t", day)
	for key, info := range collect.Weathers {
		fmt.Printf("%s\n", key)
		for _, detail := range info.DayWeathers {
			fmt.Printf("%s", detail.Temperature)
			if verbose {
				fmt.Printf(":%s", detail.Status)
			}
			fmt.Print("\t")
		}
		fmt.Println()
		fmt.Printf("%s\t", night)
		for _, detail := range info.NightWeathers {
			fmt.Printf("%s", detail.Temperature)
			if verbose {
				fmt.Printf(":%s", detail.Status)
			}
			fmt.Print("\t")
		}
		fmt.Println()
	}
}

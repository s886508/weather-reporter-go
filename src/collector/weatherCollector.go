package collector

import (
	"fmt"
	"net/http"
)

// WeeklyWeatherURL indicate the website to retrieve data
const WeeklyWeatherURL string = "https://www.cwb.gov.tw/V7/forecast/week/week.htm"

// GetWeekData : Get next 7 days weather report from website.
func GetWeekData(wait chan int) {
	response, err := http.Get(WeeklyWeatherURL)
	if err != nil {
		fmt.Println("Connection error.")
		return
	}

	// Close connection afterwards
	defer response.Body.Close()

	// Parse HTML content
	weatherCollect := parseWeeklyHTML(response.Body)
	weatherCollect.PrettyPrint()

	// Use channel to tell the function has completed.
	wait <- 1
}

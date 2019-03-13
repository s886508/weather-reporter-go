package collector

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
	"wdata"
)

// WeeklyWeatherURL indicate the website to retrieve data
const WeeklyWeatherURL string = "https://www.cwb.gov.tw/V7/forecast/week/week.htm"

var wg sync.WaitGroup
var stop = make(chan bool)
var mutex sync.Mutex
var collection *wdata.WeatherInfoCollection

// retrieveOneWeekWeather : Get next 7 days weather report from website.
func retrieveOneWeekWeather() *wdata.WeatherInfoCollection {
	//fmt.Println("Retrieving weather data...")
	response, err := http.Get(WeeklyWeatherURL)
	if err != nil {
		fmt.Println("Connection error.")
		return nil
	}

	// Close connection afterwards
	defer response.Body.Close()

	// Parse HTML content
	weatherCollect := parseWeeklyHTML(response.Body)
	//weatherCollect.PrettyPrint()
	//fmt.Println("Retrieve weather data complete.")
	return weatherCollect
}

func timeTrace(start time.Time) time.Duration {
	return time.Since(start)
}

func retrieveRoutine() {
	startTime, elapse := time.Now(), time.Hour+1
	timeElapsed := func(s time.Time) bool {
		if elapse > time.Hour {
			startTime = s
			elapse = 0
			return true
		}

		elapse = time.Since(startTime)
		return false
	}

	for {
		select {
		case <-stop:
			defer wg.Done()
			return
		default:
			if timeElapsed(time.Now()) {
				mutex.Lock()
				collection = retrieveOneWeekWeather()
				mutex.Unlock()
			}
		}
	}
}

// Start run a worker thread to get weather data asynchronously
func Start() {
	wg.Add(1)
	go retrieveRoutine()
}

// Stop terminates the work thread
func Stop() {
	stop <- true
	wg.Wait()
}

// PrintWeatherData prints the weather information retieved from website
func PrintWeatherData(city string, days int32, verbose bool) {
	mutex.Lock()
	if collection != nil {
		switch strings.Compare(city, "all") {
		case 0:
			collection.PrintAll(verbose)
		default:
			collection.Print(city, days, verbose)
		}
	}
	mutex.Unlock()
}

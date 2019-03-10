package main

import (
	"collector"
	"fmt"
)

func main() {
	wait := make(chan int)
	go collector.GetWeekData(wait)
	if <-wait == 1 {
		fmt.Println("Retrieve weather data complete.")
	}
}

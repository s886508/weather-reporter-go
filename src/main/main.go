package main

import (
	"bufio"
	"collector"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	collector.Start()

	for {
		city, days := "all", 7

		fmt.Print("Please enter city name(Q to exit): ")
		arg, _ := reader.ReadString('\n')
		arg = strings.Replace(arg, "\n", "", -1)
		if strings.Compare(arg, "Q") == 0 {
			break
		}

		if len([]rune(arg)) > 0 {
			city = arg
		}

		fmt.Print("Please enter days to show(1~7): ")
		arg, _ = reader.ReadString('\n')
		arg = strings.Replace(arg, "\n", "", -1)

		if len([]rune(arg)) > 0 {
			d, _ := strconv.Atoi(arg)
			days = int(d)
		}

		collector.PrintWeatherData(city, int32(days), false)
	}

	collector.Stop()
}

package collector

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const weekDataURL string = "https://www.cwb.gov.tw/V7/forecast/week/week.htm"

const (
	htmlTagTableBody = "tbody"
	htmlTagTable     = "table"
	htmlTagTr        = "tr"
	htmlTagTh        = "th"
	htmlTagTd        = "td"
	htmlTagImg       = "img"
)

func parseDate(tokenizer *html.Tokenizer) []string {
	var dateTable []string
	var date string
	for {
		tokenType, token := tokenizer.Next(), tokenizer.Token()
		if tokenType == html.EndTagToken && token.Data == htmlTagTr {
			break
		}

		if tokenType == html.EndTagToken && token.Data == htmlTagTh {
			if len(date) > 0 {
				dateTable = append(dateTable, strings.TrimPrefix(date, ":"))
				date = ""
			}
			continue
		}

		if tokenType == html.TextToken {
			text := strings.Trim(token.Data, " \n\t")
			if !strings.Contains(text, "白天") && !strings.Contains(text, "晚上") {
				date = strings.Join([]string{date, text}, ":")
			}
		}
	}

	return dateTable
}

func parseCityData(tokenizer *html.Tokenizer) string {
	tokenType, token := tokenizer.Next(), tokenizer.Token()
	var city string
	if tokenType == html.TextToken {
		city = strings.Trim(token.Data, " \n\t")
	}
	return city
}

func parseWeatherData(tokenizer *html.Tokenizer) ([]string, []string) {
	addToData := func(text string) bool {
		return strings.Compare(text, "白天") != 0 && strings.Compare(text, "晚上") != 0 && len([]rune(text)) > 0
	}

	parseDegree := func() []string {
		var degree []string
		var text string
		for {
			tokenType, token := tokenizer.Next(), tokenizer.Token()
			// Break when parsing to </tr>
			if tokenType == html.EndTagToken && token.Data == htmlTagTr {
				break
			}

			if tokenType == html.EndTagToken && token.Data == htmlTagTd {
				if len(text) > 0 {
					degree = append(degree, strings.TrimPrefix(text, ":"))
					text = ""
				}
				continue
			}

			if tokenType == html.SelfClosingTagToken && token.Data == htmlTagImg {
				text = strings.Join([]string{text, strings.Trim(token.Attr[2].Val, " \n\t")}, ":")
				continue
			}

			if tokenType == html.TextToken {
				data := strings.Trim(token.Data, " \n\t")
				if addToData(data) {
					text = strings.Join([]string{text, strings.Trim(data, " \n\t")}, ":")
				}
			}
		}
		return degree
	}

	var dayWeather, nightWeather []string
	for {
		if len(dayWeather) == 7 && len(nightWeather) == 7 {
			break
		}

		tokenType, token := tokenizer.Next(), tokenizer.Token()

		if tokenType == html.StartTagToken && token.Data == htmlTagTd {
			dayWeather = parseDegree()
		}

		if tokenType == html.StartTagToken && token.Data == htmlTagTd {
			nightWeather = parseDegree()
		}
	}

	return dayWeather, nightWeather
}

func parseWeeklyData(tokenizer *html.Tokenizer) {
	traversToTHTag := func() html.Token {
		var token html.Token
		tokenType := html.ErrorToken
		for token.Data != htmlTagTh && tokenType != html.StartTagToken {
			tokenType, token = tokenizer.Next(), tokenizer.Token()
		}
		return token
	}

	var weekDates []string
	for {
		tokenType, token := tokenizer.Next(), tokenizer.Token()
		if tokenType == html.EndTagToken && token.Data == htmlTagTableBody {
			break
		}

		if tokenType == html.StartTagToken && token.Data == htmlTagTr {
			// Check first child <th> tag for checking what the row stands for.
			token = traversToTHTag()

			if len(token.Attr) == 1 {
				if len(weekDates) == 0 {
					weekDates = parseDate(tokenizer)
				}
			} else {
				city := parseCityData(tokenizer)
				dayWeatherData, nightWeatherData := parseWeatherData(tokenizer)
				fmt.Println(city)
				fmt.Println(dayWeatherData)
				fmt.Println(nightWeatherData)
			}
		}
	}
}

func parseHTML(r io.Reader) {
	tokenizer := html.NewTokenizer(r)
	for {
		tokenType, token := tokenizer.Next(), tokenizer.Token()
		if tokenType == html.StartTagToken && token.Data == htmlTagTable {
			parseWeeklyData(tokenizer)
			break
		}

		if tokenType == html.ErrorToken {
			// Handle end of file
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
		}
	}
}

// GetWeekData : Get next 7 days weather report from website.
func GetWeekData(wait chan int) {
	response, err := http.Get(weekDataURL)
	if err != nil {
		fmt.Println("Connection error.")
		return
	}

	// Close connection afterwards
	defer response.Body.Close()

	// Parse HTML content
	parseHTML(response.Body)

	// Use channel to tell the function has completed.
	wait <- 1
}

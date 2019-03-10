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
	htmlTagBody  = "body"
	htmlTagTable = "table"
	htmlTagTr    = "tr"
	htmlTagTh    = "th"
)

func parseDate(tokenizer *html.Tokenizer) []string {
	var dateTable []string
	var date string
	for {
		tokenType := tokenizer.Next()
		token := tokenizer.Token()
		if tokenType == html.EndTagToken && token.Data == htmlTagTr {
			break
		}

		if tokenType == html.EndTagToken && token.Data == htmlTagTh {
			if len(date) > 0 {
				dateTable = append(dateTable, date)
				date = ""
			}
			continue
		}

		if tokenType == html.TextToken {
			text := strings.Trim(token.Data, "\n")
			if !strings.Contains(text, "白天") && !strings.Contains(text, "晚上") {
				date += text
			}
		}
	}

	return dateTable
}

func parseWeeklyData(tokenizer *html.Tokenizer) {
	traversToTHTag := func() html.Token {
		var token html.Token
		tokenType := html.ErrorToken
		for token.Data != htmlTagTh && tokenType != html.StartTagToken {
			tokenType = tokenizer.Next()
			token = tokenizer.Token()
		}
		return token
	}

	var weekDates []string
	for {
		tokenType := tokenizer.Next()
		token := tokenizer.Token()
		if tokenType == html.EndTagToken && token.Data == htmlTagBody {
			break
		}

		if tokenType == html.StartTagToken && token.Data == htmlTagTr {
			token = traversToTHTag()

			if len(token.Attr) == 1 {
				if len(weekDates) == 0 {
					weekDates = parseDate(tokenizer)
				}
			} else {
				// ToDo: Parse weather data
			}
		}
	}

	fmt.Println(weekDates)
}

func parseHTML(r io.Reader) {
	tokenizer := html.NewTokenizer(r)
	for {
		tokenType := tokenizer.Next()
		token := tokenizer.Token()
		if tokenType == html.StartTagToken && token.Data == htmlTagTable {
			parseWeeklyData(tokenizer)
			continue
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

package collector

import (
	"common"
	"io"
	"strings"
	"wdata"

	"golang.org/x/net/html"
)

func textWanted(text string) bool {
	return !strings.Contains(text, "白天") && !strings.Contains(text, "晚上") && len([]rune(text)) > 0
}

func parseDate(tokenizer *html.Tokenizer) wdata.Dates {
	var dateArr wdata.Dates
	var date string
	for {
		tokenType, token := tokenizer.Next(), tokenizer.Token()
		if tokenType == html.EndTagToken && token.Data == common.HtmlTagTr {
			break
		}

		if tokenType == html.EndTagToken && token.Data == common.HtmlTagTh {
			if len(date) > 0 {
				dateArr = append(dateArr, strings.TrimPrefix(date, ":"))
				date = ""
			}
			continue
		}

		if tokenType == html.TextToken {
			text := strings.Trim(token.Data, " \n\t")
			if textWanted(text) {
				date = strings.Join([]string{date, text}, ":")
			}
		}
	}

	return dateArr
}

func parseCity(tokenizer *html.Tokenizer) string {
	tokenType, token := tokenizer.Next(), tokenizer.Token()
	var city string
	if tokenType == html.TextToken {
		city = strings.Trim(token.Data, " \n\t")
	}
	return city
}

func parseWeatherData(tokenizer *html.Tokenizer) (wdata.WeatherDetailArr, wdata.WeatherDetailArr) {

	parseTemperature := func() wdata.WeatherDetailArr {
		var detailArr wdata.WeatherDetailArr
		detail := new(wdata.WeatherDetail)
		for {
			tokenType, token := tokenizer.Next(), tokenizer.Token()
			// Break when parsing to </tr>
			if tokenType == html.EndTagToken && token.Data == common.HtmlTagTr {
				break
			}

			if tokenType == html.EndTagToken && token.Data == common.HtmlTagTd {
				if len(detail.Temperature) > 0 && len(detail.Status) > 0 {
					detailArr = append(detailArr, detail)
					detail = new(wdata.WeatherDetail)
				}
				continue
			}

			if tokenType == html.SelfClosingTagToken && token.Data == common.HtmlTagImg {
				detail.Status = strings.Trim(token.Attr[2].Val, " \n\t")
				continue
			}

			if tokenType == html.TextToken {
				data := strings.Trim(token.Data, " \n\t")
				if textWanted(data) {
					detail.Temperature = strings.Trim(data, " \n\t")
				}
			}
		}
		return detailArr
	}

	var dayWeather, nightWeather wdata.WeatherDetailArr
	for {
		if len(dayWeather) == 7 && len(nightWeather) == 7 {
			break
		}

		tokenType, token := tokenizer.Next(), tokenizer.Token()

		if tokenType == html.StartTagToken && token.Data == common.HtmlTagTd {
			dayWeather = parseTemperature()
		}

		if tokenType == html.StartTagToken && token.Data == common.HtmlTagTd {
			nightWeather = parseTemperature()
		}
	}

	return dayWeather, nightWeather
}

func parseWeeklyData(tokenizer *html.Tokenizer) *wdata.WeatherInfoCollection {
	traversToTHTag := func() html.Token {
		var token html.Token
		tokenType := html.ErrorToken
		for token.Data != common.HtmlTagTh && tokenType != html.StartTagToken {
			tokenType, token = tokenizer.Next(), tokenizer.Token()
		}
		return token
	}

	var collection = new(wdata.WeatherInfoCollection)
	collection.Weathers = map[string]*wdata.WeatherInfo{}
	for {
		tokenType, token := tokenizer.Next(), tokenizer.Token()
		if tokenType == html.EndTagToken && token.Data == common.HtmlTagTableBody {
			break
		}

		if tokenType == html.StartTagToken && token.Data == common.HtmlTagTr {
			// Check first child <th> tag for checking what the row stands for.
			token = traversToTHTag()

			if len(token.Attr) == 1 {
				if collection.HasDate() == false {
					collection.SetDate(parseDate(tokenizer))
				}
			} else {
				var info = new(wdata.WeatherInfo)
				info.City = parseCity(tokenizer)
				dayWeatherData, nightWeatherData := parseWeatherData(tokenizer)
				info.DayWeathers = dayWeatherData
				info.NightWeathers = nightWeatherData
				collection.Weathers[info.City] = info
			}
		}
	}
	return collection
}

func parseWeeklyHTML(r io.Reader) *wdata.WeatherInfoCollection {
	tokenizer := html.NewTokenizer(r)
	for {
		tokenType, token := tokenizer.Next(), tokenizer.Token()
		if tokenType == html.StartTagToken && token.Data == common.HtmlTagTable {
			return parseWeeklyData(tokenizer)
		}

		if tokenType == html.ErrorToken {
			// Handle end of file
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
		}
	}
	return nil
}

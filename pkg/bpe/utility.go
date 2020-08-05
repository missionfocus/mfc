package bpe

import (
	"strings"
	"time"
)

const ( //TODO move this
	glTimeFormat = "2006-01-02"
)

//GetTimeParameters is used to alter the format [date] | [date] into a comparable format
func GetTimeParameters(str string) []time.Time { //TODO move this
	dates := make([]time.Time, 0)

	if len(str) == 0 {
		date := "1999-12-31"
		t, _ := time.Parse(glTimeFormat, date)
		dates = append(dates, t)

		currentTime := time.Now()
		currentTime.Format(glTimeFormat)
		dates = append(dates, currentTime)
	}

	splitDateStrings := strings.Split(str, "|")
	for _, d := range splitDateStrings {
		strToDate := strings.Replace(d, " ", "", -1)
		t, _ := time.Parse(glTimeFormat, strToDate)
		dates = append(dates, t)
	}
	return dates
}


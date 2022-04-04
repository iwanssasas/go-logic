package datechallenge

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type (
	Dest struct {
		Hasil []Week
	}

	Hasil []Week

	Week struct {
		NameWeek  string `json:"week"`
		DayInWeek int    `json:"day_in_week"`
		Date      []Date `json:"date"`
	}

	Date struct {
		Date string
	}

	Data struct {
		Tanggal string
		Week    int
	}

	Database []Data
)

func rangeDate(start, end time.Time) func() time.Time {
	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}

func Parsing(inputdate string) time.Time {
	layoutFormat := "2006-01-02"
	start, err := time.Parse(layoutFormat, inputdate)
	if err != nil {
		fmt.Println(err)
	}
	return start
}

func LastMonth(start time.Time) time.Time {
	currentYear, currentMonth, _ := start.Date()
	currentLocation := start.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	end := firstOfMonth.AddDate(0, 1, -1)

	return end
}

func GetDate(start, end time.Time) []time.Time {
	result := []time.Time{}

	for rd := rangeDate(start, end); ; {
		date := rd()
		if date.IsZero() {
			break
		}
		result = append(result, date)
	}

	return result
}
func TestNumberToAlphabet(t *testing.T) {
	var dest Dest
	var final Hasil
	var database Database
	var array [][]Date
	first := 0

	inputdate := "2022-05-01"

	start := Parsing(inputdate)
	end := LastMonth(start)
	result := GetDate(start, end)

	for _, val := range result {
		_, week := val.ISOWeek()

		database = append(database, Data{
			Tanggal: val.Format("Monday, 2006-01-02"),
			Week:    week,
		})
	}

	tempWeek := make(map[string]int)
	for _, val := range database {
		keyWeek := fmt.Sprintf("Week-%v", val.Week)

		idxWeek, hasWeek := tempWeek[keyWeek]
		if hasWeek {
			dest.Hasil[idxWeek].Date = append(dest.Hasil[idxWeek].Date, Date{
				Date: val.Tanggal,
			})
		} else {
			dest.Hasil = append(dest.Hasil, Week{
				NameWeek: keyWeek,
				Date: []Date{
					Date{
						Date: val.Tanggal,
					},
				},
			})
			tempWeek[keyWeek] = len(dest.Hasil) - 1
		}
	}

	for _, val := range dest.Hasil {
		a := val.Date
		array = append(array, a)
	}

	if len(array[0]) <= 4 {
		array[1] = append(array[1], array[0]...)
		array = append(array[:0], array[0+1:]...)
	}

	lastPosition := len(array) - 1
	if len(array[lastPosition]) <= 4 {
		array[lastPosition-1] = append(array[lastPosition-1], array[lastPosition]...)
		array = append(array[:lastPosition], array[lastPosition+1:]...)
	}

	for _, val := range array {
		first += 1
		key := fmt.Sprintf("week-%v", first)
		final = append(final, Week{
			NameWeek:  key,
			DayInWeek: len(val),
			Date:      val,
		})

	}
	response, err := json.Marshal(final)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(response))

}

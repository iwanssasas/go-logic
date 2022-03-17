package mapping_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

var jsonString = `
[
    {
        "area_name": "AREA-1",
        "seam_name": "SEAM-1",
        "actual": 200,
		"date": "01-05-2022"
    },
    {
        "area_name": "AREA-1",
        "seam_name": "SEAM-2",
        "actual": 200,
		"date": "01-07-2022"
    },
    {
        "area_name": "AREA-2",
        "seam_name": "SEAM-3",
        "actual": 200,
		"date": "01-08-2022"
    },
    {
        "area_name": "AREA-2",
        "seam_name": "SEAM-4",
        "actual": 200,
		"date": "01-16-2022"
    },
    {
        "area_name": "AREA-2",
        "seam_name": "SEAM-5",
        "actual": 200,
		"date": "01-23-2022"
    },
    {
        "area_name": "AREA-2",
        "seam_name": "SEAM-5",
        "actual": 200,
		"date": "02-01-2022"
    },
    {
        "area_name": "AREA-1",
        "seam_name": "SEAM-1",
        "actual": 200,
		"date": "02-04-2022"
    },
    {
        "area_name": "AREA-1",
        "seam_name": "SEAM-2",
        "actual": 200,
		"date": "02-04-2022"
    },
    {
        "area_name": "AREA-1",
        "seam_name": "SEAM-2",
        "actual": 200,
		"date": "02-06-2022"
    },
    {
        "area_name": "AREA-1",
        "seam_name": "SEAM-2",
        "actual": 200,
		"date": "02-18-2022"
    },
    {
        "area_name": "AREA-1",
        "seam_name": "SEAM-2",
        "actual": 200,
		"date": "02-19-2022"
    }
]
`

type (
	CustomTime struct {
		time.Time
	}

	DataInput struct {
		AreaName string     `json:"area_name"`
		SeamName string     `json:"seam_name"`
		Actual   int        `json:"actual"`
		Date     CustomTime `json:"date"`
	}
	DataInputs []DataInput

	Seam struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}
	Seams []Seam

	Area struct {
		Name  string `json:"name"`
		Total int    `json:"total"`
		Seam  Seams  `json:"seam"`
	}
	Areas []Area

	Week struct {
		Name  string `json:"name"`
		Week  int    `json:"week"`
		Total int    `json:"total"`
		Area  Areas  `json:"area"`
	}
	Weeks []Week

	Summary struct {
		Total       int `json:"total"`
		Week        int `json:"week"`
		AreaPerWeek int `json:"area_per_week"`
		SeamPerWeek int `json:"seam_per_week"`
	}

	Response struct {
		Summary Summary `json:"summary"`
		Week    Weeks   `json:"week"`
	}
)

const ctLayout = "01-02-2006"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

func TestMappingData(t *testing.T) {

	req := DataInputs{}
	err := json.Unmarshal([]byte(jsonString), &req)
	if err != nil {
		panic(err)
	}

	dest := Response{}

	tempWeek := make(map[string]int)
	tempArea := make(map[string]int)
	tempSeam := make(map[string]int)

	for _, val := range req {
		dest.Summary.Total += val.Actual

		_, week_ke := val.Date.ISOWeek()

		keyWeek := fmt.Sprintf("Week - %v", week_ke)
		keyArea := fmt.Sprintf("%v#%v", keyWeek, val.AreaName)
		keySeam := fmt.Sprintf("%v#%v#%v", week_ke, val.SeamName, val.AreaName)

		inputSeam := Seam{
			Name:  val.SeamName,
			Value: val.Actual,
		}

		idxWeek, hasWeek := tempWeek[keyWeek]
		if hasWeek {
			idxArea, hasArea := tempArea[keyArea]
			if hasArea {
				idxSeam, hasSeam := tempSeam[keySeam]
				if hasSeam {
					dest.Week[idxWeek].Area[idxArea].Seam[idxSeam].Value += val.Actual
					dest.Week[idxWeek].Area[idxArea].Total += val.Actual
					dest.Week[idxWeek].Total += val.Actual

				} else {
					dest.Week[idxWeek].Area[idxArea].Seam = append(dest.Week[idxWeek].Area[idxArea].Seam, inputSeam)
					tempSeam[keySeam] = len(dest.Week[idxWeek].Area[idxArea].Seam) - 1
					dest.Week[idxWeek].Area[idxArea].Total += val.Actual
					dest.Week[idxWeek].Total += val.Actual
					dest.Summary.SeamPerWeek++
				}

			} else {
				dest.Week[idxWeek].Area = append(dest.Week[idxWeek].Area, Area{
					Name:  val.AreaName,
					Total: val.Actual,
					Seam: Seams{
						inputSeam,
					},
				})
				tempArea[keyArea] = len(dest.Week[idxWeek].Area) - 1
				tempSeam[keySeam] = len(dest.Week[idxWeek].Area[idxArea].Seam) - 1
				dest.Week[idxWeek].Total += val.Actual
				dest.Summary.AreaPerWeek++
				dest.Summary.SeamPerWeek++

			}

		} else {
			dest.Week = append(dest.Week, Week{
				Name:  keyWeek,
				Week:  week_ke,
				Total: val.Actual,
				Area: Areas{
					Area{
						Name:  val.AreaName,
						Total: val.Actual,
						Seam: Seams{
							inputSeam,
						},
					},
				},
			})

			idxWeekArray := len(dest.Week) - 1
			tempWeek[keyWeek] = idxWeekArray

			idxAreaArray := len(dest.Week[idxWeekArray].Area) - 1
			tempArea[keyArea] = idxAreaArray

			tempSeam[keySeam] = len(dest.Week[idxWeekArray].Area[idxAreaArray].Seam) - 1

			dest.Summary.Week++
			dest.Summary.AreaPerWeek++
			dest.Summary.SeamPerWeek++

		}
	}

	result, err := json.Marshal(dest)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}

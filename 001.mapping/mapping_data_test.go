package mapping

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonString = `
[
    {
        "area_name": "AREA-1",
        "seam_name": "SEAM-1",
        "actual": 200
    },
    {
        "area_name": "AREA-1",
        "seam_name": "SEAM-2",
        "actual": 200
    },
    {
        "area_name": "AREA-2",
        "seam_name": "SEAM-3",
        "actual": 300
    },
    {
        "area_name": "AREA-2",
        "seam_name": "SEAM-4",
        "actual": 300
    },
    {
        "area_name": "AREA-2",
        "seam_name": "SEAM-5",
        "actual": 300
    },
    {
        "area_name": "AREA-2",
        "seam_name": "SEAM-5",
        "actual": 300
    }
]
`

type (
	DataInput struct {
		AreaName string `json:"area_name"`
		SeamName string `json:"seam_name"`
		Actual   int    `json:"actual"`
	}
	DataInputs []DataInput

	Seam struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	Seams []Seam

	Data struct {
		Name  string `json:"name"`
		Total int    `json:"total"`
		Seam  Seams  `json:"seam"`
	}

	Datas []Data

	Summary struct {
		Total int `json:"total"`
		Seam  int `json:"seam"`
	}
	Response struct {
		Summary Summary `json:"summary"`
		Data    Datas   `json:"data"`
	}
)

func TestMappingData(t *testing.T) {
	var req DataInputs

	err := json.Unmarshal([]byte(jsonString), &req)
	if err != nil {
		panic(err)
	}

	dest := Response{}

	tempArea := make(map[string]int)
	tempSeam := make(map[string]int)

	for _, val := range req {
		dest.Summary.Total += val.Actual

		keyArea := val.AreaName
		keySeam := fmt.Sprintf("%v#%v", val.AreaName, val.SeamName)

		idxArea, hasArea := tempArea[keyArea]
		if hasArea {
			idxSeam, hasSeam := tempSeam[keySeam]
			if hasSeam {
				dest.Data[idxArea].Seam[idxSeam].Value += val.Actual
			} else {
				dest.Data[idxArea].Seam = append(dest.Data[idxArea].Seam, Seam{ // masukkan data seam dalam area yg sudah ada
					Name:  val.SeamName,
					Value: val.Actual,
				})
				tempSeam[keySeam] = len(dest.Data[idxArea].Seam) - 1 // membuat posisi array seam yg baru aja dimasukkan
				dest.Summary.Seam++
			}
			dest.Data[idxArea].Total += val.Actual
		} else {
			dest.Data = append(dest.Data, Data{ // masukkan data area belum ada
				Name: val.AreaName,
				Seam: Seams{ // masukkan data seam dari area yg baru
					Seam{
						Name:  val.SeamName,
						Value: val.Actual,
					},
				},
			})

			idx := len(dest.Data) - 1
			tempArea[keyArea] = idx // 0, 1 membuat posisi array area yg baru aja dimasukkan (isi dalam data ada berapa area)

			tempSeam[keySeam] = len(dest.Data[idx].Seam) - 1 // membuat posisi array seam yg baru aja dimasukkan (isi dalam area ada berapa seam)
			dest.Summary.Seam++
			dest.Data[idx].Total += val.Actual
		}

	}

	result, err := json.Marshal(dest)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}

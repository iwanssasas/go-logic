package main

import (
	"fmt"
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

type DataInput struct {
	AreaName string `json:"area_name"`
	SeamName string `json:"seam_name"`
	Actual   int    `json:"actual"`
}

type DataInputs []DataInput

func main() {

	var chicken = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	var slice []int

	for _, val := range chicken {
		slice = append(slice, val)
	}

	fmt.Println(slice)
}

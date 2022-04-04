package utils

import "math"

var Alfabets = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func GetColumnFromInt(i int) string {
	// Get max sheet to 2 digits (ZZ)
	maxThreeSheet := len(Alfabets)*len(Alfabets) + len(Alfabets)
	if i == 0 || i > maxThreeSheet {
		return ""
	}

	if len(Alfabets) >= i {
		return Alfabets[i-1]
	}

	first := int(math.Floor(float64(i) / float64(len(Alfabets))))
	second := (i % len(Alfabets))
	if second == 0 {
		first = first - 1
		second = len(Alfabets)
	}

	return Alfabets[first-1] + Alfabets[second-1]
}

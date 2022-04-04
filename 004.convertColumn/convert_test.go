package convertcolumn

import (
	"fmt"
	"testing"
)

func numberToAlphabet(number int) string {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var letter string
	for {
		number -= 1
		letter = string(alphabet[number%26]) + letter
		number = (number / 26) >> 0
		if !(number > 0) {
			break
		}
	}

	return letter
}

func TestNumberToAlphabet(t *testing.T) {
	fmt.Println(numberToAlphabet(703))
}

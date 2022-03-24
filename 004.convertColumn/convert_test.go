package convertcolumn

import (
	"fmt"
	"testing"
)

func Coba(col int) string {
	alphabets := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if col <= 26 {
		return alphabets[col-1]
	} else if col >= 26 && col < 703 {
		if col%26 == 0 {
			return alphabets[col/26-2] + alphabets[25]
		} else {
			return alphabets[col/26-1] + alphabets[col%26-1]
		}

	} else {
		if col%26 == 0 {
			return alphabets[(col/26/26)-1] + alphabets[(col/26)%26-2] + alphabets[25]
		} else {
			return alphabets[(col/26/26)-1] + alphabets[(col/26)%26-1] + alphabets[(col%26)-1]
		}
	}
}

func TestIwan(t *testing.T) {
	fmt.Println(Coba(701))

	fmt.Println((18278/26/26)-2, (18278/26)%26+24, (18278%26)+25)

}

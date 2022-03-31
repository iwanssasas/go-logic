package datechallenge

import (
	"fmt"
	"testing"
)

func TestNumberToAlphabet(t *testing.T) {
	array := [][]string{{"1", "2", "3", "4"}, {"1", "2", "3", "4", "5", "6", "7"}, {"1", "2", "3", "4", "5", "6", "7"}, {"1", "2", "3", "4", "5", "6", "7"}, {"1", "2"}}

	if len(array[0]) <= 4 {
		array[1] = append(array[1], array[0]...)
		array = append(array[:0], array[0+1:]...)
	}

	lastPosition := len(array) - 1
	if len(array[lastPosition]) <= 4 {
		array[lastPosition-1] = append(array[lastPosition-1], array[lastPosition]...)
		array = append(array[:lastPosition], array[lastPosition+1:]...)
	}

	a := []int{1, 2, 3}
	b := []int{5, 6, 7}

	a = append(a, b...)
	fmt.Println(a)

	fmt.Println(array)
	fmt.Println(len(array))

}

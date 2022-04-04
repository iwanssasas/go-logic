package iwan

import (
	"fmt"
	"testing"
)

func TestNumberToAlphabet(t *testing.T) {
	// Example - 1
	str := "GOLANG"
	runes := []rune(str)

	var result []int

	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}

	fmt.Println(result)

	// Example - 2
	s := "GOLANG"
	for _, r := range s {
		fmt.Printf("%c - %d\n", r, r)
	}

}

package stringchalleng

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func GetAplhabetAndSplit(input string) []string {

	re := regexp.MustCompile("[^A-Z][^a-z]")
	strFill := re.ReplaceAllString(input, "")
	strLower := strings.ToLower(strFill)

	arr := strings.Split(strLower, "")

	return arr
}

func TestString(t *testing.T) {
	var hasil string
	input := "WWWwwwggoppwwwxxxxaaa"
	str := GetAplhabetAndSplit(input)
	fmt.Println(str)

	jumlah := 0

	for idx, val := range str {
		if idx < len(str)-1 {
			if val != str[idx+1] {
				jumlah++
				strHasil := fmt.Sprintf("%v%v", jumlah, val)
				hasil += strHasil

				jumlah = 0
			} else {
				jumlah++
			}
		} else {
			strHasil := fmt.Sprintf("%v%v", jumlah+1, val)
			hasil += strHasil
		}
	}

	fmt.Println(hasil)
}

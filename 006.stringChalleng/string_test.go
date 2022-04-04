package stringchalleng

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

type Data struct {
	Huruf  string
	Jumlah int
}

type ArrData []Data

func GetAplhabetAndSplit(input string) []string {

	re := regexp.MustCompile("[^A-Z][^a-z]")
	strFill := re.ReplaceAllString(input, "")
	strLower := strings.ToLower(strFill)

	arr := strings.Split(strLower, "")

	return arr
}

func TestString(t *testing.T) {
	var dest ArrData
	var hasil string
	input := "WWWwwwggoppwwwxxxx"
	str := GetAplhabetAndSplit(input)
	fmt.Println(str)

	jumlah := 0

	for idx, val := range str {
		if idx < len(str)-1 {
			if val != str[idx+1] {
				jumlah++
				dest = append(dest, Data{
					Huruf:  val,
					Jumlah: jumlah,
				})
				jumlah = 0
			} else {
				jumlah++
			}
		} else {
			dest = append(dest, Data{
				Huruf:  val,
				Jumlah: jumlah + 1,
			})
		}
	}

	for _, val := range dest {
		strHasil := fmt.Sprintf("%v%v", val.Jumlah, val.Huruf)
		hasil += strHasil
	}

	fmt.Println(hasil)
}

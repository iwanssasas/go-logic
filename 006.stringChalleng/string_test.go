package stringchalleng

import (
	"fmt"
	"regexp"
	"strconv"
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

	arr := strings.Split(strFill, "")

	return arr
}

func TestString(t *testing.T) {
	var dest ArrData
	var hasil []string
	input := "123WWWwwwggopp##%^&*><:{}+_)(*&"

	arr := GetAplhabetAndSplit(input)

	tempMap := make(map[string]int)
	for _, val := range arr {
		idx, ok := tempMap[val]
		if ok {
			dest[idx].Jumlah++
		} else {

			dest = append(dest, Data{
				Huruf:  val,
				Jumlah: 1,
			})

			tempMap[val] = len(dest) - 1
		}
	}

	for _, val := range dest {
		hasil = append(hasil, strconv.Itoa(val.Jumlah))
		hasil = append(hasil, val.Huruf)

	}

	fmt.Println(strings.Join(hasil, ""))
}

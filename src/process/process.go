package process

import (
	"strings"
)


func SentenceLengths(bd [][]string) []int {
	var l []int
	for i := 0; i < len(bd); i++ {
		l = append(l, len(bd[i]))
	}
	return l
}

func Average(s []int) int {
	var val int
	for i := 0; i < len(s); i++ {
		val = val + s[i]
	}
	return val / len(s)
}

func BreakDown(s string,) [][]string {
	var bd [][]string
	var sliceA, sliceB []string

	//after splitting by ".", sliceA has master slice
	sliceA = strings.SplitAfter(s, ".")

	//after splitting by "?", sliceB has master slice
	for i := 0; i < len(sliceA); i++ {
		j := strings.SplitAfter(sliceA[i], "?")

		for k := 0; k < len(j); k++ {
			sliceB = append(sliceB, j[k])
		}
	}

	sliceA = nil

	//after splitting by "!", sliceA has master slice
	for x := 0; x < len(sliceB); x++ {
		y := strings.SplitAfter(sliceB[x], "!")

		for z := 0; z < len(y); z++ {
			sliceA = append(sliceA, y[z])
		}
	}

	//split each sentence by space
	for a := 0; a < len(sliceA); a++ {
		b := strings.Split(sliceA[a], " ")

		//Remove excess punctuation
		for c := 0; c < len(b); c++ {
			// b[c] = strings.Replace(b[c], ",", "", -1)
			b[c] = strings.Replace(b[c], "\"", "", -1)
			b[c] = strings.Replace(b[c], "'", "", -1)
			// b[c] = strings.Replace(b[c], ";", "", -1)
			// b[c] = strings.Replace(b[c], ":", "", -1)
		}

		bd = append(bd, b)
	}

	return bd
}
package analytics

import (
	"fmt"
	"sort"
)

//NextWord returns an slice of slices where each word is followed by 
func NextWord(bd [][]string) [][]string {

	var nwA, nwB [][]string
	//also nwB

	//Looping thru sentences
	for i := 0; i < len(bd); i++ {

		//Looping thru words
		for j := 0; j < len(bd[i]); j++ {
			// fmt.Println(bd[i][j])

			//Looping thru next word slice
			if len(nwA) > 0 {

				for k := 0; k < len(nwA); k++ {

					index := sort.SearchStrings(nwB, bd[i][j])

					if nwA[k][0] != bd[i][j] {
						fmt.Println(nwA[k])
						return nwA[k] = append(nwA[k], bd[i][j + 1])
					}
				}
			} else {
				nwA[k][0] = append(nwA[0], bd[i][j + 1])
			}
		}
	}
	return nwA
}

//loop thru sentences
//loop thru words
//once we have word, loop thru nwA to find if word exists already.

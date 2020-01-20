package analytics

import (
)

func NextWord(bd [][]string) [][]string {

	var nwA, nwB [][]string

	//Looping thru sentences
	for i := 0; i < len(bd); i++ {

		//Looping thru words
		for j := 0; j < len(bd[i]); j++ {

			//Looping thru next word slice
			for k := 0; k < len(nw); k++ {

				if nw[k][0] == bd[i][j] {
					nw[k] = append(nw[k], bd[i][j + 1])
					break;
				}
			}


		}
		
	}




	return nwB
}
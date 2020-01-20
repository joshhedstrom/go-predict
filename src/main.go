package main

import (
	"fmt"

	// "analytics"
	"process"
)

func main() {

	//Break Down the document
	document := string("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ac tortor vitae purus faucibus ornare suspendisse sed nisi. Pellentesque habitant morbi tristique senectus et netus et malesuada. Suspendisse potenti nullam ac tortor. Viverra ipsum nunc aliquet bibendum enim facilisis gravida neque. Blandit volutpat maecenas volutpat blandit aliquam etiam erat. Diam ut venenatis tellus in metus vulputate eu scelerisque. Egestas egestas fringilla phasellus faucibus. Et ultrices neque ornare aenean euismod elementum nisi quis eleifend. Mus mauris vitae ultricies leo integer malesuada nunc.")
	
	bd := process.BreakDown(document)

	avgLength := process.Average(process.SentenceLengths(bd))

	//Analyze the document

	//Predict the document


	//Print results
	// fmt.Println("-------------")
	fmt.Print(bd)
	fmt.Println("-------------")

}

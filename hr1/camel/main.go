package main

import "fmt"

var input = "oneTwoThreeFourFive"

func main() {
	foundWords := 0

	for i := 0; i < len(input); i++ {
		charVariable := input[i]
		if charVariable < 96 {
			foundWords++
		}
		//fmt.Println(charVariable)
	}

	fmt.Printf("The input string had %d words.", foundWords+1)
}

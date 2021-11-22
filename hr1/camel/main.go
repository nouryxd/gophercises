package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	fmt.Println("Input is:", input)

	// Solution 1
	// answer := 1
	//for _, ch := range input {
	//	min, max := 'A', 'Z'
	//	if ch >= min && ch <= max {
	//		answer++
	//	}

	// Solution 2
	// str := string(ch)
	// if strings.ToUpper(str) == str {
	// 	answer++
	// }
	//}

	// Solution 3 (my own one)
	foundWords := 0
	for i := 0; i < len(input); i++ {
		charVariable := input[i]
		if charVariable < 96 {
			foundWords++
		}
	}
	fmt.Printf("The input string had %d words.", foundWords+1)
}

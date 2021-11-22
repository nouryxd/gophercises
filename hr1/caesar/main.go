package main

import "fmt"

func main() {
	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%d\n", &input)
	fmt.Scanf("%d\n", &delta)

	//	fmt.Printf("length: %d\n", length)
	//	fmt.Printf("input: %s\n", input)
	//	fmt.Printf("delta: %d\n", delta)

	alphabet := []rune("abcdefghijklmnopqrstuvwxyz")
	newRune := rotate('z', 2, alphabet)
	fmt.Println(string(newRune))
}

func rotate(s rune, delta int, key []rune) rune {
	idx := -1
	for i, r := range key {
		if r == s {
			idx = i
			break
		}
	}
	if idx < 0 {
		panic("idx less than 0")
	}
	for i := 0; i < delta; i++ {
		idx++
		if idx >= len(key) {
			idx = 0
		}
	}
	return key[idx]
}

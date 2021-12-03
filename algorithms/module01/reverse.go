package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	n := len(word)
	runes := make([]rune, n)
	for _, rune := range word {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

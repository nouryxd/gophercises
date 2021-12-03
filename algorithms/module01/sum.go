package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	ret := 0
	for _, v := range numbers {
		ret += v
	}
	return ret

}

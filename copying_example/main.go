package main

import "fmt"

type Demo struct {
	Name    string
	Numbers []int
}

func main() {
	var d Demo
	d.Name = "jon"
	d.Numbers = []int{1, 5, 8}

	d2 := d
	fmt.Println("original d2=", d2)
	d2.Name = "bob"
	d2.Numbers[0], d2.Numbers[1] = d2.Numbers[1], d2.Numbers[0]
	d2.Numbers = []int{4, 5, 6}
	fmt.Println("d =", d)
	fmt.Println("d2 =", d2)
}

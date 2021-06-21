package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type question struct {
	question string
	answer   string
}

func main() {
	in, err := os.ReadFile("problems.csv")
	if err != nil {
		log.Fatal("Couldn't read csv file")
	}
	str := string(in)

	r := csv.NewReader(strings.NewReader(str))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(records)
}

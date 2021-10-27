package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	"github.com/lyx0/gophercises/quiz-game/pkg/parser"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {

		exit(fmt.Sprintln("Failed to open the csv file"))
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Error while reading the csv file")
	}

	problems := parser.ParseLines(lines)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s\n", i+1, p.Question)

	}

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

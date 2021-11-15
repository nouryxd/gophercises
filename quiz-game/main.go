package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {

		fmt.Printf("Failed to open the csv file: %s", *csvFilename)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Error while reading the csv file")
	}

	problems := parseLines(lines)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s\n", i+1, p.Question)

	}

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			Question: line[0],
			Answer:   line[1],
		}
	}

	return ret
}

type problem struct {
	Question string
	Answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

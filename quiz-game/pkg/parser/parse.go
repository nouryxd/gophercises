package parser

func ParseLines(lines [][]string) []problem {
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

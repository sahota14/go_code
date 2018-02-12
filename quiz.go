package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	filename := flag.String("csv", "problems.csv", "a csv file in the format of question answer")
	flag.Parse()

	file, err := os.Open(*filename)

	if err != nil {
		exit(fmt.Sprintf("could not open file: %s\n", *filename))

	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("faiiled to parse CSV file")

	}

	problems := parseLines(lines)

	correct_answers := 0
	for i, p := range problems {
		fmt.Printf("problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct_answers++
		}
	}

	fmt.Printf("you scored %d out of %d\n", correct_answers, len(problems))

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines)) //return a slice of problems which is the length of the total number of lines
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {

	fmt.Println(msg)
	os.Exit(1)

}

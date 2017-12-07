package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	question string
	answer   string
}

func main() {
	filename := flag.String("csv", "quiz.csv", "a csv file in the format 'question,answer'")

	flag.Parse()

	lines, err := readFile(*filename)

	if err != nil {
		exit(err)
	}

	problems := parseQuestions(lines)

	correct := makeQuiz(problems)

	fmt.Printf("\nDone!\n %d correct answers over %d questions\n", len(problems), correct)
}

func readFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	r := csv.NewReader(file)

	return r.ReadAll()
}

func parseQuestions(lines [][]string) []problem {
	problems := make([]problem, 0, len(lines))

	for _, line := range lines {
		problems = append(problems, problem{
			question: line[0],
			answer:   line[1],
		})
	}

	return problems
}

func makeQuiz(problems []problem) int {
	var correct int

	for i, p := range problems {
		if makeQuestion(i+1, p) {
			correct++
		}
	}

	return correct
}

func makeQuestion(number int, p problem) bool {
	var answer string

	fmt.Printf("%d) %s: ", number, p.question)
	fmt.Scanf("%s\n", &answer)

	return p.answer == answer
}

func exit(err error) {
	fmt.Printf(fmt.Sprintf("Failed on %s", err))
	os.Exit(1)
}

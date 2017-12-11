package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

var input = bufio.NewScanner(os.Stdin)

func main() {
	filename := flag.String("csv", "quiz.csv", "a csv file in the format 'question,answer'")
	timeout := flag.Int("timer", 30, "time limit for the quiz in seconds")

	flag.Parse()

	lines, err := readFile(*filename)

	if err != nil {
		exit(err)
	}

	problems := parseQuestions(lines)

	fmt.Println("Press [enter] to start quiz...")
	input.Scan()

	correct := makeQuiz(problems, *timeout)

	fmt.Printf("\nDone!\n %d correct answers over %d questions\n", correct, len(problems))
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

func makeQuiz(problems []problem, timeout int) int {
	var correct int
	timer := time.NewTimer(time.Second * time.Duration(timeout))

quizloop:
	for i, p := range problems {
		c := make(chan bool)

		go func() {
			c <- makeQuestion(i+1, p)
		}()

		select {
		case <-timer.C:
			fmt.Println("\n\nYour time is over!")
			break quizloop

		case answer := <-c:
			if answer {
				correct++
			}
		}
	}

	return correct
}

func makeQuestion(number int, p problem) bool {
	var answer string

	fmt.Printf("%d) %s: ", number, p.question)
	input.Scan()

	answer = input.Text()

	return p.answer == answer
}

func exit(err error) {
	fmt.Printf(fmt.Sprintf("Failed on %s", err))
	os.Exit(1)
}

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "A question answer file")
	quizTime := flag.Int("timer", 30, "Quiz time")
	flag.Parse()

	fmt.Println("Preparing quiz...")
	f := openFile(*csvFileName)
	defer f.Close()
	records := readContent(f)

	fmt.Println("Ready")
	fmt.Println("Enter any key to start...")
	fmt.Scanf("%c")

	var correct, total int
	total = len(records)

	problems := parseLines(records)

	timer := time.NewTimer(time.Duration(*quizTime) * time.Second)

	for i, p := range problems {
		fmt.Printf("Q.%d %s ", i + 1,  p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			printResult(correct, total)
			return
		case answer := <-answerCh:
			if strings.ToLower(strings.TrimSpace(answer)) == strings.ToLower(strings.TrimSpace(p.a)) {
				correct++
			}
		}
	}

	printResult(correct, total)
}

func parseLines (lines [][]string) []problem {
	problems := make([] problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			q : line[0],
			a : line[1],
		}
	}

	return problems
}
type problem struct {
	q string
	a string
}

func printResult(correct int, total int) {
	fmt.Println("\nCorrect: ", correct)
	fmt.Println("Total: ", total)
}

func readContent(f *os.File) [][]string {
	reader := csv.NewReader(f)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Error reading records: ", err)
	}
	return records
}

func openFile(file string) *os.File {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	return f
}

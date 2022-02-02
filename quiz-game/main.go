package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	csvFileName string
	timeLimit   int
	shuffleVar  bool
)

// Initiate and declare default flags
func init() {
	flag.StringVar(&csvFileName, "csv",
		"problems.csv", "a csv file in the format of question,answer")
	flag.BoolVar(&shuffleVar, "shuffle", true, "shuffle the problems if true, else ignore")
	flag.IntVar(&timeLimit, "limit", 3, "time limit for the test in seconds")
}

type problem struct {
	// Question for the problem
	question string
	// Answer for the problem
	answer string
	// Available options for the problem
	options []string
}

// Reads the CSV file and return the records present in it
func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read input file %v! %v"+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to parse file as CSV for %V!\n %v"+filePath, err)
	}

	return records
}

func parseLines(lines [][]string) []problem {
	res := make([]problem, len(lines))

	for ind, line := range lines {
		res[ind] = problem{
			question: strings.TrimSpace(line[0]),
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return res
}

func main() {
	// Parsing all the flags
	flag.Parse()

	// Reads the CSV file `problems.csv`
	questionnaire := parseLines(readCsvFile(csvFileName))

	var value string
	var ordering []int

	if shuffleVar {
		ordering = rand.Perm(len(questionnaire))
	} else {
		for ind := 0; ind < len(questionnaire); ind++ {
			ordering = append(ordering, ind)
		}
	}
	correct := 0

	fmt.Printf("\n\tPress the Enter key to start!")
	fmt.Scanf("%*c")

	start := time.Now()
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	for ind, _ := range questionnaire {
		prob := questionnaire[ordering[ind]]
		fmt.Printf("\tQ.%v\t%v = ", ind+1, prob.question)
		answerCh := make(chan string)

		go func() {
			fmt.Scanf("%s\n", &value)
			answerCh <- value
		}()

		select {
		// When the scheduled time occurs, stop everything
		case <-timer.C:
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Printf("\n\n\tYou scored %v out of %v!", correct, len(questionnaire))
			fmt.Printf("\n\tTime taken: %2.4v seconds!\n", elapsed)
			return

		// When the answer arrives before scheduled time
		case answer := <-answerCh:
			if strings.Compare(answer, prob.answer) == 0 {
				correct++
			}
		}
	}

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("\n\tYou scored %v out of %v!\n", correct, len(questionnaire))
	fmt.Printf("\tTime taken: %2.4v seconds!\n", elapsed)
}

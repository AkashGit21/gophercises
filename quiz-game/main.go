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

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

var limitVar int

func init() {
	// Initiate shuffle and limit flags
	flag.Bool("shuffle", false, "Shuffle the questionnaire if true else ignore")
	flag.IntVar(&limitVar, "limit", 30, "Time limit for the test")
}

func main() {
	// Parsing all the flags
	flag.Parse()

	// Reads the CSV file `problems.csv`
	questionnaire := readCsvFile("problems.csv")
	fmt.Println("Limit is: ", limitVar)

	correct := 0
	fmt.Printf("\n\tPress the Enter key to start!")
	fmt.Scanf("%*c")

	start := time.Now()
	timer := time.NewTimer(time.Duration(limitVar) * time.Second)

	go func() {
		var value string
		for ind, question := range questionnaire {
			fmt.Printf("\tQ.%v\t%v = ", ind+1, question[0])
			fmt.Scanf("%s", &value)
			if strings.Compare(strings.TrimSpace(value), question[1]) == 0 {
				correct++
			}
		}
	}()

	<-timer.C

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("\n\tYou scored %v out of %v!\n", correct, len(questionnaire))
	fmt.Printf("\tTime taken: %2.4v seconds!\n", elapsed)
}

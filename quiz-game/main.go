package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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

func main() {
	flag.Bool("shuffle", false, "Shuffle the questionnaire if true else ignore")
	flag.Int("limit", 5, "Time limit for the test")
	flag.Parse()

	questionnaire := readCsvFile("problems.csv")

	var value string
	correct := 0
	for ind, question := range questionnaire {
		fmt.Printf("Q.%v\t%v = ", ind+1, question[0])
		fmt.Scanf("%s", &value)
		if strings.Compare(strings.TrimSpace(value), question[1]) == 0 {
			correct++
		}
	}
	fmt.Printf("You scored %v out of %v!\n", correct, len(questionnaire))
}

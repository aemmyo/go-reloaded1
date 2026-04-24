package main

import (
	"fmt"
	"os"
	"strings"
)

func ProcessText(s string) string {
	result := HandleCases(s)
	result = HexBin(result)
	result = aOrAn(result)
	result = ItQuote(result)
	result = ItPunc(result)
	return result
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . sample.txt result.txt")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error in reading input file")
		return
	}
	result := NewLine(string(data))
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error in writing output file")
		return
	}

}

func NewLine(text string) string {
	s := strings.Split(text, "\n")
	for i, line := range s {
		s[i] = ProcessText(line)
	}
	return strings.Join(s, "\n")
}


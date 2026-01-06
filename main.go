package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("ERROR = Enter the command in this format - run main.go <filename>")
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error in OPENING the file %v\n", err)
	}

	defer file.Close()
	var logs []string
	lineCount := 0
	INFO := 0
	Warn := 0
	Error := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
		line := scanner.Text()
		if line != "" {
			logs = append(logs, line)
			if strings.Contains(line, "INFO") {
				INFO++
			} else if strings.Contains(line, "WARN") {
				Warn++
			} else if strings.Contains(line, "ERROR") {
				Error++
			} else {
				// Do nothing for unrecognized log levels
			}

		}

	}
	for _, Log := range logs {
		fmt.Println(Log)
	}
	fmt.Printf("\n--- Done! Total lines: %d ---\n", lineCount)

	fmt.Printf("INFO: %d\n Warn: %d\n Error: %d\n", INFO, Warn, Error)
}

func processLine(line string, INFO *int, Warn *int, Error *int) {
	
}
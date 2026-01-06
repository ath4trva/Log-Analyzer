package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LogStats struct {
	TotalLines int
	InfoCount  int
	WarnCount  int
	ErrorCount int
	Lines      []string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]

	stats, err := parseLogFile(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, line := range stats.Lines {
		fmt.Println(line)
	}

	printSummary(stats)
}

func parseLogFile(path string) (LogStats, error) {
	file, err := os.Open(path)
	if err != nil {
		return LogStats{}, err
	}
	defer file.Close()

	var s LogStats
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		s.TotalLines++
		s.Lines = append(s.Lines, line)

		if strings.Contains(line, "INFO") {
			s.InfoCount++
		} else if strings.Contains(line, "WARN") {
			s.WarnCount++
		} else if strings.Contains(line, "ERROR") {
			s.ErrorCount++
		}
	}

	return s, scanner.Err()
}

func printSummary(s LogStats) {
	fmt.Printf("\n--- LOG ANALYSIS COMPLETE ---\n")
	fmt.Printf("Total processed: %d\n", s.TotalLines)
	fmt.Printf("INFO messages:   %d\n", s.InfoCount)
	fmt.Printf("WARN messages:   %d\n", s.WarnCount)
	fmt.Printf("ERROR messages:  %d\n", s.ErrorCount)
	fmt.Println("-----------------------------")
}

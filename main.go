package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LogStats struct {
	Counts     map[string]int
	TotalLines int
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
	s := LogStats{
		Counts: make(map[string]int),
	}

	levels := []string{"INFO", "WARN", "ERROR"}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		s.TotalLines++
		s.Lines = append(s.Lines, line)
		match:= false
		for _, level := range levels {
			if strings.Contains(line, level) {
				s.Counts[level]++
				match = true
				break
			}
		}
		if !match{
			s.Counts["UNKNOWN"]++
		}
	}

	return s, scanner.Err()
}

func printSummary(s LogStats) {
	fmt.Printf("\n--- LOG ANALYSIS COMPLETE ---\n")
	fmt.Printf("Total processed: %d\n", s.TotalLines)

	for level, count := range s.Counts {
		fmt.Printf("%-5s messages:   %d\n", level, count)
	}
	fmt.Println("-----------------------------")
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type LogStats struct {
	Counts     map[string]int
	TotalLines int
	Lines      []string
	mu 	   sync.Mutex
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}
	start:= time.Now()

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

	duration:= time.Since(start)
	fmt.Printf("Processing time: %v\n", duration)
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

	lineChan:= make(chan string, 100)
	var wg sync.WaitGroup
	numWorkers:=4	

	

	for i:=0; i<numWorkers; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			levels:=[]string{"INFO", "WARN", "ERROR"}
			for line:= range lineChan {
				match:=false
				for _, level:= range levels{
					if strings.Contains(line, level){
						s.increment(level)
						match = true
						break
				}
			}

			if !match{
				s.increment("UNKNOWN")
			}
		}
	}()
	}
	scanner:= bufio.NewScanner(file)
	for scanner.Scan(){
		line:= scanner.Text()
		if strings.TrimSpace(line)==""{
			continue
		}
		s.TotalLines++
		s.Lines = append(s.Lines, line)
		lineChan<-line
	}
	close(lineChan)
	wg.Wait()
	return s, scanner.Err()
	
}
 func(s *LogStats) increment(level string){
	s.mu.Lock()
	s.Counts[level]++
	s.mu.Unlock()
 }

func printSummary(s LogStats) {
	fmt.Printf("\n--- LOG ANALYSIS COMPLETE ---\n")
	fmt.Printf("Total processed: %d\n", s.TotalLines)

	for level, count := range s.Counts {
		fmt.Printf("%-5s messages:   %d\n", level, count)
	}
	fmt.Println("-----------------------------")
}

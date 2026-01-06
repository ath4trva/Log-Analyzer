# Concurrent Log Analyzer (Go)

A CLI-based log analyzer written in Go that efficiently processes large log files using concurrency.  
The tool reads log files, detects log levels, and produces a summarized report.

This project was built to strengthen understanding of Go fundamentals such as file I/O, goroutines, channels, worker pools, mutexes, and synchronization.

---

## Features

- Reads log files line by line
- Detects log levels:
  - INFO
  - WARN
  - ERROR
  - UNKNOWN
- Counts total number of log lines
- Concurrent processing using a worker pool
- Thread-safe aggregation of results
- Clean CLI output

---

## Project Structure

.
├── main.go
├── app.log
└── README.md


---

## How It Works

1. The log file is read line by line using `bufio.Scanner`
2. Each line is sent into a buffered channel
3. Multiple worker goroutines consume log lines concurrently
4. Workers detect log levels in each line
5. Shared statistics are updated safely using a mutex
6. A `sync.WaitGroup` ensures all workers complete before exiting
7. A summary of log statistics is printed

---

## Usage

### Run the program

```bash
go run main.go <logfile>
$ go run main.go log.txt

2026-01-06 09:15:02 ERROR Authentication failed: Invalid API key provided
2026-01-06 09:18:45 ERROR FileSystemException: Permission denied at /var/www/uploads/
2026-01-06 09:25:30 WARN High disk usage detected: 88% on volume /dev/sda1
...
2026-01-06 10:25:00 CRITICAL Temperature threshold exceeded!
2026-01-06 10:30:12 PANIC Segment fault at address 0x0045f

--- LOG ANALYSIS COMPLETE ---
Total processed: 15
INFO  messages:   3
WARN  messages:   4
ERROR messages:   3
UNKNOWN messages: 5
-----------------------------
⚡ Execution Time: 2.1055ms

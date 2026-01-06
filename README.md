# Concurrent Log Analyzer (Go)

A CLI-based log analyzer written in Go that processes large log files efficiently using concurrency.  
It reads log files, identifies log levels, and produces a summarized analysis.

This project was built to deeply understand Go fundamentals such as file I/O, goroutines, channels, worker pools, mutexes, and synchronization.

---

## Features

- Reads log files line by line
- Detects log levels:
  - INFO
  - WARN
  - ERROR
  - UNKNOWN
- Counts total log lines
- Concurrent processing using a worker pool
- Thread-safe aggregation of results
- Clean and readable CLI output

---

## Project Structure
.
├── main.go
├── log.txt
└── README.md


---

## How It Works (High-Level)

1. The log file is read line by line using `bufio.Scanner`
2. Each line is sent to a buffered channel
3. Multiple worker goroutines consume lines concurrently
4. Workers analyze log levels
5. Shared statistics are updated safely using a mutex
6. A `sync.WaitGroup` ensures all workers finish before exit
7. A summary report is printed

---

## Usage

### Run the program

```bash
go run main.go <logfile>

## Sample Output

2026-01-06 09:15:02 ERROR Authentication failed: Invalid API key provided
2026-01-06 09:18:45 ERROR FileSystemException: Permission denied at /var/www/uploads/
2026-01-06 09:22:10 ERROR OutOfMemoryError: Java heap space exhausted
2026-01-06 09:25:30 WARN High disk usage detected: 88% on volume /dev/sda1
2026-01-06 09:30:12 WARN Deprecated method call: 'getUserDataV1' will be removed soon
2026-01-06 09:35:55 WARN Latency spike: Database query took 1500ms
2026-01-06 09:40:18 WARN Retrying connection to Message Queue (Attempt 2/5)
2026-01-06 09:45:00 INFO System startup complete: Listening on port 8080
2026-01-06 09:50:22 INFO User ID 4492 successfully updated profile settings
2026-01-06 10:00:05 INFO Daily backup job initiated for 'customer_db'
2026-01-06 10:15:22 DEBUG Initializing internal cache...
2026-01-06 10:20:45 TRACE Packet received from 192.168.1.1
2026-01-06 10:22:10 SYSTEM Kernel update initiated
2026-01-06 10:25:00 CRITICAL Temperature threshold exceeded!
2026-01-06 10:30:12 PANIC Segment fault at address 0x0045f

--- LOG ANALYSIS COMPLETE ---
Total processed: 15
INFO  messages:   3
UNKNOWN messages:   5
ERROR messages:   3
WARN  messages:   4
-----------------------------
Processing time: 2.1055ms


# nutwrk

A lightweight network diagnostics and benchmarking CLI written in Go.

**nutwrk** helps developers and enthusiasts measure network performance directly from the terminal by providing latency analysis, packet loss detection, jitter statistics, and download speed benchmarking.

Built from scratch as a learning project to understand how modern networking and speed-testing tools work under the hood.

---

## Features

### Implemented

* DNS Resolution
* TCP Latency Measurement
* Packet Loss Detection
* Average Latency Statistics
* Jitter Calculation
* Download Speed Benchmark
* Download Progress Tracking
* Custom Download Test Sizes
* Configurable Timeout & Ports
* IPv4 Resolution

### Planned

* Upload Speed Testing
* Multi-stream Download Benchmarking
* Automatic Server Selection
* ICMP Ping Support
* Historical Metrics
* Exportable Reports
* Geo-based Server Discovery

---

## Why nutwrk?

Most developers use networking tools every day but rarely understand how they work internally.

nutwrk is designed to be both:

1. A practical network benchmarking utility.
2. A hands-on systems programming project for learning networking concepts in Go.

The project explores:

* DNS lookups
* TCP connection establishment
* Latency measurement
* Packet loss detection
* Jitter calculation
* Throughput measurement
* HTTP streaming
* Concurrent network applications

---

# Installation & Usage

## Download a Prebuilt Release

1. Go to the Releases page.
2. Download the archive for your operating system.
3. Extract the archive.
4. Run the executable.

### Linux

```bash
chmod +x nutwrk
./nutwrk --help
```

### Windows

```powershell
nutwrk.exe --help
```

### macOS

```bash
chmod +x nutwrk
./nutwrk --help
```

---

## Build From Source

### Prerequisites

- Go 1.24 or later
- Git

### Clone the Repository

```bash
git clone https://github.com/sanskarOH/speedT.git
cd speedT
```

### Install Dependencies

```bash
go mod tidy
```

### Build

```bash
go build -o nutwrk
```

### Run

```bash
./nutwrk --help
```

---

# Commands

## TCP Latency Check

Measure TCP connection latency to a host.

### Usage

```bash
nutwrk ltc <host> [flags]
```

### Flags

| Flag | Description | Default |
|--------|-------------|---------|
| `-c` | Number of requests | `4` |
| `-p` | Port | `443` |
| `-t` | Timeout in seconds | `3` |

### Examples

```bash
nutwrk ltc google.com
```

```bash
nutwrk ltc google.com -c 10
```

```bash
nutwrk ltc google.com -p 80
```

```bash
nutwrk ltc google.com -c 5 -p 443 -t 2
```

---

# Example Output

```text
Resolving google.com...

PING google.com (142.250.182.142)

Reply from 142.250.182.142: time=54ms
Reply from 142.250.182.142: time=58ms
Reply from 142.250.182.142: time=55ms
Reply from 142.250.182.142: time=53ms

--- Statistics ---
Packets Sent:     4
Packets Received: 4
Packets Lost:     0
Average Latency:  55ms
```

---

# Development

Run directly without building:

```bash
go run . ltc google.com
```

Run tests:

```bash
go test ./...
```

Format code:

```bash
go fmt ./...
```

---

# Troubleshooting

### Command Not Found

Make sure the binary is executable:

```bash
chmod +x nutwrk
```

### Connection Timed Out

Verify:

- Internet connectivity
- Hostname correctness
- Firewall settings
- Port availability

---

## Commands

### LTC (Latency Test Command)

Measure TCP latency to a host.

```bash
nutwrk ltc google.com
```

Example:

```text
pinging google.com [142.250.xxx.xxx] on port 443

Reply from 142.250.xxx.xxx time= 41 ms
Reply from 142.250.xxx.xxx time= 39 ms
Reply from 142.250.xxx.xxx time= 42 ms
Reply from 142.250.xxx.xxx time= 40 ms

Ping Statistics
---------------
Packets Sent     : 4
Packets Received : 4
Packet Loss      : 0%
Average Latency  : 40 ms
Jitter           : 1.3 ms
```

### Download Benchmark

Measure download throughput using configurable test file sizes.

```bash
nutwrk test
```

```bash
nutwrk test -s 10mb
```

```bash
nutwrk test -s 100mb
```

```bash
nutwrk test -s 1gb
```

Example:

```text
Connected to the server successfully
Status: 200 OK

Progress: 25%
Progress: 50%
Progress: 75%
Progress: 100%

--- Download Statistics ---

File Size          : 100.00 MB
Downloaded         : 100.00 MB
Time Taken         : 12.34 s
Download Speed     : 67.97 Mbps
```

---

## Project Structure

```text
nutwrk/

├── cmd/
│   ├── root.go
│   ├── ltc.go
│   └── test.go
│
├── internal/
│   ├── ltc/
│   │   └── ltc.go
│   │
│   └── test/
│       └── test.go
│
├── go.mod
├── go.sum
└── README.md
```

---

## Roadmap

### Networking

* [x] DNS Resolution
* [x] TCP Latency Measurement
* [x] Packet Loss Detection
* [x] Jitter Calculation
* [x] Download Benchmark

### Performance Testing

* [ ] Upload Benchmark
* [ ] Multi-stream Downloads
* [ ] Adaptive Benchmark Duration
* [ ] Dynamic Server Selection

### Systems Programming

* [ ] Worker Pools
* [ ] Concurrent Downloads
* [ ] Context Cancellation
* [ ] Metrics Collection

---

## Technologies

* Go
* Cobra CLI
* TCP
* DNS
* HTTP
* TLS

---
## Development Phases

### Phase 1 — Core Networking

* [x] Project Setup
* [x] Cobra CLI Integration
* [x] Command Structure
* [x] DNS Resolution
* [x] IPv4 Address Extraction
* [x] TCP Connection Testing

### Phase 2 — Latency Analysis

* [x] TCP-based Latency Measurement
* [x] Packet Loss Detection
* [x] Average Latency Statistics
* [x] Min / Max Latency Tracking
* [x] Jitter Calculation

### Phase 3 — Download Benchmarking

* [x] HTTP Connectivity Testing
* [x] Streaming Downloads
* [x] Throughput Measurement
* [x] Download Progress Tracking
* [x] Configurable Test Sizes
* [x] Human-readable Statistics

### Phase 4 — Upload Benchmarking

* [ ] HTTP Upload Benchmark
* [ ] Upload Throughput Measurement
* [ ] Upload Progress Tracking
* [ ] Configurable Upload Payloads

### Phase 5 — Advanced Benchmarking

* [ ] Multi-stream Downloads
* [ ] Parallel Connection Testing
* [ ] Adaptive Benchmark Duration
* [ ] Dynamic Test Scaling
* [ ] Sustained Throughput Analysis

### Phase 6 — Intelligent Server Selection

* [ ] Server Discovery
* [ ] Latency-based Ranking
* [ ] Geo-aware Selection
* [ ] Automatic Best Server Detection

### Phase 7 — Systems & Concurrency

* [ ] Goroutine Worker Pools
* [ ] Context Cancellation
* [ ] Concurrent Benchmark Execution
* [ ] Rate Limiting
* [ ] Connection Pooling

### Phase 8 — Observability

* [ ] Metrics Collection
* [ ] Historical Benchmark Results
* [ ] JSON Output Support
* [ ] Exportable Reports
* [ ] Prometheus Integration

### Phase 9 — Protocol Expansion

* [ ] ICMP Ping Support
* [ ] UDP Benchmarking
* [ ] Traceroute
* [ ] DNS Benchmarking
* [ ] TLS Handshake Timing

### Phase 10 — Custom Infrastructure

* [ ] Custom Benchmark Server
* [ ] Self-hosted Test Nodes
* [ ] Distributed Benchmarking
* [ ] Regional Performance Comparison
---

## Author

**Sanskar Diwedi**

GitHub: https://github.com/sanskarOH

Built to learn networking, systems programming, and performance engineering through real-world tooling.

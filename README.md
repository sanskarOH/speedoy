# nutwrk

A lightweight, fast, and educational network diagnostics CLI written in Go.

nutwrk helps measure network performance by testing latency, jitter, download speed, upload speed, and packet loss directly from the terminal.

This project is built as a learning exercise to explore:

* Go
* Networking fundamentals
* TCP & UDP
* DNS Resolution
* Concurrency with Goroutines
* CLI Development with Cobra
* System Design

---

## Features

### Current

* DNS Resolution
* Latency Measurement

### Planned

* Jitter Calculation
* Download Speed Test
* Upload Speed Test
* Multi-threaded Testing
* Packet Loss Detection
* Automatic Server Selection
* Metrics & Monitoring
* ICMP Ping Support

---

## Project Goals

The objective of nutwrk is not only to provide a speed testing utility but also to serve as a hands-on learning project for understanding how network diagnostic tools work under the hood.

By building nutwrk, you'll learn:

* How DNS works
* How TCP connections are established
* How latency is measured
* How bandwidth is calculated
* How speed test servers operate
* How concurrent network applications are built in Go

---

## Architecture

```text
User
 │
 ▼
nutwrk CLI
 │
 ├── DNS Resolution
 │
 ├── TCP Connection
 │
 ├── Latency Measurement
 │
 ├── Download Benchmark
 │
 └── Upload Benchmark
```

---

## Project Structure

```text
nutwrk/

├── cmd/
│   ├── root.go
│   ├── ping.go
│   ├── download.go
│   └── upload.go
│
├── internal/
│   ├── ping/
│   │   ├── ping.go
│   │   ├── dns.go
│   │   ├── tcp.go
│   │   └── stats.go
│   │
│   ├── download/
│   ├── upload/
│   ├── jitter/
│   └── packetloss/
│
├── pkg/
│   ├── network/
│   ├── metrics/
│   └── utils/
│
├── tests/
│
├── go.mod
└── README.md
```

---

## Installation

### Clone Repository

```bash
git clone https://github.com/sanskarOH/nutwrk.git

cd nutwrk
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
./nutwrk
```

---

## Usage

### Check latency of a Host

```bash
nutwrk ltc google.com
```

Example Output:

```text
Host: google.com
IP: 142.250.xxx.xxx

Latency: 14ms
```

---

### check latency Multiple Times

```bash
nutwrk ltc google.com -c 5
```

Example Output:

```text
Reply 1: 14ms
Reply 2: 12ms
Reply 3: 15ms
Reply 4: 13ms
Reply 5: 14ms

Min: 12ms
Max: 15ms
Avg: 13.6ms
```

---

## Development Roadmap

### Phase 1

* [x] Project Setup
* [x] DNS Resolution
* [x] Latency Check
* [x] Latency Statistics

### Phase 2

* [x] Jitter Calculation

### Phase 3

* [x] Download Speed Test
* [ ] Throughput Measurement

### Phase 4

* [ ] Upload Speed Test

### Phase 5

* [ ] Multi-threaded Connections
* [ ] Goroutine Worker Pool

### Phase 6

* [ ] Custom Benchmark Server

### Phase 7

* [ ] Server Discovery
* [ ] Geo-based Selection


---

## Technologies

* Go
* Cobra CLI
* TCP
* UDP
* DNS
* HTTP

---

## Learning Outcomes

This project explores:

### Networking

* DNS
* TCP
* UDP
* HTTP
* TLS
* Packet Routing

### Go

* Packages
* Modules
* Interfaces
* Error Handling
* Goroutines
* Channels
* Context

### System Design

* Client-Server Architecture
* Benchmarking
* Concurrent Systems
* Observability

---

## License

MIT License

---

## Author

**Sanskar Diwedi**

GitHub: https://github.com/sanskarOH

Built to learn how networking tools work under the hood.

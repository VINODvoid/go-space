# ratelimiter

A token-bucket rate limiter — the fourth intermediate project in the go-space learning series.

---

## Usage

```bash
go run main.go
```

### Example

```bash
job1 -- 18:40:30
job2 -- 18:40:33
job3 -- 18:40:36
job4 -- 18:40:39
job5 -- 18:40:42
```

Jobs are processed at a fixed interval — one every N seconds.

---

## Concepts Covered

- `time.NewTicker` — fires a signal on a fixed interval
- `ticker.C` — the channel that receives on each tick
- `<-ticker.C` — blocks until next tick, providing rate control without sleep loops
- `defer ticker.Stop()` — frees the internal goroutine, prevents resource leaks
- `time.Duration` — Go's type for expressing time intervals
- Channels as flow control — blocking behavior as a timing mechanism

---

## Why This Matters

Without a rate limiter, code can hammer external APIs, rack up costs, or get banned.
The `<-ticker.C` pattern is idiomatic Go — no sleep loops, no manual timing, just channel blocking.

For production use, see `golang.org/x/time/rate` which implements a full token bucket with burst support.

---

![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=flat-square&logo=go&logoColor=white)
![Level](https://img.shields.io/badge/Level-Intermediate-orange?style=flat-square)
![Concepts](https://img.shields.io/badge/Concepts-channels%20%7C%20tickers%20%7C%20goroutines%20%7C%20rate%20limiting-blue?style=flat-square)

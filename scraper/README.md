# scraper

A concurrent web scraper — the second intermediate project in the go-space learning series.

---

## Usage

```bash
go run main.go <url1> <url2> <url3> ...
```

### Example

```bash
go run main.go https://example.com https://httpbin.org/get https://httpbin.org/status/404
```

### Output

```
ERROR https://example.com: URL is invalid
200 https://httpbin.org/get — 19 words
404 https://httpbin.org/status/404 — 0 words
```

All URLs are fetched **concurrently** — results arrive as goroutines complete.

---

## Concepts Covered

- Goroutines — one per URL, launched with `go func()`
- `sync.WaitGroup` — tracking when all goroutines finish
- Channels — passing results safely between goroutines
- `close(ch)` + `for range ch` — draining a channel until done
- `http.Get` — making HTTP client requests
- `io.ReadAll` — reading response body
- Closure variable capture bug — why `go func(u string){}(u)` matters

---

![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=flat-square&logo=go&logoColor=white)
![Level](https://img.shields.io/badge/Level-Intermediate-orange?style=flat-square)
![Concepts](https://img.shields.io/badge/Concepts-goroutines%20%7C%20channels%20%7C%20WaitGroup%20%7C%20concurrency-blue?style=flat-square)

# httpserver

A simple HTTP server — the fourth project in the go-space learning series.

---

## Usage

```bash
go run main.go
```

Server starts on `http://localhost:8080`.

### Routes

```bash
curl http://localhost:8080/                      # Welcome to my server
curl "http://localhost:8080/hello?name=John"     # Hello John
curl http://localhost:8080/health                # OK
```

### Terminal logs

```
Server is listening on port 8080:
GET /
GET /hello
GET /health
```

---

## Concepts Covered

- `net/http` package — routing and server setup
- Handler functions (`http.ResponseWriter`, `*http.Request`)
- Writing responses with `fmt.Fprintln` and `fmt.Fprintf`
- Reading query parameters with `r.URL.Query().Get()`
- Inspecting requests via `r.Method` and `r.URL.Path`
- Middleware pattern — wrapping handlers for cross-cutting concerns

---

![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=flat-square&logo=go&logoColor=white)
![Level](https://img.shields.io/badge/Level-Beginner-brightgreen?style=flat-square)
![Concepts](https://img.shields.io/badge/Concepts-net%2Fhttp%20%7C%20handlers%20%7C%20middleware%20%7C%20routing-blue?style=flat-square)

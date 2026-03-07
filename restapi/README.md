# restapi

A REST API for managing books — the first intermediate project in the go-space learning series.

---

## Usage

```bash
go run .
```

Server starts on `http://localhost:8080`.

### Endpoints

```bash
# List all books
curl http://localhost:8080/books

# Get a book by ID
curl http://localhost:8080/books/1

# Add a new book
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"title":"Clean Code","author":"Martin","year":2008}'

# Delete a book
curl -X DELETE http://localhost:8080/books/1
```

### Example Response

```json
{"id":1,"title":"The Go Programming Language","author":"Donovan","year":2015}
```

---

## Concepts Covered

- Struct tags for JSON field mapping (`json:"title"`)
- JSON encoding and decoding with `encoding/json`
- Handler methods on structs for shared state
- HTTP status codes (200, 201, 400, 404)
- Routing GET/POST/DELETE on the same path via `r.Method`
- Manual URL path parsing for path parameters
- Two-value returns `(Book, bool)` for found/not found pattern

---

![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=flat-square&logo=go&logoColor=white)
![Level](https://img.shields.io/badge/Level-Intermediate-orange?style=flat-square)
![Concepts](https://img.shields.io/badge/Concepts-JSON%20%7C%20REST%20%7C%20interfaces%20%7C%20status%20codes-blue?style=flat-square)

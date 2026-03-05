# calculator

A CLI calculator — the first project in the go-space learning series.

---

## Usage

```bash
go run main.go <number> <operator> <number>
```

### Examples

```bash
go run main.go 10 + 5    # 15
go run main.go 10 - 3    # 7
go run main.go 6 '*' 4   # 24
go run main.go 10 / 2    # 5
go run main.go 10 / 0    # cannot divide by zero
```

---

## Concepts Covered

- `package main` and program entry points
- CLI argument parsing with `os.Args`
- Type conversion with `strconv.ParseFloat`
- `switch` statements
- Input validation with `len()`
- Multiple return values and the `_` blank identifier
- Early exits with `return`

---

![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=flat-square&logo=go&logoColor=white)
![Level](https://img.shields.io/badge/Level-Beginner-brightgreen?style=flat-square)
![Concepts](https://img.shields.io/badge/Concepts-types%20%7C%20functions%20%7C%20switch%20%7C%20os.Args-blue?style=flat-square)

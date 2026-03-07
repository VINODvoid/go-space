# wordcounter

A CLI file word counter — the third project in the go-space learning series.

---

## Usage

```bash
go run main.go <filename>
```

### Examples

```bash
go run main.go sample.txt    # count words in sample.txt
go run main.go missing.txt   # open missing.txt: no such file or directory
```

---

## Concepts Covered

- File I/O with `os.Open`
- `defer` for guaranteed resource cleanup
- Buffered reading with `bufio.Scanner`
- Word splitting with `strings.Fields`
- `map[string]int` for frequency counting
- Proper error handling (`if err != nil` — no more `_`)
- Extracting and sorting map keys with `sort.Strings`

---

![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=flat-square&logo=go&logoColor=white)
![Level](https://img.shields.io/badge/Level-Beginner-brightgreen?style=flat-square)
![Concepts](https://img.shields.io/badge/Concepts-file%20I%2FO%20%7C%20maps%20%7C%20bufio%20%7C%20error%20handling-blue?style=flat-square)

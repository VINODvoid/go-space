# todo

A CLI to-do manager (in-memory) — the second project in the go-space learning series.

---

## Usage

```bash
go run . add "buy milk"    # add a new todo
go run . list              # list all todos
go run . done 1            # mark todo #1 as done
go run . delete 1          # delete todo #1
```

> Note: data is in-memory only — it does not persist between runs. File persistence is added in Project 3.

---

## Concepts Covered

- Structs and custom types
- Named slice types (`type TodoList []Todo`)
- Methods with pointer receivers (`*TodoList`)
- Struct literals
- Range loops and index-based mutation
- Multi-file packages (`go run .`)
- `strconv.Atoi` for string to int conversion
- Exported fields (capitalized = public in Go)

---

![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=flat-square&logo=go&logoColor=white)
![Level](https://img.shields.io/badge/Level-Beginner-brightgreen?style=flat-square)
![Concepts](https://img.shields.io/badge/Concepts-structs%20%7C%20methods%20%7C%20slices-blue?style=flat-square)

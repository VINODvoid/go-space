# taskmanager

A CLI task manager with file persistence — the third intermediate project in the go-space learning series.

---

## Usage

```bash
go run . -add "buy milk"    # add a task
go run . -list              # list all tasks
go run . -done 1            # mark task #1 as done
go run . -delete 1          # delete task #1
```

Tasks are saved to `tasks.json` and persist between runs.

### Example

```bash
$ go run . -add "buy milk"
$ go run . -add "write code"
$ go run . -list
buy milk ( false )
write code ( false )
$ go run . -done 1
$ go run . -list
buy milk ( true )
write code ( false )
```

---

## Concepts Covered

- `flag` package — CLI flag parsing (`-add`, `-list`, `-done`)
- Pointer flags — `flag.String` returns `*string`, dereference with `*`
- `os.Create` — creates file if not exists, truncates if it does
- JSON file persistence — `json.NewEncoder/Decoder` for reading and writing
- Load on startup, save after every mutation
- Separation of concerns — `store.go` for data, `main.go` for CLI

---

![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=flat-square&logo=go&logoColor=white)
![Level](https://img.shields.io/badge/Level-Intermediate-orange?style=flat-square)
![Concepts](https://img.shields.io/badge/Concepts-flags%20%7C%20JSON%20%7C%20file%20persistence%20%7C%20CLI-blue?style=flat-square)

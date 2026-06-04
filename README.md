# 🐹 golang-lab

A hands-on learning repository for experimenting with **Go (Golang)**. This project serves as a structured collection of small, focused mini-projects built to explore core language concepts, standard library usage, concurrency patterns, networking, and backend development practices in Go.

> **Goal:** Learn Go by building real, working things — not just reading docs.

---

## 📁 Repository Structure

```
golang-lab/
├── todocli/           # CLI-based Todo manager (file persistence, CRUD)
├── ...                # More mini-projects coming soon
└── README.md
```

---

## 🚀 Mini-Projects

### 1. `todocli` — Command-Line Todo Manager

A fully functional CLI todo application demonstrating foundational Go patterns.

**Concepts covered:**
- Struct design and method receivers (`*Todos`)
- Pointer semantics and slice manipulation
- JSON marshalling / unmarshalling (`encoding/json`)
- File I/O with the `os` and `path/filepath` packages
- Error handling idioms (`errors.New`, `fmt.Errorf`)
- Working with `time.Time` for timestamps

**Features:**
- Add, complete, delete, and list todo items
- Persist todos to a JSON file in the OS temp directory
- Graceful handling of missing or empty files on load

**Usage example:**
```go
var todos todocli.Todos

todos.LoadFromFile("todos.json")
todos.Add("Learn Go interfaces")
todos.Add("Build a REST API")
todos.Complete(1)
todos.Delete(2)
todos.SaveToFile("todos.json")

for i, item := range todos.List() {
    fmt.Printf("[%d] %v\n", i+1, item)
}
```

**Key data model:**
```go
type item struct {
    Done        bool      `json:"done"`
    CreatedAt   time.Time `json:"created_at"`
    CompletedAt time.Time `json:"completed_at,omitempty"`
    Task        string    `json:"description"`
}
```

---

## 🧠 Topics This Repo Will Explore

| Area | Status |
|---|---|
| Structs, methods, interfaces | ✅ Started (`todocli`) |
| Error handling patterns | ✅ Started (`todocli`) |
| File I/O & JSON | ✅ Started (`todocli`) |
| CLI argument parsing (`flag`, `cobra`) | 🔜 Planned |
| Concurrency (`goroutines`, `channels`) | 🔜 Planned |
| HTTP servers & REST APIs | 🔜 Planned |
| Middleware & routing | 🔜 Planned |
| Database access (`database/sql`, GORM) | 🔜 Planned |
| Testing (`testing` package, table-driven tests) | 🔜 Planned |
| Context & cancellation | 🔜 Planned |

---

## 🛠 Prerequisites

- [Go 1.21+](https://go.dev/dl/) installed
- Basic familiarity with a terminal

Verify your installation:
```bash
go version
```

---

## ▶️ Running a Mini-Project

Each mini-project lives in its own package directory. To run or test one:

```bash
# Navigate to a project
cd todocli

# Run tests (if available)
go test ./...

# Or integrate the package in a main.go at the repo root
go run main.go
```

---

## 📖 Learning Philosophy

Each mini-project in this repo follows a simple philosophy:

1. **Small scope** — one concept or feature area per project
2. **Real output** — every project produces something that actually works
3. **Commented intent** — code comments explain *why*, not just *what*
4. **Idiomatic Go** — following community conventions and effective Go guidelines

---

## 📚 Resources

- [The Go Tour](https://go.dev/tour/) — official interactive intro
- [Effective Go](https://go.dev/doc/effective_go) — idiomatic Go guide
- [Go by Example](https://gobyexample.com/) — annotated code examples
- [Go Standard Library](https://pkg.go.dev/std) — official package docs

---

## 📝 License

This repository is licensed under the [MIT License](./LICENSE). Feel free to fork and use it as a template for your own Go experiments.
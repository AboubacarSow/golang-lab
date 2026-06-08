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

`todocli` is a completed CLI todo application that demonstrates idiomatic Go, JSON persistence, and simple flag-based command handling.

**What it supports:**
- Add a new todo item
- List all todos in a formatted table
- Mark a todo as completed by index
- Delete a todo by index
- Persist todos to `todos.json` in the current working directory

**Concepts covered:**
- Go structs and method receivers
- Pointer semantics and slice updates
- JSON serialization and deserialization
- File I/O and path handling
- CLI argument parsing with `flag`
- Error handling and exit status
- Third-party package usage for terminal output

**CLI usage:**
```bash
cd todocli

go run ./cmd/todo -list

go run ./cmd/todo -add "Buy groceries"

go run ./cmd/todo -complete 1

go run ./cmd/todo -delete 1
```

You can also provide task text through a pipe when using `-add`:
```bash
echo "Buy groceries" | go run ./cmd/todo -add
```

**Available flags:**
- `-add` : add a new todo (supports inline text or stdin via pipe)
- `-complete <number>` : mark the todo at the given index as completed
- `-delete <number>` : delete the todo at the given index
- `-list` : list all todos

**Persistence:**
The todo list is stored in `todos.json` in the current working directory. If the file does not exist or is empty, the application starts with an empty list.

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
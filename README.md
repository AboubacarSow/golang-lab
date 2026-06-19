# 🐹 golang-lab

A hands-on learning repository for experimenting with **Go (Golang)**. This project serves as a structured collection of small, focused mini-projects built to explore core language concepts, standard library usage, concurrency patterns, networking, and backend development practices in Go.

> **Goal:** Learn Go by building real, working things — not just reading docs.

---

## 📁 Repository Structure

```
golang-lab/
├── todocli/           # CLI-based Todo manager (file persistence, CRUD)
├── inventory/         # Terminal inventory manager with TUI and JSON persistence
├── pokedex/           # In-progress PokeAPI exploration CLI with cache and exploration features planned
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

### 2. `inventory` — Terminal Inventory Manager

`inventory` is a terminal user interface application that manages stock items and stores data in `data/inventory.json`.

> **Refactored:** Started as a single `main.go` file, then restructured to apply **Single Responsibility** and **Clean Architecture** principles — separating domain logic, UI, and entry point into distinct packages.

**Package structure:**
```
inventory/
├── cmd/                  # Entry point — wires everything together
├── data/                 # JSON persistence (inventory.json)
└── internal/
    ├── inventory/        # Domain logic — structs, CRUD, file I/O
    └── ui/               # TUI layer — tview forms, tables, event handling
```

**What it supports:**
- Load inventory from `data/inventory.json`
- Display current inventory in a TUI table
- Add a new item with name and stock quantity
- Delete an item by index
- Persist inventory changes back to `data/inventory.json`

**Concepts covered:**
- Clean architecture and package separation (`internal/`)
- Single Responsibility Principle applied to a real refactor
- Terminal UI with `tview`
- Tabular formatting with `simpletable`
- JSON file persistence and path handling
- Slice mutation and index-based CRUD operations
- Error handling and interactive form input

**Run the app:**
```bash
cd inventory
go run ./cmd
```

**Behavior:**
If `data/inventory.json` does not exist, the app starts with an empty inventory and saves updates automatically.

---

### 3. `pokedex` — PokeAPI REPL CLI (in-progress)

`pokedex` is a REPL-based command-line tool for exploring Pokémon data using the public PokeAPI (https://pokeapi.co/).

**What it does:**
- Provides an interactive REPL for listing location areas, exploring which Pokémon appear in an area, attempting to "catch" Pokémon, and viewing a local list of caught Pokémon.
- Uses an internal `pokeapi` client (`internal/pokeapi`) and a simple in-memory cache implementation (`internal/pokecache`).

**Available REPL commands:**
- `help` — show the help menu and available commands
- `map` — list location areas (paginated)
- `mapback` — show the previous page of location areas
- `explore {location_area_name}` — list Pokémon encountered in the given location area
- `catch {pokemon_name}` — attempt to catch a Pokémon by name (simple probability check)
- `pokedex` — list all caught Pokémon
- `inspect {pokemon_name}` — show detailed information for a caught Pokémon
- `exit` — exit the REPL

**Run the app:**
```bash
cd pokedex
go run .
```

Once running, type `help` to see usage. Example REPL session:
```
> map
- pallet-town
- viridian-forest
> explore viridian-forest
Pokeman in :viridian-forest
- pidgey
- rattata
> catch pidgey
pidgey pokeman was catched
> pokedex
Pokedex:
- 1. pidgey
> inspect pidgey
pidgey Information Details:
Height:3
Weight:18
...
```

**Implementation notes & internal layout**
- **Client:** `pokedex/internal/pokeapi` — wraps HTTP calls (`net/http`) to fetch location areas and Pokémon details from PokeAPI; defines JSON structs for unmarshalling responses.
- **Cache:** `pokedex/internal/pokecache` — a concurrent, TTL-based in-memory cache:
  - Uses `sync.RWMutex` to protect concurrent reads/writes to the cache map.
  - Spawns a background goroutine (`reapLoop()`) that runs on a ticker to evict expired entries every `interval`.
  - `Add(val []byte, key string)` stores entries with a timestamp; `Get(key string)` retrieves them.
  - Entries are considered expired if they were created more than `interval` ago.
- **Entry point:** `pokedex/main.go` sets up a `config` struct (default cache TTL = 45s) and starts the interactive REPL in `repl.go`.

**Concepts demonstrated:**
- **Concurrency:** background goroutine using `time.Ticker` for periodic cache reaping.
- **Synchronization:** `sync.RWMutex` ensures safe concurrent access to the cache from multiple goroutines.
- **HTTP networking:** `net/http.Client` with a 1-minute timeout for API calls.
- **TTL caching:** time-based entry expiration and automatic cleanup.

**Known issues / TODOs**
- The project is actively developed; expect rough edges and incomplete error handling.
- **pokecache.go:** There are two versions of the reap loop (`reaploop()` and `reapLoop()`); only one should be active.
- Pagination and richer navigation could be improved (page links are stored in the `config`).

**Tests & exploration**
- There is a small test harness in `pokedex/repl_test.go` to guide behavior-driven checks.

---

## 🧠 Topics This Repo Will Explore

| Area | Status |
|---|---|
| Structs, methods, interfaces | ✅ Started (`todocli`, `inventory`) |
| Error handling patterns | ✅ Started (`todocli`, `inventory`, `pokedex`) |
| File I/O & JSON | ✅ Started (`todocli`, `inventory`) |
| CLI / TUI input handling | ✅ Started (`inventory`) |
| Clean architecture & package design | ✅ Started (`inventory`, `pokedex`) |
| HTTP clients & REST APIs | ✅ Started (`pokedex` via PokeAPI client) |
| Concurrency (`goroutines`, `sync`, `time.Ticker`) | ✅ Started (`pokedex` via pokecache background reaping) |
| In-memory caching & TTL strategies | ✅ Started (`pokedex` via pokecache) |
| CLI argument parsing (`flag`, `cobra`) | 🔜 Planned |
| HTTP servers & routing | 🔜 Planned |
| Channels & select | 🔜 Planned |
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

go run ./cmd/todo -list

# Or switch to the inventory app
cd ../inventory

go run ./cmd

# Run tests (if available)
go test ./...
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
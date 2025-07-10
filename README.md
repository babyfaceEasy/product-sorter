# Product Sorter

A flexible and extensible product catalog sorting engine built in Go 1.24+.  
Designed to allow teams to experiment with different sorting strategies such as sorting by price, sales/view ratio, or custom algorithms — without modifying core logic.

---

## Features

- Built with **Go 1.24+**
- Clean and idiomatic Go architecture
- Extensible with the **Strategy pattern**
- Uses **Dependency Injection** for logger and sorters
- Pluggable logging via a custom logger interface (e.g., uses `zap`)
- Dockerized for local or CI/CD use
- Unit tested

---

## Project Structure

```bash
product-sorter/
├── cmd/
│   └── product-sorter/         # App entrypoint (main.go)
├── internal/
│   ├── logger/                 # Logger interface + zap implementation
│   ├── models/                 # Product data model
│   ├── sorter/                 # Sorter interface and concrete implementations
│   └── utils/                  # Sample product generator
├── go.mod / go.sum            # Go module files
├── Dockerfile                 # For containerized usage
├── .dockerignore / .gitignore
└── README.md
```

---

## Design Patterns & Architecture

### Strategy Pattern
Each sorter (e.g., `ByPrice`, `BySalesViewRatio`) implements a common `Sorter` interface:

```go
type Sorter interface {
    Sort(products []Product) []Product
}
```

This allows dynamic selection of sorting logic at runtime — ideal for A/B testing.

---

### Dependency Injection
Logger and sorters are injected into functions or structs, allowing:

- Mocking in tests
- Swapping logger backends (zap → logrus, or a no-op logger)
- Loose coupling and easier testing

---

### Logger Interface
The logger is abstracted via a simple interface:

```go
type Logger interface {
    Info(msg string, fields ...zap.Field)
    Error(msg string, fields ...zap.Field)
}
```

Implemented using [Uber's zap](https://github.com/uber-go/zap).

---

## Running the App

### Prerequisites
- Go 1.24+
- Docker (optional)

### Run with Go

```bash
go run cmd/product-sorter/main.go
```

### Run with Docker

```bash
docker build -t product-sorter .
docker run --rm product-sorter
```

---

## Run Tests

```bash
go test ./...
```

---

## How to Add a New Sorter (for other teams)

To add a new sorting strategy:

1. **Create a file** in `internal/sorter/`, e.g. `by_created_date.go`
2. **Implement the `Sorter` interface:**

```go
type ByCreatedDate struct {
    logger logger.Logger
}

func (s ByCreatedDate) Sort(products []models.Product) []models.Product {
    // your sort logic
}
```

3. **Register or use it** in your main logic or experiments:

```go
newSorter := sorter.ByCreatedDate{Logger: injectedLogger}
sorted := newSorter.Sort(products)
```

4. (Optional) Add a unit test in `by_created_date_test.go`

---

## Author

Olakunle Michael Odegbaro — [LinkedIn](https://www.linkedin.com/in/olakunle-odegbaro-ab819421)  
Position: Applying for **Senior Golang Engineer**

---

## License

MIT — free to use, share, and extend.

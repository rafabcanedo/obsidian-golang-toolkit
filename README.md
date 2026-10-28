# obsidian-golang-toolkit

A collection of framework-agnostic utility packages for Go projects. No external dependencies beyond the standard library (except `gopkg.in/yaml.v3` for YAML config loading).

**Requires Go 1.22+**

## Installation

```bash
go get github.com/rafabcanedo/obsidian-golang-toolkit
```

Import only the packages you need:

```go
import "github.com/rafabcanedo/obsidian-golang-toolkit/strutil"
import "github.com/rafabcanedo/obsidian-golang-toolkit/validator"
```

## Packages

### Config & Environment

| Package | Description | Functions |
|---------|-------------|-----------|
| `config` | Load configuration from files | `LoadEnv`, `LoadJSON`, `LoadYAML` |
| `env` | Read environment variables with defaults | `GetString` |

### HTTP

| Package | Description | Functions |
|---------|-------------|-----------|
| `httputil` | HTTP response helpers | `WriteJSON`, `WriteError`, `WriteNoContent` |
| `middleware` | HTTP middleware | `Logger`, `Recovery`, `RequestID`, `GetRequestID`, `CORS` |
| `health` | Health check handler | `New`, `AddCheck`, `Handler` |
| `pagination` | Pagination params and metadata | `New`, `Params.Offset`, `Params.NewMeta` |

### Errors & Validation

| Package | Description | Functions |
|---------|-------------|-----------|
| `errutil` | Typed HTTP errors | `NotFound`, `BadRequest`, `Unauthorized` |
| `validator` | Fluent input validation | `New`, `Required`, `MinLen`, `MaxLen`, `Email`, `Validate` |

### Data Structures

| Package | Description | Functions |
|---------|-------------|-----------|
| `sliceutil` | Slice utilities with generics | `Contains`, `Filter`, `Map`, `Unique`, `Chunk` |
| `maputil` | Map utilities with generics | `Keys`, `Values`, `Merge`, `Filter` |

### Strings & Random

| Package | Description | Functions |
|---------|-------------|-----------|
| `strutil` | String manipulation | `IsEmpty`, `Capitalize`, `Truncate`, `CamelToSnake`, `Slugify` |
| `randutil` | Random value generation | `String`, `Int`, `UUID` |

### Context & Logging

| Package | Description | Functions |
|---------|-------------|-----------|
| `contextutil` | Type-safe context values | `WithValue`, `Value[T]` |
| `logutil` | Structured logger via `slog` | `NewLogger` |

### Testing

| Package | Description | Functions |
|---------|-------------|-----------|
| `testutil` | Test assertion helpers | `Equal[T]` |

---

## Roadmap

### Done

- [x] `config` — load env, JSON, YAML files
- [x] `env` — environment variable helpers
- [x] `errutil` — typed HTTP errors
- [x] `httputil` — JSON response helpers
- [x] `middleware` — logger, recovery, request ID, CORS
- [x] `health` — health check HTTP handler
- [x] `pagination` — page params and metadata
- [x] `sliceutil` — generic slice operations
- [x] `maputil` — generic map operations
- [x] `strutil` — string manipulation utilities
- [x] `contextutil` — type-safe context values
- [x] `logutil` — structured logger
- [x] `validator` — fluent input validation
- [x] `randutil` — random string, int, UUID v4
- [x] `testutil` — test helpers

### Planned

- [ ] `mathutil` — generic Min, Max, Clamp, Abs
- [ ] `ptr` — generic pointer helpers (`ToPtr`, `FromPtr`)
- [ ] `timeutil` — BeginOf, EndOf, IsBetween, format duration
- [ ] `retry` — retry with exponential backoff
- [ ] `circuitbreaker` — simple circuit breaker
- [ ] `worker` — goroutine pool
- [ ] `cache` — in-memory cache with TTL

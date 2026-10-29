# obsidian-golang-toolkit

A collection of framework-agnostic utility packages for Go APIs. No external dependencies beyond the standard library (except `gopkg.in/yaml.v3` for YAML config loading).

**Requires Go 1.22+**

## Installation

```bash
go get github.com/rafabcanedo/obsidian-golang-toolkit
```

Import only the packages you need:

```go
import "github.com/rafabcanedo/obsidian-golang-toolkit/api/auth"
import "github.com/rafabcanedo/obsidian-golang-toolkit/util/strutil"
import "github.com/rafabcanedo/obsidian-golang-toolkit/validator"
```

See [`examples/`](./examples) for runnable usage examples.

## Packages

### `api/` — HTTP & API

| Package | Description |
|---------|-------------|
| `api/auth` | JWT generation, parsing, and middleware |
| `api/binding` | Decode JSON body and query params into structs |
| `api/httpclient` | HTTP client with timeout, headers, and retry |
| `api/middleware` | Logger, Recovery, RequestID, CORS |
| `api/ratelimit` | Token bucket rate limiter with middleware |
| `api/respond` | Standardized JSON response envelope |

### `util/` — Generic Utilities

| Package | Description |
|---------|-------------|
| `util/contextutil` | Type-safe context value helpers |
| `util/errutil` | Typed HTTP errors (`NotFound`, `BadRequest`, `Unauthorized`) |
| `util/httputil` | Low-level JSON response helpers |
| `util/logutil` | Structured logger via `slog` |
| `util/maputil` | Generic map operations (`Keys`, `Values`, `Merge`, `Filter`) |
| `util/randutil` | Random string, int, and UUID v4 |
| `util/sliceutil` | Generic slice operations (`Contains`, `Filter`, `Map`, `Unique`, `Chunk`) |
| `util/strutil` | String utilities (`IsEmpty`, `Capitalize`, `Truncate`, `CamelToSnake`, `Slugify`) |

### Root

| Package | Description |
|---------|-------------|
| `config` | Load configuration from `.env`, JSON, and YAML files |
| `env` | Read environment variables with defaults |
| `health` | Health check HTTP handler with named checks |
| `pagination` | Page params, offset calculation, and metadata |
| `testutil` | Test assertion helpers (`Equal`, `NoError`, `Error`, `EqualError`) |
| `validator` | Fluent input validation (`Required`, `MinLen`, `MaxLen`, `Email`) |

---

## Roadmap

### Done

- [x] `api/auth` — JWT generation, parsing, and middleware
- [x] `api/binding` — JSON body and query param binding
- [x] `api/httpclient` — HTTP client with options
- [x] `api/middleware` — logger, recovery, request ID, CORS
- [x] `api/ratelimit` — token bucket rate limiter
- [x] `api/respond` — standardized JSON response envelope
- [x] `config` — load env, JSON, YAML files
- [x] `env` — environment variable helpers
- [x] `health` — health check HTTP handler
- [x] `pagination` — page params and metadata
- [x] `testutil` — test helpers
- [x] `util/contextutil` — type-safe context values
- [x] `util/errutil` — typed HTTP errors
- [x] `util/httputil` — JSON response helpers
- [x] `util/logutil` — structured logger
- [x] `util/maputil` — generic map operations
- [x] `util/randutil` — random string, int, UUID v4
- [x] `util/sliceutil` — generic slice operations
- [x] `util/strutil` — string manipulation utilities
- [x] `validator` — fluent input validation

### Planned

- [ ] `mathutil` — generic Min, Max, Clamp, Abs
- [ ] `ptr` — generic pointer helpers (`ToPtr`, `FromPtr`)
- [ ] `timeutil` — BeginOf, EndOf, IsBetween, format duration
- [ ] `retry` — retry with exponential backoff
- [ ] `circuitbreaker` — simple circuit breaker
- [ ] `worker` — goroutine pool
- [ ] `cache` — in-memory cache with TTL

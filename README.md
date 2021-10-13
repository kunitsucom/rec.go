# rec.go - Simple JSON Logger

[![pkg](https://pkg.go.dev/badge/github.com/rec-logger/rec.go)](https://pkg.go.dev/github.com/rec-logger/rec.go)
[![goreportcard](https://goreportcard.com/badge/github.com/rec-logger/rec.go)](https://goreportcard.com/report/github.com/rec-logger/rec.go)
[![workflow](https://github.com/rec-logger/rec.go/workflows/CI/badge.svg)](https://github.com/rec-logger/rec.go/tree/v1)
[![codecov](https://codecov.io/gh/rec-logger/rec.go/branch/v1/graph/badge.svg?token=LUn3V3D61I)](https://codecov.io/gh/rec-logger/rec.go)
[![sourcegraph](https://sourcegraph.com/github.com/rec-logger/rec.go/-/badge.svg)](https://sourcegraph.com/github.com/rec-logger/rec.go)

rec.go is a lightweight and no dependencies simple JSON (NDJSON, JSONLines, JSONL) logger for Go.  

## Project Goal

To provide a JSON logger with a simple interface, easy setup, and no dependencies.  

## Features

- **Simple and easy-to-use interface**
  - Anyone can start using rec.go by just reading the README for a minute.
- **Lightweight**
  - [Benchmark comparison rec.go with major logger libraries.](https://gist.github.com/newtstat/68941357e83adb7e441e0197e94e21ea)
    - excerpt:
      > 
      > Output of log entries with common JSON log format and 1 `error` field:
      > 
      > | logger  |     time/op |    delta |
      > |:--------|------------:|---------:|
      > | rec.go  | 2.25µs ± 7% |        - |
      > | zap     | 8.24µs ±10% | +266.51% |
      > | zerolog | 7.76µs ± 6% | +245.24% |
      > 
      > | logger  | alloc/op bytes |    delta |
      > |:--------|---------------:|---------:|
      > | rec.go  |      208B ± 0% |        - |
      > | zap     |      392B ± 0% |  +88.65% |
      > | zerolog |      336B ± 0% |  +61.54% |
      > 
      > | logger  | allocs/op times |    delta |
      > |:--------|----------------:|---------:|
      > | rec.go  |       1.00 ± 0% |        - |
      > | zap     |       4.00 ± 0% | +300.00% |
      > | zerolog |       4.00 ± 0% | +300.00% |
      > 
- **No dependencies**
  - Take a look at this clean [`go.mod`](/go.mod).
  - rec.go will never corrupt your `go.mod` and `go.sum`.

## Examples of how to use

### Use default logger ([go.dev/play](https://go.dev/play/p/CzpS8hSUVAe))

```go
package main

import (
    "io"
    "os"
    "time"

    "golang.org/x/xerrors"

    "github.com/rec-logger/rec.go"
)

func main() {
    t := time.Minute
    err := xerrors.Errorf("error: %w", io.EOF)
    wrap := xerrors.Errorf("wrap: %w", err)

    // Use default logger
    rec.L().Info(
        "default logger",
        rec.DurationFormat("duration", t),
        rec.Error(wrap),
        rec.ErrorStacktrace(wrap),
    )
}
```

output:  

```console
$ go run main.go
{"timestamp":"2009-11-10T23:00:00Z","severity":"INFO","caller":"sandbox1448316521/prog.go:17","message":"default logger","duration":"1m0s","error":"wrap: error: EOF","errorStacktrace":"wrap:\n    main.main\n        /tmp/sandbox1448316521/prog.go:15\n  - error:\n    main.main\n        /tmp/sandbox1448316521/prog.go:14\n  - EOF"}
```

### Setup logger using default config ([go.dev/play](https://go.dev/play/p/cYrNtlbnp0_W))

```go
package main

import (
    "io"
    "os"
    "time"

    "golang.org/x/xerrors"

    "github.com/rec-logger/rec.go"
)

func main() {
    // Setup logger using default config
    logger := rec.Must(rec.New(os.Stderr))

    t := time.Minute
    err := xerrors.Errorf("error: %w", io.EOF)
    wrap := xerrors.Errorf("wrap: %w", err)

    logger.Info(
        "logger generated from default config",
        rec.DurationFormat("duration", t),
        rec.Error(wrap),
        rec.ErrorStacktrace(wrap),
    )
}
```

output:  

```console
$ go run main.go
{"timestamp":"2009-11-10T23:00:00Z","severity":"INFO","caller":"sandbox1333751707/prog.go:20","message":"logger generated from default config","duration":"1m0s","error":"wrap: error: EOF","errorStacktrace":"wrap:\n    main.main\n        /tmp/sandbox1333751707/prog.go:18\n  - error:\n    main.main\n        /tmp/sandbox1333751707/prog.go:17\n  - EOF"}
```


### Setup logger using rec.Option ([go.dev/play](https://go.dev/play/p/irGaIhNS5MD))

```go
package main

import (
    "io"
    "os"
    "time"

    "golang.org/x/xerrors"

    "github.com/rec-logger/rec.go"
)

func main() {
    // Setup logger using rec.Option (Functional options)
    logger := rec.Must(rec.New(os.Stderr, rec.WithTimestampFieldKey("time"), rec.WithUseHostnameField(true), rec.WithHostnameFieldKey("host")))

    t := time.Minute
    err := xerrors.Errorf("error: %w", io.EOF)
    wrap := xerrors.Errorf("wrap: %w", err)

    logger.Info(
        "logger generated from rec.Option",
        rec.DurationFormat("duration", t),
        rec.Error(wrap),
        rec.ErrorStacktrace(wrap),
    )
}
```

output:  

```console
$ go run main.go
{"time":"2009-11-10T23:00:00Z","severity":"INFO","host":"f054dfbe1788","caller":"sandbox480936273/prog.go:20","message":"logger generated from rec.Option","duration":"1m0s","error":"wrap: error: EOF","errorStacktrace":"wrap:\n    main.main\n        /tmp/sandbox480936273/prog.go:18\n  - error:\n    main.main\n        /tmp/sandbox480936273/prog.go:17\n  - EOF"}
```

### Setup logger using rec.Config ([go.dev/play](https://go.dev/play/p/4BYLTFiMPTu))

```go
package main

import (
    "io"
    "os"
    "time"

    "golang.org/x/xerrors"

    "github.com/rec-logger/rec.go"
)

func main() {
    config := &rec.Config{
        // "timestamp":"...",
        UseTimestampField:    true,
        TimestampFieldKey:    "timestamp",
        TimestampFieldFormat: time.RFC3339Nano,
        // "severity":"...",
        UseSeverityField:     true,
        SeverityFieldKey:     "severity",
        SeverityThreshold:    rec.DEFAULT,
        UseUppercaseSeverity: true,
        // "hostname":"...",
        UseHostnameField:   true,
        HostnameFieldKey:   "hostname",
        HostnameFieldValue: os.Getenv("HOSTNAME"),
        // "caller":"...",
        UseCallerField: true,
        CallerFieldKey: "caller",
        CallerSkip:     4,
        UseShortCaller: true,
        // "message":"...",
        UseMessageField: true,
        MessageFieldKey: "message",
        // \n
        LineSeparator: "\n",
    }

    // Setup logger using rec.Config
    logger := rec.Must(rec.NewWithConfig(os.Stderr, config))

    t := time.Minute
    err := xerrors.Errorf("error: %w", io.EOF)
    wrap := xerrors.Errorf("wrap: %w", err)

    logger.Info(
        "logger generated from rec.Config",
        rec.DurationFormat("duration", t),
        rec.Error(wrap),
        rec.ErrorStacktrace(wrap),
    )
}
```

output:  

```console
$ go run main.go
{"timestamp":"2009-11-10T23:00:00Z","severity":"INFO","hostname":"acab9130628a","caller":"sandbox2890955676/prog.go:46","message":"logger generated from rec.Config","duration":"1m0s","error":"wrap: error: EOF","errorStacktrace":"wrap:\n    main.main\n        /tmp/sandbox2890955676/prog.go:44\n  - error:\n    main.main\n        /tmp/sandbox2890955676/prog.go:43\n  - EOF"}
```

### Setup buffered logger ([go.dev/play](https://go.dev/play/p/Ph3Iq4SbFAP))

```go
package main

import (
    "bufio"
    "io"
    "os"
    "time"

    "golang.org/x/xerrors"

    "github.com/rec-logger/rec.go"
)

func main() {
	// new *bufio.Writer
	bufStderr := bufio.NewWriter(os.Stderr)
	defer bufStderr.Flush()

	// Setup buffered logger
	logger := rec.Must(rec.New(bufStderr))

    t := time.Minute
    err := xerrors.Errorf("error: %w", io.EOF)
    wrap := xerrors.Errorf("wrap: %w", err)

    logger.Info(
        "buffered logger",
        rec.DurationFormat("duration", t),
        rec.Error(wrap),
        rec.ErrorStacktrace(wrap),
    )
}
```

output:  

```console
$ go run main.go
{"timestamp":"2009-11-10T23:00:00Z","severity":"INFO","caller":"sandbox499319467/prog.go:26","message":"buffered logger","duration":"1m0s","error":"wrap: error: EOF","errorStacktrace":"wrap:\n    main.main\n        /tmp/sandbox499319467/prog.go:24\n  - error:\n    main.main\n        /tmp/sandbox499319467/prog.go:23\n  - EOF"}
```

### Setup logger that context fields added ([go.dev/play](https://go.dev/play/p/Zc4p9fArvnY))

```go
package main

import (
    "bufio"
    "io"
    "os"
    "time"

    "golang.org/x/xerrors"

    "github.com/rec-logger/rec.go"
)

func main() {
	// new *bufio.Writer
	bufStderr := bufio.NewWriter(os.Stderr)
	defer bufStderr.Flush()

	// Setup buffered logger
	logger := rec.Must(rec.New(bufStderr))
    contextLogger := logger.With(rec.String("name", "myLogger"), rec.Int("id", 100))

    t := time.Minute
    err := xerrors.Errorf("error: %w", io.EOF)
    wrap := xerrors.Errorf("wrap: %w", err)

    contextLogger.Info(
        "buffered logger",
        rec.DurationFormat("duration", t),
        rec.Error(wrap),
        rec.ErrorStacktrace(wrap),
    )
}
```

output:  

```console
$ go run main.go
{"timestamp":"2009-11-10T23:00:00Z","severity":"INFO","caller":"sandbox1328679084/prog.go:27","message":"buffered logger","name":"myLogger","id":100,"duration":"1m0s","error":"wrap: error: EOF","errorStacktrace":"wrap:\n    main.main\n        /tmp/sandbox1328679084/prog.go:25\n  - error:\n    main.main\n        /tmp/sandbox1328679084/prog.go:24\n  - EOF"}
```

### Replace the logger in the Go standard log package with rec.Logger, and rollback ([go.dev/play](https://go.dev/play/p/mtvvTnH39zf))

```go
package main

import (
    "log"
    "os"

    "github.com/rec-logger/rec.go"
)

func main() {
    logger := rec.Must(rec.New(os.Stderr))

    // Replace the logger in the Go standard log package with rec.Logger
    rollback := rec.ReplaceStdLogger(logger, rec.INFO)

    log.Println("replaced")

    // To revert back to the original logger, run rollback.
    rollback()

    log.Println("rollback")
}
```

output:  

```console
$ go run main.go
{"timestamp":"2009-11-10T23:00:00Z","severity":"INFO","caller":"sandbox1964866315/prog.go:16","message":"replaced"}
2009/11/10 23:00:00 rollback
```

# logger

**logger** is a logging library for Go with multiple output destinations and semantic logging methods.

## Installation

```sh
go get crdx.org/logger
```

## Usage

```go
import "crdx.org/logger"

func main() {
    logger.InitStderr() // see below for other types

    logger.Info("starting up")
    logger.Added("new device: %s", name)
    logger.Warn("disk usage at %d%%", usage)
    logger.Fatal(err)
}
```

### Initialisation

```go
logger.InitStderr()
logger.InitSyslog()
logger.InitNull()
```

If you need dynamic initialisation, pass a destination directly.

```go
logger.Init(logger.DestinationStderr)
logger.Init(logger.DestinationSyslog)
logger.Init(logger.DestinationNull)
```

## Methods

All methods accept a format string and variadic arguments.

### Status

| Method      | Prefix | Colour |
| ----------- | ------ | ------ |
| `Added`     | `[+]`  | Green  |
| `Removed`   | `[-]`  | Yellow |
| `Succeeded` | `[✓]`  | Green  |
| `Failed`    | `[𐄂]`  | Red    |
| `Enabled`   | `[✓]`  | Green  |
| `Disabled`  | `[𐄂]`  | Red    |

### General

| Method  | Description                                                      |
| ------- | ---------------------------------------------------------------- |
| `Print` | General message.                                                 |
| `Info`  | Info message.                                                    |
| `Warn`  | Warning (yellow, prefixed with "Warning:").                      |
| `Err`   | Error (red, prefixed with "Error:"). Accepts `error` or `string`.|
| `Fatal` | Like `Err`, but calls `os.Exit(1)` after logging.                |

## Destinations

Three output destinations are available.

### Stderr

Logs to standard error with colour-coded output. Status methods are prefixed with relevant status symbols.

### Syslog

Logs to the system log with the program name as the tag. Messages are routed to info or error level as appropriate.

### Null

Silently discards all log output. Useful for testing.

## Contributions

Open an [issue](https://github.com/crdx/logger/issues) or send a [pull request](https://github.com/crdx/logger/pulls).

## Licence

[GPLv3](LICENCE).

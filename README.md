# traceerr

`traceerr` is a lightweight Go library that enhances error handling and logging by providing:

1. **Stack Traces for Errors**: Automatically captures stack traces when errors occur.
2. **Log Trace Inclusion**: Includes the last N logs (configurable) in the error trace for better debugging context.
3. **Concise Logging**: Provides simplified logging methods (`Print`, `Printf`, etc.), similar to Go's standard `log` package.
4. **Real-Time Debug Mode**: Optionally outputs logs to `stdout` when the environment variable `DEBUG=1` is set.

---

## Features

- **Stack trace**: Every error created with `traceerr.Error` or `traceerr.Errorf` automatically includes a detailed stack trace.
- **Log trace**: Logs leading up to the error are included for better debugging context.
- **Customizable log storage**: The number of logs stored in memory can be configured using the `LOG_PERSIST_TRACE` environment variable.
- **Standard log functions**: Provides methods like `Print`, `Printf`, `Println`, `Fatal`, `Fatalf`, `Panic`, and `Panicf` for ease of use.

---

## Installation

```bash
go get github.com/devalexandre/traceerr
```

---
# Usage

## Setup

1- Set environment variables (optional):

    LOG_PERSIST_TRACE: Defines the number of logs to store in memory (default: 20).
    DEBUG: If set to 1, logs are also printed to stdout in real-time.
```bash
export LOG_PERSIST_TRACE=50  # Store up to 50 logs
export DEBUG=1               # Enable real-time logging to stdout
```

2- Import the package:

```go
import "github.com/devalexandre/traceerr"
```

## Examples
### Logging and Error Handling

```go
package main

import (
	"github.com/devalexandre/traceerr"
)

func funcA() error {
	traceerr.Print("Starting funcA")
	traceerr.Printf("Processing step %d in funcA", 1)
	return traceerr.Errorf("Something went wrong in funcA")
}

func funcB() error {
	traceerr.Print("Starting funcB")
	if err := funcA(); err != nil {
		return traceerr.Errorf("funcB encountered an error: %v", err)
	}
	return nil
}

func main() {
	err := funcB()
	if err != nil {
		traceerr.Errorln("Program completed with error:", err)
	} else {
		traceerr.Print("Program completed successfully")
	}
}

```

### Output
If funcA encounters an error, the output will include:

1- Log Trace: All logs leading up to the error.
2- Error Message: The error message provided in Error or Errorf.
3- Stack Trace: A detailed trace of the function calls leading to the error.

```bash
01-12-2025 21:19:26 - Starting funcB
01-12-2025 21:19:26 - Starting funcA
01-12-2025 21:19:26 - Processing step 1 in funcA
---- LOGS ----
01-12-2025 21:19:26 - Starting funcB
01-12-2025 21:19:26 - Starting funcA
01-12-2025 21:19:26 - Processing step 1 in funcA
---- ERROR ----
Error: funcB encountered an error: ---- LOGS ----
01-12-2025 21:19:26 - Starting funcB
01-12-2025 21:19:26 - Starting funcA
01-12-2025 21:19:26 - Processing step 1 in funcA
---- ERROR ----
Error: Something went wrong in funcA
Stack trace:
	main.funcA
		/home/alexandre/projects/devalexandre/traceerr/example/main.go:15
	main.funcB
		/home/alexandre/projects/devalexandre/traceerr/example/main.go:23
	main.main
		/home/alexandre/projects/devalexandre/traceerr/example/main.go:38
	runtime.main
		/home/alexandre/.gvm/gos/go1.23/src/runtime/proc.go:272
	runtime.goexit
		/home/alexandre/.gvm/gos/go1.23/src/runtime/asm_amd64.s:1700

Stack trace:
	github.com/devalexandre/traceerr.Errorf
		/home/alexandre/projects/devalexandre/traceerr/traceerr.go:46
	main.funcB
		/home/alexandre/projects/devalexandre/traceerr/example/main.go:24
	main.main
		/home/alexandre/projects/devalexandre/traceerr/example/main.go:38
	runtime.main
		/home/alexandre/.gvm/gos/go1.23/src/runtime/proc.go:272
	runtime.goexit
		/home/alexandre/.gvm/gos/go1.23/src/runtime/asm_amd64.s:1700


```

### Real-Time Debugging
If DEBUG=1 is set, the logs (Starting funcA, Starting funcB, etc.) will also appear in the console as they are generated.

## API Reference

### Logging
- `traceerr.Print(v ...interface{})`: Logs a message.
- `traceerr.Printf(format string, v ...interface{})`: Logs a formatted message.
- `traceerr.Println(v ...interface{})`: Logs a message with a newline.
- `traceerr.Fatal(v ...interface{})`: Logs a message and exits the program.
- `traceerr.Fatalf(format string, v ...interface{})`: Logs a formatted message and exits the program.
- `traceerr.Panic(v ...interface{})`: Logs a message and panics.
- `traceerr.Panicf(format string, v ...interface{})`: Logs a formatted message and panics.

### Error Handling
- `traceerr.Error(msg string) error`: Creates an error with a stack trace and logs trace.
- `traceerr.Errorf(format string, v ...interface{}) error`: Creates a formatted error with a stack trace and logs trace.

### Environment Variables
| Variable          | Description                             | Default |
|-------------------|-----------------------------------------|---------|
| LOG_PERSIST_TRACE | Number of logs to store in memory.      | 20      |
| DEBUG             | If set to 1, logs are printed to stdout.| 0       |

## License

This project is licensed under the MIT License.

## Contribution

Feel free to open issues or pull requests to contribute to the library.

Happy debugging! 🚀

# YukiLog

A simple, colorful structured logging wrapper around Go's slog package. YukiLog provides an easy-to-use logging interface with support for multiple log levels, color-coded output, and configurable formats.

## Features

- Simple API with Info, Debug, Error, Warn, and Fatal log levels
- Color-coded output for better readability in terminal environments
- Support for both text and JSON output formats
- Environment variable configuration
- Programmatic configuration via Config struct
- Built on top of Go's standard slog package

## Installation

```bash
go get github.com/yourusername/yukilog
```

## Quick Start

```go
package main

import (
	"github.com/yourusername/yukilog"
)

func main() {
	// Initialize with environment variables
	yukilog.Initialize(nil)
	
	// Basic usage
	yukilog.Info("Hello World")
	yukilog.Debug("Debug message")
	yukilog.Error("Error message")
	
	// With structured fields
	yukilog.Info("User logged in", "user_id", 123, "username", "johndoe")
}
```

## Configuration

### Environment Variables

- `LOG_LEVEL`: Sets the minimum log level (debug, info, warn, error, fatal)
- `LOG_FORMAT`: Determines output format (json or text)

### Programmatic Configuration

```go
config := yukilog.Config{
	Level:  "debug",
	Format: "json",
	Output: os.Stdout,
}
yukilog.Initialize(&config)
```

## Log Levels

- **Debug**: Cyan - Detailed information for debugging
- **Info**: Green - General information messages
- **Warn**: Yellow - Warning messages
- **Error**: Red - Error messages
- **Fatal**: High-intensity Red - Critical errors that cause program termination

## License

[MIT](LICENSE)

// Package yukilog provides a simple, colorful logging library for Go applications.
// It supports multiple log levels with color-coded output and configurable formats.
package yukilog

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

// Log levels
const (
	LevelDebug = "DEBUG"
	LevelInfo  = "INFO"
	LevelWarn  = "WARN"
	LevelError = "ERROR"
	LevelFatal = "FATAL"
)

// Config holds the logger configuration options
type Config struct {
	// Level sets the minimum log level (debug, info, warn, error, fatal)
	Level string
	// Format determines the output format (json or text)
	Format string
	// Output specifies where logs are written (defaults to os.Stdout)
	Output io.Writer
	// DisableColors disables colored output
	DisableColors bool
	// DisableTimestamp disables timestamp in logs
	DisableTimestamp bool
}

// DefaultConfig returns a Config with default values
func DefaultConfig() Config {
	return Config{
		Level:            "info",
		Format:           "text",
		Output:           os.Stdout,
		DisableColors:    false,
		DisableTimestamp: false,
	}
}

// Logger is the main logger struct
type Logger struct {
	cfg    Config
	output io.Writer
	mutex  sync.Mutex
}

// default logger instance
var defaultLogger *Logger
var loggerMutex sync.Mutex

// Initialize sets up the logger with the given configuration
// If config is nil, environment variables will be used for configuration
func Initialize(config *Config) {
	loggerMutex.Lock()
	defer loggerMutex.Unlock()
	
	var cfg Config
	if config != nil {
		cfg = *config
	} else {
		// Use environment variables
		cfg = DefaultConfig()
		if envLevel := os.Getenv("LOG_LEVEL"); envLevel != "" {
			cfg.Level = envLevel
		}
		if envFormat := os.Getenv("LOG_FORMAT"); envFormat != "" {
			cfg.Format = envFormat
		}
	}
	
	// Create logger
	defaultLogger = NewLogger(cfg)
	
	// Print initialization message
	printInitMessage(cfg.Level)
}

// NewLogger creates a new Logger instance with the given configuration
func NewLogger(cfg Config) *Logger {
	output := cfg.Output
	if output == nil {
		output = os.Stdout
	}

	return &Logger{
		cfg:    cfg,
		output: output,
		mutex:  sync.Mutex{},
	}
}

// printInitMessage prints an initialization message with the log level
func printInitMessage(level string) {
	levelStr := strings.ToLower(level)
	var c *color.Color

	switch levelStr {
	case "debug":
		c = color.New(color.FgCyan)
	case "info":
		c = color.New(color.FgGreen)
	case "warn":
		c = color.New(color.FgYellow)
	case "error":
		c = color.New(color.FgRed)
	case "fatal":
		c = color.New(color.FgHiRed)
	default:
		c = color.New(color.FgGreen)
		levelStr = "info"
	}

	c.Printf("YukiLog initialized with level: %s\n", strings.Title(levelStr))
}

// log logs a message with the given level
func (l *Logger) log(level, msg string, args ...interface{}) {
	// Lock to prevent concurrent writes from garbling output
	l.mutex.Lock()
	defer l.mutex.Unlock()
	
	// Check if level is enabled
	if !isLevelEnabled(l.cfg.Level, level) {
		return
	}

	// Format timestamp
	timeStr := ""
	if !l.cfg.DisableTimestamp {
		timeStr = time.Now().Format(time.RFC3339)
	}

	// Format level with color if enabled
	levelStr := level
	if !l.cfg.DisableColors {
		switch strings.ToUpper(level) {
		case LevelDebug:
			levelStr = color.CyanString(level)
		case LevelInfo:
			levelStr = color.GreenString(level)
		case LevelWarn:
			levelStr = color.YellowString(level)
		case LevelError:
			levelStr = color.RedString(level)
		case LevelFatal:
			levelStr = color.HiRedString(level)
		}
	}

	// Process message and structured fields
	logMsg := msg
	structuredFields := make(map[string]interface{})
	
	if len(args) > 0 && len(args)%2 == 0 {
		for i := 0; i < len(args); i += 2 {
			key, ok := args[i].(string)
			if !ok {
				key = fmt.Sprintf("%v", args[i])
			}
			value := args[i+1]
			
			// Add to text log message
			logMsg += fmt.Sprintf(" %s=%v", key, value)
			
			// Store for JSON
			structuredFields[key] = value
		}
	}
	
	// Output log message
	if l.cfg.Format == "json" {
		// Build JSON manually for simplicity
		jsonMsg := fmt.Sprintf("{\"time\":\"%s\",\"level\":\"%s\",\"msg\":\"%s\"", 
			timeStr, strings.ToUpper(level), msg)
		
		// Add structured fields if any
		if len(structuredFields) > 0 {
			for k, v := range structuredFields {
				jsonMsg += fmt.Sprintf(",\"%s\":\"%v\"", k, v)
			}
		}
		
		// Close JSON object
		jsonMsg += "}"
		
		fmt.Fprintln(l.output, jsonMsg)
	} else {
		// Text format
		if timeStr != "" {
			fmt.Fprintf(l.output, "%s %s %s\n", timeStr, levelStr, logMsg)
		} else {
			fmt.Fprintf(l.output, "%s %s\n", levelStr, logMsg)
		}
	}

	// Exit on fatal
	if strings.ToUpper(level) == LevelFatal {
		os.Exit(1)
	}
}

// isLevelEnabled checks if the given level is enabled based on the configured level
func isLevelEnabled(configLevel, msgLevel string) bool {
	levels := map[string]int{
		"debug": 0,
		"info":  1,
		"warn":  2,
		"error": 3,
		"fatal": 4,
	}

	configLevelVal, ok := levels[strings.ToLower(configLevel)]
	if !ok {
		configLevelVal = levels["info"] // Default to info
	}

	msgLevelVal, ok := levels[strings.ToLower(msgLevel)]
	if !ok {
		return true // If unknown level, allow it
	}

	return msgLevelVal >= configLevelVal
}

// Info logs a message at Info level
func Info(msg string, args ...interface{}) {
	loggerMutex.Lock()
	logger := defaultLogger
	if logger == nil {
		Initialize(nil)
		logger = defaultLogger
	}
	loggerMutex.Unlock()
	
	logger.log(LevelInfo, msg, args...)
}

// Debug logs a message at Debug level
func Debug(msg string, args ...interface{}) {
	loggerMutex.Lock()
	logger := defaultLogger
	if logger == nil {
		Initialize(nil)
		logger = defaultLogger
	}
	loggerMutex.Unlock()
	
	logger.log(LevelDebug, msg, args...)
}

// Error logs a message at Error level
func Error(msg string, args ...interface{}) {
	loggerMutex.Lock()
	logger := defaultLogger
	if logger == nil {
		Initialize(nil)
		logger = defaultLogger
	}
	loggerMutex.Unlock()
	
	logger.log(LevelError, msg, args...)
}

// Warn logs a message at Warn level
func Warn(msg string, args ...interface{}) {
	loggerMutex.Lock()
	logger := defaultLogger
	if logger == nil {
		Initialize(nil)
		logger = defaultLogger
	}
	loggerMutex.Unlock()
	
	logger.log(LevelWarn, msg, args...)
}

// Fatal logs a message at Fatal level and exits the program
func Fatal(msg string, args ...interface{}) {
	loggerMutex.Lock()
	logger := defaultLogger
	if logger == nil {
		Initialize(nil)
		logger = defaultLogger
	}
	loggerMutex.Unlock()
	
	logger.log(LevelFatal, msg, args...)
}

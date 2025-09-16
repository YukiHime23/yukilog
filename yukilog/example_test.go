package yukilog_test

import (
	"os"

	"yukilog/yukilog"
)

// This example demonstrates basic usage of the yukilog package
func Example() {
	// Initialize with default configuration
	yukilog.Initialize(nil)

	// Log messages at different levels
	yukilog.Debug("This is a debug message")
	yukilog.Info("This is an info message")
	yukilog.Warn("This is a warning message")
	yukilog.Error("This is an error message")
	// yukilog.Fatal("This is a fatal message") // This would exit the program
}

// This example demonstrates using structured logging
func Example_structuredLogging() {
	// Initialize with default configuration
	yukilog.Initialize(nil)

	// Log with structured fields
	yukilog.Info("User logged in", 
		"user_id", 123,
		"username", "johndoe",
		"role", "admin",
	)

	// Log with nested structured fields
	yukilog.Info("API request completed",
		"request", map[string]interface{}{
			"method": "GET",
			"path":   "/api/users",
			"query":  "limit=10",
		},
		"response", map[string]interface{}{
			"status": 200,
			"time":   "120ms",
		},
	)
}

// This example demonstrates custom configuration
func Example_customConfiguration() {
	// Create a custom configuration
	config := yukilog.Config{
		Level:           "debug",
		Format:          "json",
		Output:          os.Stdout,
		DisableColors:   true,
		DisableTimestamp: false,
	}

	// Initialize with custom configuration
	yukilog.Initialize(&config)

	// Log messages
	yukilog.Debug("Debug message will be shown")
	yukilog.Info("Info message with JSON formatting")
}

// This example demonstrates creating a custom logger instance
func Example_customLogger() {
	// Create a custom configuration
	config := yukilog.Config{
		Level:  "info",
		Format: "text",
		Output: os.Stderr, // Log to stderr instead of stdout
	}

	// Create a new logger instance
	logger := yukilog.NewLogger(config)
	
	// Use the logger (note: this doesn't affect the global default logger)
	_ = logger
	
	// To make this the default logger:
	// slog.SetDefault(logger.Logger())
}

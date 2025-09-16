package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"yukilog/yukilog"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Load environment variables from .env file
	fmt.Println("=== YukiLog Demo ===")
	fmt.Println()
	
	// Initialize with environment variables
	fmt.Println("1. Initializing with environment variables:")
	yukilog.Initialize(nil)
	
	// Basic usage
	fmt.Println("\n2. Basic usage:")
	yukilog.Info("This is an info message")
	yukilog.Debug("This is a debug message")
	yukilog.Error("This is an error message")
	yukilog.Warn("This is a warning message")
	
	// Using with structured fields
	fmt.Println("\n3. Structured logging:")
	yukilog.Info("User logged in", "user_id", 123, "username", "johndoe")
	
	// Example with JSON format
	fmt.Println("\n4. JSON format:")
	jsonConfig := yukilog.Config{
		Level:           "debug",
		Format:          "json",
		Output:          os.Stdout,
		DisableColors:   false,
		DisableTimestamp: false,
	}
	yukilog.Initialize(&jsonConfig)
	yukilog.Info("This is a JSON formatted log")
	
	// Example with custom configuration (colored text)
	fmt.Println("\n5. Custom colored text format:")
	coloredConfig := yukilog.Config{
		Level:           "debug",
		Format:          "text",
		Output:          os.Stdout,
		DisableColors:   false,
		DisableTimestamp: false,
	}
	yukilog.Initialize(&coloredConfig)
	yukilog.Debug("Debug message with colors")
	yukilog.Info("Info message with colors")
	yukilog.Warn("Warning message with colors")
	yukilog.Error("Error message with colors")
	
	// Compare with standard logger
	fmt.Println("\n6. Standard logger comparison:")
	log.Println("Hello from standard logger")
}

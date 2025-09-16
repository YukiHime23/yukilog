package yukilog_test

import (
	"bytes"
	"strings"
	"testing"

	"yukilog/yukilog"
)

func TestLogLevels(t *testing.T) {
	// Create a buffer to capture log output
	var buf bytes.Buffer

	// Configure logger to write to our buffer
	config := yukilog.Config{
		Level:           "debug",
		Format:          "text",
		Output:          &buf,
		DisableColors:   true,
		DisableTimestamp: false,
	}

	// Initialize the logger
	yukilog.Initialize(&config)

	// Test each log level
	yukilog.Debug("debug message")
	if !strings.Contains(buf.String(), "DEBUG") {
		t.Error("Debug log level not working correctly")
	}
	buf.Reset()

	yukilog.Info("info message")
	if !strings.Contains(buf.String(), "INFO") {
		t.Error("Info log level not working correctly")
	}
	buf.Reset()

	yukilog.Warn("warn message")
	if !strings.Contains(buf.String(), "WARN") {
		t.Error("Warn log level not working correctly")
	}
	buf.Reset()

	yukilog.Error("error message")
	if !strings.Contains(buf.String(), "ERROR") {
		t.Error("Error log level not working correctly")
	}
	buf.Reset()
}

func TestStructuredLogging(t *testing.T) {
	// Create a buffer to capture log output
	var buf bytes.Buffer

	// Configure logger to write to our buffer
	config := yukilog.Config{
		Level:           "info",
		Format:          "text",
		Output:          &buf,
		DisableColors:   true,
		DisableTimestamp: false,
	}

	// Initialize the logger
	yukilog.Initialize(&config)

	// Test structured logging
	yukilog.Info("user login", "user_id", 123, "username", "testuser")
	output := buf.String()

	// Check that the structured fields are included
	if !strings.Contains(output, "user_id=123") {
		t.Error("Structured logging not including numeric field correctly")
	}

	if !strings.Contains(output, "username=testuser") {
		t.Error("Structured logging not including string field correctly")
	}
}

func TestJSONFormat(t *testing.T) {
	// Create a buffer to capture log output
	var buf bytes.Buffer

	// Configure logger to use JSON format
	config := yukilog.Config{
		Level:           "info",
		Format:          "json",
		Output:          &buf,
		DisableColors:   true,
		DisableTimestamp: false,
	}

	// Initialize the logger
	yukilog.Initialize(&config)

	// Test JSON output
	yukilog.Info("test message")
	output := buf.String()

	// Check that the output is in JSON format
	if !strings.Contains(output, "{") || !strings.Contains(output, "}") {
		t.Error("JSON format not working correctly")
	}

	if !strings.Contains(output, "\"level\":\"INFO\"") {
		t.Error("JSON format not including level correctly")
	}

	if !strings.Contains(output, "\"msg\":\"test message\"") {
		t.Error("JSON format not including message correctly")
	}
}

func TestNewLogger(t *testing.T) {
	// Create a buffer to capture log output
	var buf bytes.Buffer

	// Create a new logger instance
	config := yukilog.Config{
		Level:           "debug",
		Format:          "text",
		Output:          &buf,
		DisableColors:   true,
		DisableTimestamp: false,
	}

	logger := yukilog.NewLogger(config)
	if logger == nil {
		t.Fatal("NewLogger returned nil")
	}
}

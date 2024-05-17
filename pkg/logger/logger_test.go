package logger

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	buffer := new(bytes.Buffer)
	logger := NewLogger(buffer, "", log.LstdFlags)

	// Test functionA
	functionA(logger)
	outputA := buffer.String()
	fmt.Println("Output from functionA tests:") // Print the buffer content
	fmt.Println(outputA)                        // Print the captured log output
	expectedOutputs := []string{
		"Logging from functionA",
		"Logging from functionB",
		"Logging from functionC",
		"Log with skip=0",
		"Log with skip=1",
		"Log with skip=2",
	}

	for _, msg := range expectedOutputs {
		if !strings.Contains(outputA, msg) {
			t.Errorf("Expected log to contain %q, got %q", msg, outputA)
		}
	}

	// Clear buffer
	buffer.Reset()

	// Test WithCaller
	logger.WithCaller(1).Info("Test WithCaller")
	outputCaller := buffer.String()
	expectedCaller := "logger_test.go" // assuming this test file is named logger_test.go

	if !strings.Contains(outputCaller, expectedCaller) {
		t.Errorf("Expected log to contain caller info, got %q", outputCaller)
	}
}

// Implementations of functionA, functionB, and functionC with the modified Logger method usage
func functionA(logger *Logger) {
	logger.Info("Logging from functionA")
	functionB(logger)
}

func functionB(logger *Logger) {
	logger.Info("Logging from functionB")
	functionC(logger)
}

func functionC(logger *Logger) {
	logger.Info("Logging from functionC")

	// Log with different skip values
	logger.WithCaller(0).Info("Log with skip=0")
	logger.WithCaller(1).Info("Log with skip=1")
	logger.WithCaller(2).Info("Log with skip=2")
}

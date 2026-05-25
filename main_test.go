package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// Helper function to capture stdout during execution
func captureOutput(args []string, runFunc func()) string {
	oldArgs := os.Args
	oldStdout := os.Stdout
	defer func() {
		os.Args = oldArgs
		os.Stdout = oldStdout
	}()

	os.Args = args
	r, w, _ := os.Pipe()
	os.Stdout = w

	runFunc()

	w.Close()
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	return buf.String()
}

func TestArgumentParsingAndValidation(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedOutput string
	}{
		{
			name:           "Too few arguments",
			args:           []string{"cmd"},
			expectedOutput: "Usage: go run . <text> <banner>",
		},
		{
			name:           "Missing banner defaults to standard",
			args:           []string{"cmd", "hello"},
			expectedOutput: "Usage: go run . <text> <banner>",
		},
		{
			name:           "Invalid banner name triggers error message",
			args:           []string{"cmd", "hello", "invalid_banner"},
			expectedOutput: "Error: Invalid or missing banner. Only standard.txt, shadow.txt, and thinkertoy.txt are allowed. Using standard.txt instead.",
		},
		{
			name:           "Ungrouped multi-word text with valid banner",
			args:           []string{"cmd", "hello", "world", "standard.txt"},
			expectedOutput: "Ungrouped, multi word text found. Printing only first word",
		},
		{
			name:           "Ungrouped multi-word text with invalid banner",
			args:           []string{"cmd", "hello", "world", "invalid_banner"},
			expectedOutput: "Error: Invalid or missing banner", // Checks for partial match of the error sequence
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Wrap main logic in an anonymous function to pass to the capturer
			output := captureOutput(tt.args, func() {
				// We call main directly to test its entry behavior
				// Note: if readbanner() fails during tests because files are missing,
				// it will print "error reading file:" which we can also assert.
				main()
			})

			if !strings.Contains(output, tt.expectedOutput) {
				t.Errorf("Expected output to contain %q, but got %q", tt.expectedOutput, output)
			}
		})
	}
}

func TestEmptyInputExitsQuietly(t *testing.T) {
	args := []string{"cmd", "", "standard.txt"}
	output := captureOutput(args, main)

	if output != "" {
		t.Errorf("Expected no output for empty input string, but got %q", output)
	}
}

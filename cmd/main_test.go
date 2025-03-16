package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMainOutput(t *testing.T) {
	// Save the original stdout
	originalStdout := os.Stdout

	// Create a pipe to capture stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the main function
	main()

	// Close the writer and restore stdout
	w.Close()
	os.Stdout = originalStdout

	// Read the captured output
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	// Check if the output is as expected
	expectedOutput := "Hello World\n"
	if output != expectedOutput {
		t.Errorf("Expected output %q but got %q", expectedOutput, output)
	}
}

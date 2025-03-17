package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	// Capture the output of the main function
	output := captureOutput(func() {
		main()
	})

	expected := "Hello World\n"
	if output != expected {
		t.Errorf("Expected %q but got %q", expected, output)
	}
}

// Helper function to capture output
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

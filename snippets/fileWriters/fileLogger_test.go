package filewriters_test

import (
	"errors"
	"os"
	"testing"

	fileWriters "github.com/stefanovazzocell/GoDojo/snippets/fileWriters"
)

func TestLogFile(t *testing.T) {
	// Setup
	testFile, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	testFile = testFile + string(os.PathSeparator) + "testlogfile"
	t.Logf("Using %q as test file", testFile)
	_ = os.Remove(testFile)
	// Open
	lf, err := fileWriters.NewLogFile(testFile)
	if err != nil {
		t.Fatalf("Error on NewLogFile: %v", err)
	}
	// Write 2 entries
	if err := lf.WriteLine([]byte("Hello")); err != nil {
		t.Fatalf("Failed to write the first entry: %v", err)
	}
	if err := lf.WriteLine([]byte("World")); err != nil {
		t.Fatalf("Failed to write the second entry: %v", err)
	}
	// Close
	if err := lf.Close(); err != nil {
		t.Fatalf("Failed to close file: %v", err)
	}
	// Write after close
	if err := lf.WriteLine([]byte("Write after close")); !errors.Is(err, os.ErrClosed) {
		t.Fatalf("Failed to error on write after closing file: %v", err)
	}
	// Re-Open
	lf, err = fileWriters.NewLogFile(testFile)
	if err != nil {
		t.Fatalf("Error on NewLogFile (re-opening): %v", err)
	}
	// Write 1 entry
	if err := lf.WriteLine([]byte("!")); err != nil {
		t.Fatalf("Failed to write the third entry (post re-opening): %v", err)
	}
	// Close
	if err := lf.Close(); err != nil {
		t.Fatalf("Failed to close (re-opened) file: %v", err)
	}
	// Cleanup
	if err := os.Remove(testFile); err != nil {
		t.Fatalf("Failed to cleanup: %v", err)
	}
}

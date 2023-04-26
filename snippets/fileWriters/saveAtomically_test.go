package filewriters_test

import (
	"os"
	"testing"

	fileWriters "github.com/stefanovazzocell/GoDojo/snippets/fileWriters"
)

func TestSaveToFile(t *testing.T) {
	// Setup
	testFile, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	testFile = testFile + string(os.PathSeparator) + "testfile"
	t.Logf("Using %q as test file", testFile)
	_ = os.Remove(testFile)
	_ = fileWriters.CleanupTempFile(testFile)
	// Attempt to create a file
	if err := fileWriters.SaveToFile(testFile, []byte("Hello World")); err != nil {
		t.Fatalf("Failed on first write: %v", err)
	}
	// Attempt to update file
	if err := fileWriters.SaveToFile(testFile, []byte("Hello World, again")); err != nil {
		t.Fatalf("Failed on second write: %v", err)
	}
	// Cleanup
	if err := os.Remove(testFile); err != nil {
		t.Fatalf("Failed to cleanup: %v", err)
	}
}

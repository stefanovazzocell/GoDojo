package filewriters_test

import (
	"os"
	"testing"
)

// Get a test file
func getTestFileHelper() (string, error) {
	testFile, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return testFile + string(os.PathSeparator) + "testfile", nil
}

// Cleanup a test file
func cleanupTestFileHelper(testFile string) {
	_ = os.Remove(testFile)
}

func BenchmarkFileSync(b *testing.B) {
	testFile, err := getTestFileHelper()
	if err != nil {
		b.Fatalf("Failed to get test file: %v", err)
	}
	data := []byte("HelloWorld_")

	b.Run("FlagAndCall", func(b *testing.B) {
		// Setup
		cleanupTestFileHelper(testFile)
		fp, err := os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0600)
		if err != nil {
			b.Fatalf("Failed to create file: %v", err)
		}
		b.ResetTimer()
		// Test
		for i := 0; i < b.N; i++ {
			if _, err := fp.Write(append(data, byte(i))); err != nil {
				fp.Close()
				b.Fatalf("Failed write %d: %v", i, err)
			}
			if err := fp.Sync(); err != nil {
				fp.Close()
				b.Fatalf("Failed sync %d: %v", i, err)
			}
		}
		// Cleanup
		b.StopTimer()
		if err := fp.Close(); err != nil {
			b.Fatalf("Failed to close file: %v", err)
		}
		cleanupTestFileHelper(testFile)
	})
	b.Run("Flag", func(b *testing.B) {
		// Setup
		cleanupTestFileHelper(testFile)
		fp, err := os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0600)
		if err != nil {
			b.Fatalf("Failed to create file: %v", err)
		}
		b.ResetTimer()
		// Test
		for i := 0; i < b.N; i++ {
			if _, err := fp.Write(append(data, byte(i))); err != nil {
				fp.Close()
				b.Fatalf("Failed write %d: %v", i, err)
			}
		}
		// Cleanup
		b.StopTimer()
		if err := fp.Close(); err != nil {
			b.Fatalf("Failed to close file: %v", err)
		}
		cleanupTestFileHelper(testFile)
	})
	b.Run("Call", func(b *testing.B) {
		// Setup
		cleanupTestFileHelper(testFile)
		fp, err := os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			b.Fatalf("Failed to create file: %v", err)
		}
		b.ResetTimer()
		// Test
		for i := 0; i < b.N; i++ {
			if _, err := fp.Write(append(data, byte(i))); err != nil {
				fp.Close()
				b.Fatalf("Failed write %d: %v", i, err)
			}
			if err := fp.Sync(); err != nil {
				fp.Close()
				b.Fatalf("Failed sync %d: %v", i, err)
			}
		}
		// Cleanup
		b.StopTimer()
		if err := fp.Close(); err != nil {
			b.Fatalf("Failed to close file: %v", err)
		}
		cleanupTestFileHelper(testFile)
	})
}

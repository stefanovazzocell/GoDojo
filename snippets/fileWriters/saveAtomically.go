package filewriters

import "os"

// Save some data to a file atomically
func SaveToFile(path string, data []byte) error {
	tmpFilePath := path + ".tmp"
	// Create tmp file
	fp, err := os.OpenFile(tmpFilePath, os.O_WRONLY|os.O_CREATE|os.O_EXCL|os.O_SYNC, 0600)
	if err != nil {
		return err
	}
	// Write and force-sync data
	if _, err := fp.Write(data); err != nil {
		fp.Close()
		return err
	}
	if err := fp.Sync(); err != nil {
		fp.Close()
		return err
	}
	// Attempt to close file
	if err := fp.Close(); err != nil {
		return err
	}
	// Atomically replace original
	return os.Rename(tmpFilePath, path)
}

// Cleanups any temporary file from SaveToFile in case of errors
func CleanupTempFile(path string) error {
	return os.Remove(path + ".tmp")
}

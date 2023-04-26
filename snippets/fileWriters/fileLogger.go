package filewriters

import "os"

// A log file
type logFile struct {
	fp *os.File
}

// Open/Create log file
func NewLogFile(path string) (logFile, error) {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	return logFile{
		fp: fp,
	}, err
}

// Write a line in the log file
func (lf logFile) WriteLine(data []byte) error {
	_, err := lf.fp.Write(append(data, '\n'))
	if err != nil {
		return err
	}
	return lf.fp.Sync()
}

// Close the log file
func (lf logFile) Close() error {
	return lf.fp.Close()
}

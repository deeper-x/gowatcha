package lwatch

import (
	"fmt"
	"os"
	"strings"
)

// WLog file object
type WLog struct {
	F *os.File
}

// New return new wlog instance
func New(filePath string) (*WLog, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return &WLog{}, err
	}

	return &WLog{F: file}, nil
}

// GetFileSize return file size
func (tr *WLog) GetFileSize() (int64, error) {
	s, err := tr.F.Stat()
	if err != nil {
		return 0, err
	}

	return s.Size(), nil
}

// GetTail get the last part, starting from offset
func (tr *WLog) GetTail(start int64) (int64, error) {
	// read file from offset
	r, err := tr.F.Seek(start, 1)
	if err != nil {
		return 0, err
	}

	return r, nil

}

// GetAppendedString return appended string since last read
func (tr *WLog) GetAppendedString(s, r int64) (string, error) {
	b := make([]byte, s)

	n, err := tr.F.Read(b)
	if err != nil {
		return "", err
	}

	// read appended string from file
	res := fmt.Sprintf("%v", string(b[:n]))

	return res, nil
}

// NeedleExists search for string occurrence
func (tr *WLog) NeedleExists(needle, haystack string) bool {
	return strings.Contains(haystack, needle)
}

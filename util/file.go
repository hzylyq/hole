package util

import (
	"io"
	"os"
)

func ReadFile(fileName string) ([]byte, error) {
	fs, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(fs)
}

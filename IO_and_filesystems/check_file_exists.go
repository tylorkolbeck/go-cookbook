package file

import (
	"errors"
	"io/fs"
	"os"
)

func FileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	// if some error other than file does not exist
	if !errors.Is(err, fs.ErrNotExist) {
		return false, err
	}
	// file does not exist
	return false, nil
}

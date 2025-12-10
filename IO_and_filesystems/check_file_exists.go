package files

import (
	"fmt"
	"os"
)

func FileExists(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); err != nil {
		return false, fmt.Errorf("error checking for file: %v")
	}

	return true, nil
}

package fileutil

import (
	"os"
)

func PathIsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	return false
}
package utils

import (
	"os"
	"path/filepath"
)

func GetExecPath() (root string) {
	bin, err := os.Executable()
	if err != nil {
		return
	}
	return filepath.Dir(bin)
}

package utils

import (
	"fmt"
	C "github.com/Weclone-org/core/constant"
	"os"
	"path/filepath"
	"runtime"
)

func ShowVersion() {
	fmt.Printf("WecloneCore %s %s %s with %s %s\n", C.Version, runtime.GOOS, runtime.GOARCH, runtime.Version(), C.BuildTime)
}

func GetExecPath() (root string) {
	bin, err := os.Executable()
	if err != nil {
		return
	}
	return filepath.Dir(bin)
}

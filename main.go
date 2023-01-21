package main

import (
	"flag"
	"fmt"
	"github.com/Weclone-org/core/executor"
	"github.com/Weclone-org/core/log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/Weclone-org/core/config"
	C "github.com/Weclone-org/core/constant"
	"github.com/Weclone-org/core/utils"
)

var (
	flagset        map[string]bool
	version        bool
	homeDir        string
	configFilePath string
	testConfig     bool
)

func init() {
	flag.StringVar(&homeDir, "d", "", "set configuration directory")
	flag.StringVar(&configFilePath, "f", "", "specify configuration file")
	flag.BoolVar(&version, "v", false, "show current version of WecloneCore")
	flag.BoolVar(&testConfig, "t", false, "test configuration and exit")
	flag.Parse()

	flagset = map[string]bool{}
	flag.Visit(func(f *flag.Flag) {
		flagset[f.Name] = true
	})
}

func main() {
	if version {
		utils.ShowVersion()
		return
	}
	if homeDir != "" {
		if !filepath.IsAbs(homeDir) {
			currentDir, _ := os.Getwd()
			homeDir = filepath.Join(currentDir, homeDir)
		}
		C.SetHomeDir(homeDir)
	} else {
		C.SetHomeDir(utils.GetExecPath())
	}
	if configFilePath != "" {
		C.SetConfigPath(configFilePath)
	} else {
		C.SetConfigPath(filepath.Join(utils.GetExecPath(), "config.yaml"))
	}
	if err := config.Init(C.HomeDir); err != nil {
		log.Fatalln("Initial configuration directory error: %s\n", err.Error())
	}
	if testConfig {
		if _, err := config.Parse(); err != nil {
			log.Throwln("Config", "configuration file test failed", err)
			os.Exit(1)
		}
		fmt.Printf("configuration file %s test is successful\n", C.ConfigFilePath)
		return
	}
	go executor.Run()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}

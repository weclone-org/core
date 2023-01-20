package constant

import (
	"path/filepath"
)

var (
	HomeDir        = ""
	ConfigFileName = ""
	ConfigFilePath = ""
)

// SetHomeDir is used to set the configuration path
func SetHomeDir(root string) {
	HomeDir = root
	ConfigFilePath = filepath.Join(root, "config.yaml")
}

// SetConfigPath is used to set the configuration file
func SetConfigPath(file string) {
	ConfigFilePath = file
	path, fileName := filepath.Split(file)
	HomeDir = path
	ConfigFileName = fileName
}

func GetHomeDir() string {
	return HomeDir
}

func GetConfigPath() string {
	return ConfigFilePath
}

func GetConfigName() string {
	return ConfigFileName
}

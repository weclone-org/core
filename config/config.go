package config

import (
	"fmt"
	C "github.com/Weclone-org/core/constant"
	"github.com/Weclone-org/core/log"
	"os"
)

// PostgreSQL config
type PostgreSQL struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	DBName string `yaml:"db_name"`
	DBUser string `yaml:"db_user"`
	DBPass string `yaml:"db_pass"`
}

// Redis config
type Redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

// Network config
type Network struct {
	Name      string     `yaml:"name"`
	Owner     string     `yaml:"owner"`
	StartTime string     `yaml:"start_time"`
	RateLimit *RateLimit `yaml:"rate_limit"`
}

// RateLimit config
type RateLimit struct {
	Minute string `yaml:"minute"`
	Hour   string `yaml:"hour"`
	Day    string `yaml:"day"`
}

// Server config
type Server struct {
	Root     string   `yaml:"root"`
	JsonURL  []string `yaml:"json_url"`
	Redirect string   `yaml:"redirect"`
}

// CDN config
type CDN struct {
	Root     string   `yaml:"root"`
	JsonURL  []string `yaml:"json_url"`
	Redirect string   `yaml:"redirect"`
}

// Root config
type Root struct {
	Network *Network `yaml:"network"`
}

// Config config struct
type Config struct {
	Mode       string      `yaml:"mode"`
	Listen     string      `yaml:"listen"`
	Level      string      `yaml:"level"`
	PostgreSQL *PostgreSQL `yaml:"postgresql"`
	Redis      *Redis      `yaml:"redis"`
	*Server
	*CDN
	*Root
}

// Init prepare necessary files
func Init(dir string) error {
	// initial homedir
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0o777); err != nil {
			return fmt.Errorf("can't create config directory %s: %s", dir, err.Error())
		}
	}

	// initial config.yaml
	if _, err := os.Stat(C.ConfigFilePath); os.IsNotExist(err) {
		log.Println("Can't find config, create a initial config file")
		f, err := os.OpenFile(C.ConfigFilePath, os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			return fmt.Errorf("can't create file %s: %s", C.ConfigFilePath, err.Error())
		}
		_, err = f.Write([]byte(`mode: server`))
		if err != nil {
			return err
		}
		err = f.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

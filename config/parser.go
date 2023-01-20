package config

import (
	"errors"
	C "github.com/Weclone-org/core/constant"
	"gopkg.in/yaml.v2"
	"os"
)

func Parse() (*Config, error) {
	rawCfg, err := os.ReadFile(C.ConfigFilePath)
	if err != nil {
		return nil, err
	}
	Cfg := Config{}
	if err := yaml.Unmarshal(rawCfg, &Cfg); err != nil {
		return nil, err
	}
	if Cfg.Mode == "" || Cfg.Listen == "" || Cfg.Level == "" || Cfg.Redis.Host == "" || Cfg.Redis.Port == "" || Cfg.PostgreSQL.Host == "" || Cfg.PostgreSQL.Port == "" || Cfg.PostgreSQL.DBName == "" || Cfg.PostgreSQL.DBPass == "" || Cfg.PostgreSQL.DBUser == "" {
		return nil, errors.New("the config file item has no value")
	}
	if Cfg.Mode == "server" {
		if Cfg.Server.Root == "" && len(Cfg.Server.JsonURL) == 0 {
			return nil, errors.New("the config item \"server\" has an unexpected value")
		}
		return &Cfg, nil
	}
	if Cfg.Mode == "cdn" {
		if Cfg.CDN.Root == "" && len(Cfg.CDN.JsonURL) == 0 {
			return nil, errors.New("the config item \"cdn\" has an unexpected value")
		}
		return &Cfg, nil
	}
	if Cfg.Mode == "root" {
		if Cfg.Root.Network.Name == "" || Cfg.Root.Network.Owner == "" || Cfg.Root.Network.StartTime == "" || Cfg.Root.Network.RateLimit.Minute == "" || Cfg.Root.Network.RateLimit.Day == "" || Cfg.Root.Network.RateLimit.Hour == "" {
			return nil, errors.New("the config item \"root\" has an unexpected value")
		}
		return &Cfg, nil
	}
	return nil, errors.New("the config item \"mode\" has an unexpected value")
}

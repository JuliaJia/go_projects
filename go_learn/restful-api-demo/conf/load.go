package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

func LoadonfigFromToml(path string) error {
	cfg := NewDefaultConfig()
	_, err := toml.DecodeFile(path, cfg)
	if err != nil {
		return err
	}
	SetGlobalConfig(cfg)

	return nil
}

func LoadConfigFromEnv() error {
	cfg := NewDefaultConfig()
	err := env.Parse(cfg)
	if err != nil {
		return err
	}
	SetGlobalConfig(cfg)
	return nil
}

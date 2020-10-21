package config

import (
	"io/ioutil"
	"os"

	yaml "github.com/jesseduffield/yaml"
)

type DefaultConfig struct {
	Cmd   string   `yaml:"cmd"`
	Hosts []string `yaml:"hosts"`
}

func LoadConfig(fileName string) (*DefaultConfig, error) {
	cfg := &DefaultConfig{}
	if _, err := os.Stat(fileName); err != nil {
		return nil, err
	}

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(content, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

package config

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Https         bool   `yaml:"https"`
	RemoteAddress string `yaml:"remoteAddress"`
	DirName       string `yaml:"dirName"`
}

func LoadConfigFile() (*Config, error) {
	data, err := os.ReadFile("conf/config.yaml")
	if err != nil {
		return &Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return &Config{}, err
	}

	return &config, nil
}

func (c *Config) GetRemoteAddress(path string) string {
	url := ""

	if c.Https {
		url += "https://"
	} else {
		url += "http://"
	}

	url += c.RemoteAddress

	if strings.HasSuffix(url, "/") {
		url = strings.TrimRight(url, "/")
	}

	url = url + path

	return url
}

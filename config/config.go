package config

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/kataras/golog"
)

// Config scheme struct for unmarshall
type Config = struct {
	TransformationsPath string `json:"transformations_path"`
}

func LoadConfig(path string) *Config {
	cfg := Config{}
	configBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error(err)
		return &cfg
	}

	if err := json.Unmarshal(configBytes, &cfg); err != nil {
		log.Error(err)
		return &cfg
	}
	return &cfg
}

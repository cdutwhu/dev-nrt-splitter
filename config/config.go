package config

import (
	"github.com/BurntSushi/toml"
)

// Config :
type ReportConfig struct {
	InFolder   string
	TrimColumn struct {
		Columns   []string
		Enable    bool
		OutFolder string
	}
	Splitting struct {
		Enable    bool
		OutFolder string
		Schema    []string
	}
}

// GetConfig :
func GetConfig(configs ...string) *ReportConfig {
	for _, config := range configs {
		cfg := &ReportConfig{}
		_, err := toml.DecodeFile(config, cfg)
		if err != nil {
			continue
		}
		return cfg
	}
	panic("Report Config File is missing or error")
}

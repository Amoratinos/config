package config_test

import (
	"fmt"
	"github.com/go-bumbu/config"
	"os"
)

type sampleConfig struct {
	Sub      submodule
	Number   int
	FromFile string `config:"from_file"`
	FromEnv  string `config:"EnvVarConfig"`
	Override string
}

type submodule struct {
	Name  string
	Value float64
}

var defaultCfg = sampleConfig{
	Number:   42,
	Override: "default Value",
}

// this is a full example of all features
func ExampleLoad() {
	// this example ignores error handling

	_ = os.Setenv("ENVPREFIX_SUB_NAME", "Superman")

	// loadconfig simulates how a user implement an app specific config loader
	loadConfig := func(overrides sampleConfig) (sampleConfig, error) {
		cfg := sampleConfig{}
		_, err := config.Load(
			config.Defaults{Item: defaultCfg},                                   // use default values
			config.Overrides{Item: overrides},                                   // set possible override values
			config.Unmarshal{Item: &cfg},                                        // marshal result into cfg
			config.CfgFile{Path: "sampledata/example_test/example.config.json"}, // load config from file
			config.EnvVar{Prefix: "ENVPREFIX"},                                  // load a config value from an env
		)
		if err != nil {
			return sampleConfig{}, err
		}
		return cfg, nil
	}

	// these values override the config, e.g. if they are set as cli flags
	overrides := sampleConfig{Override: "changed"}

	cfg, _ := loadConfig(overrides)

	// print the output
	fmt.Println(cfg.Number)   // using default value
	fmt.Println(cfg.FromFile) // from config file
	fmt.Println(cfg.Sub.Name) // loaded from env var
	fmt.Println(cfg.Override) // using override

	// Output:
	// 42
	// Spiderman
	// Superman
	// changed
}

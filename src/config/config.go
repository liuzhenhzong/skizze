package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/njpatel/loggo"

	"utils"
)

//go:generate bash bake_default_config.sh
const defaultTomlConfig = `
# This is where top level info is stored for the counter manager we
# could also use a boltDB in the DataDir but this would make it harder
# to sync over replicas since not all replicas will hold the all the
# counters.
info_dir = "~/.skizze"

# This is where the data is stored either as json or .count (pure bytes)
data_dir = "~/.skizze/data"

# The host interface for the server
host = "localhost"

# The port number for the server
port = 3596

# Treshold for saving a sketch to disk
save_threshold_seconds = 1
`

var logger = loggo.GetLogger("config")

// Config stores all configuration parameters for Go
type Config struct {
	InfoDir              string `toml:"info_dir"`
	DataDir              string `toml:"data_dir"`
	Host                 string `toml:"host"`
	Port                 int    `toml:"port"`
	SaveThresholdSeconds uint   `toml:"save_threshold_seconds"`
}

var config *Config
// InfoDir initialized from config file
var InfoDir              string
// DataDir initialized from config file
var DataDir              string
// Host initialized from config file
var Host                 string
// Port initialized from config file
var Port                 int
// SaveThresholdSeconds initialized from config file
var SaveThresholdSeconds uint

// MaxKeySize ...
const MaxKeySize int = 32768 // max key size BoltDB in bytes

func parseConfigTOML() *Config {
	cfg := &Config{}
	if _, err := toml.Decode(defaultTomlConfig, &cfg); err != nil {
		utils.PanicOnError(err)
	}

	configPath := os.Getenv("SKIZZE_CONFIG")
	if configPath != "" {
		_, err := os.Open(configPath)
		if err != nil {
			logger.Warningf("Unable to find config file, using defaults")
			return cfg
		}
		if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
			logger.Warningf("Error parsing config file, using defaults")
		}
	}

	return cfg
}

// GetConfig returns a singleton Configuration
func GetConfig() *Config {
	if config == nil {
		config = parseConfigTOML()

		if err := os.MkdirAll(config.InfoDir, os.ModePerm); err != nil {
			panic(err)
		}

		if err := os.MkdirAll(config.DataDir, os.ModePerm); err != nil {
			panic(err)
		}

		InfoDir = config.InfoDir
		DataDir = config.DataDir
		Host = config.Host
		Port = config.Port
		SaveThresholdSeconds = config.SaveThresholdSeconds
	}
	return config
}

// init initializes a singleton Configuration
func init() {
	GetConfig()
}

// Reset ...
func Reset() {
	GetConfig()
}

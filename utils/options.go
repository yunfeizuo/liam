package utils

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

// Options defines the available set of configuration options available for the
// service. The entire options struct is unmarshalled from JSON files in the
// config/ dir, to match the convention we use in the Hurley node apps.
type Options struct {
	ProjectRoot string
	Port        int
	Environment string

	// DB connection options
	Postgres struct {
		Host               string
		Port               int
		Database           string
		User               string
		Password           string
		SSL                bool
		Readonly           bool
		ConnectTimeoutSec  int
		MaxConnections     int
		MaxIdleConnections int
	} `json:"postgres"`
}

// ParseConfigFile attempts to load an Options struct, using the data in a JSON
// file in the config/ directory. The `environment` string is used to select a
// file in the directory; for example: parseConfigFile("development") would
// load config/development.json, etc.
func ParseConfigFile(configFile string) (Options, error) {
	var opts Options
	byts, err := ioutil.ReadFile(configFile)
	if err != nil {
		return opts, err
	}

	err = json.Unmarshal(byts, &opts)

	// Determine the ProjectRoot directory
	// - The PROJECT_ROOT env var takes priority over any configured value.
	// - If nothing is configured, use the default working dir.
	rootEnv := os.Getenv("PROJECT_ROOT")
	if rootEnv != "" {
		opts.ProjectRoot = rootEnv
	} else if opts.ProjectRoot == "" {
		opts.ProjectRoot, err = os.Getwd()
	}

	return opts, err
}

// ParseCommandLine parses all available command-line flags and returns the
// path to the config file to use
func ParseCommandLine() string {
	var configFile string
	flag.StringVar(&configFile, "config", "config/development.json", "The path to the JSON config file to load (defaults to 'config/development.json')")
	flag.Parse()

	return configFile
}

// ParseOptions all togather
func ParseOptions() Options {
	configFile := ParseCommandLine()
	log.Println("using config file", configFile)

	opts, err := ParseConfigFile(configFile)
	if err != nil {
		log.Fatal("fail to parse config file", err)
	}
	return opts
}

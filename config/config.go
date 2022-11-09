package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/iskrapw/utils/misc"
)

func LoadConfigFromArgs[T any]() (T, error) {
	var cfg T

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return cfg, misc.NewError("no configuration file specified")
	}
	configFile := args[0]

	log.Println("Loading configuration from", configFile)
	err := LoadConfig(configFile, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func LoadConfig[T any](path string, object *T) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	jsonFile.Close()

	err = json.Unmarshal(content, object)
	if err != nil {
		return err
	}

	return nil
}

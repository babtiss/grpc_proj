package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type CfgMap map[string]CfgMapV2

type CfgMapV2 map[string]string

func ParseConfigFile() (CfgMap, error) {

	data, err := fileRead()
	if err != nil {
		return nil, fmt.Errorf("can't read config file: %v", err)
	}
	var cfgMap CfgMap
	err = yaml.Unmarshal(data, &cfgMap)
	if err != nil {
		return nil, fmt.Errorf("can't unmarhal config file: %v", err)
	}
	return cfgMap, nil
}

func fileRead() ([]byte, error) {
	runMode := os.Getenv("RUN_MODE")
	if runMode == "local" {
		log.Println("RUN_MODE = LOCAL")
		return os.ReadFile(configLocalPath)
	}
	log.Println("RUN_MODE = ANOTHER")
	return os.ReadFile(configPath)
}

package config

import (
	// "gopkg.in/yaml.v2"
	// "io/ioutil"
	"flag"
)


type FlagValue struct {
	configFile *string
}

func ParseFlags() FlagValue {
	var flagValue FlagValue
	flagValue.configFile = flag.String("f", "", "Path to configuration file")
	flag.Parse()

	return flagValue	
}

// func (*Configurations) ParseConfigFile(path string) {
	
// }
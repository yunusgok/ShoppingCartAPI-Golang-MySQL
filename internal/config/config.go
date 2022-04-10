package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var cfgReader *configReader

type (
	Configuration struct {
		DatabaseSettings
		JwtSettings
	}

	DatabaseSettings struct {
		DatabaseURI  string
		DatabaseName string
		Username     string
		Password     string
	}

	JwtSettings struct {
		SecretKey string
	}
	configReader struct {
		configFile string
		v          *viper.Viper
	}
)

// GetAllConfigValues reads config file given as a parameter
// returns Configuration object
func GetAllConfigValues(configFile string) (configuration *Configuration, err error) {
	newConfigReader(configFile)
	if err = cfgReader.v.ReadInConfig(); err != nil {
		fmt.Printf("Failed to read config file : %s", err)
		return nil, err
	}

	err = cfgReader.v.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Failed to unmarshal yaml file to configuration struct : %s", err)
		return nil, err
	}

	return configuration, err
}

// newConfigReader creates a new reader with given fileName
func newConfigReader(configFile string) {
	v := viper.GetViper()
	v.SetConfigType("yaml")
	v.SetConfigFile(configFile)
	cfgReader = &configReader{
		configFile: configFile,
		v:          v,
	}
}

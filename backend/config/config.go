package config

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/viper"
)

const (
	configDir      = "./"
	configName     = "config"
	configFileType = ".yml"
	configType     = "yaml"
	configDefault  = "config.default.yml"
)

func Init() bool {
	isDone, err := checkForConfig()

	if err != nil {
		fmt.Println("Something went wrong Initializing the config: ", err)
		return false
	}

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configDir)

	err = viper.ReadInConfig()

	if err != nil {
		fmt.Println("Something went wrong Initializing the config 1: ", err)
		return false
	}

	return isDone
}

func checkForConfig() (bool, error) {
	fileInfo, err := os.Stat(configDir + configName + configFileType)

	if err == nil {
		return !fileInfo.IsDir(), nil
	}

	err = createConfig()

	if err != nil {
		return true, nil
	}

	return false, err
}

func createConfig() error {
	defaultConf, err := os.Open(configDir + configDefault)

	if err != nil {
		return err
	}
	defer defaultConf.Close()

	conf, err := os.Create(configDir + configName + configFileType)
	if err != nil {
		return err
	}
	defer conf.Close()

	_, err = io.Copy(conf, defaultConf)
	if err != nil {
		return err
	}

	err = conf.Sync()
	return err
}

package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var config *Config

// Config load file a valid JSON with all settings
func LoadConfig(dir string) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(currentDir, fmt.Sprintf("%s/%s", dir, "/local.json"))

	raw, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(raw, &config); err != nil {
		log.Fatal(err)
	}
}

// GetConfig returns a pointer to a Config struct which holds a valid config
func GetConfig() *Config {
	if config == nil {
		log.Fatal("config was not successfully loaded")
	}
	return config
}

func NewConfig(p string) *viper.Viper {
	envConf := os.Getenv("APP_CONF")
	if envConf == "" {
		envConf = p
	}
	fmt.Println("load conf file:", envConf)
	return getViper(envConf)
}

func getViper(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}

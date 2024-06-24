package viper

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func NewViper(p string) *viper.Viper {
	envConf := os.Getenv("APP_CONF")
	if envConf == "" {
		envConf = p
	}
	fmt.Println("load conf file:", envConf)
	return getViper(envConf)
}

func getViper(dir string) *viper.Viper {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(currentDir, dir)
	conf := viper.New()
	conf.SetConfigFile(path)
	err = conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}

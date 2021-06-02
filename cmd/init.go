package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)
var (
	CurrentEnv	=  os.Getenv("CUSTOM_RUNTIME_ENV")
	LocationEnv	=  os.Getenv("LOCATION")
)

func init()  {
	if CurrentEnv == "" {
		CurrentEnv = "dev"
	}
	if LocationEnv == "" {
		LocationEnv = "cn"
	}
	log.Println("step 0: init start env:", CurrentEnv, LocationEnv)

	dir,_ := os.Getwd()
	// init config 部署上线时，要制定下生产配置路径
	cfgFile := fmt.Sprintf(dir+"/configs/%s/%s/conf.toml",strings.ToLower(CurrentEnv), strings.ToLower(LocationEnv))
	viper.SetConfigFile(cfgFile)
	err := viper.ReadInConfig()
	if err!=nil {
		panic(err)
	}
	log.Println("step 1: load config success", viper.ConfigFileUsed())
}

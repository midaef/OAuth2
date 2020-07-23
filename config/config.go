package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	LogLevel string `toml:"log_level"`
}

type DataBaseConfig struct {
	Name     string `toml:"db_name"`
	Host     string `toml:"db_host"`
	Password string `toml:"db_password"`
}

func NewConfig() (*ServerConfig, *DataBaseConfig) {
	return &ServerConfig{
			Host:     "localhost",
			Port:     ":65000",
			LogLevel: "debug",
		}, &DataBaseConfig{
			Name:     "MyTestCluster",
			Host:     "2.95.155.99",
			Password: "8093",
		}
}

func ReadConfig(configName string) (*ServerConfig, *DataBaseConfig) {
	serverConfig, dbConfig := NewConfig()
	viper.SetConfigFile(fmt.Sprintf("config/%s", configName))
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		return serverConfig, dbConfig
	}
	return &ServerConfig{
			Host:     viper.GetString("host"),
			Port:     viper.GetString("port"),
			LogLevel: viper.GetString("log_level"),
		}, &DataBaseConfig{
			Name:     viper.GetString("db_name"),
			Host:     viper.GetString("db_host"),
			Password: viper.GetString("db_password"),
		}
}

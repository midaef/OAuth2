package main

import (
	"flag"
	"log"
	"packages/config"
	"packages/server"
)

var (
	configName string
)

func init() {
	flag.StringVar(&configName, "config", "config.toml", "config name")
}

func main() {
	flag.Parse()
	serverConfig, dbConfig := config.ReadConfig(configName)
	server := server.NewServer(serverConfig, dbConfig)
	if err := server.StartServer(); err != nil {
		log.Fatal()
	}
}

package main

import (
	"backups/internal/server"
	"flag"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "config/config.yaml", "Path to config file")
}

func main() {
	flag.Parse()

	server, err := server.NewServer(configPath)
	if err != nil {
		panic(err)
	}
	log.Println("Init Server")

	if err = server.Start(); err != nil {
		panic(err)
	}
	log.Println("Kill Server")
}

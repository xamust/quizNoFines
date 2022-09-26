package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/xamust/quizNoFines/service-grpc-gw/internal/app/service"
	"log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "Path to config file")
}

func main() {

	flag.Parse()
	config := service.NewConfig()
	meta, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	//check undecoded conf...
	if len(meta.Undecoded()) != 0 {
		log.Fatal("Undecoded configs param: ", meta.Undecoded())
	}

	if err := service.New(config).Start(); err != nil {
		log.Fatal(err)
	}
}

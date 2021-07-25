package main

import (
	"Call_A_Professional/internal/app/apiserver"
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config_path", "configs/apiserver.toml", "path to config file")
}

func main() {
	fmt.Print("Hello!!")
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

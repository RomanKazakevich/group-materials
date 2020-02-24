package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/RomanKazakevich/group-materials/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	fmt.Println("1")

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal("Decode error:   ", err)
	}

	fmt.Println("2")

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal("Server start error: ", err)
	}

	fmt.Println("3")
}

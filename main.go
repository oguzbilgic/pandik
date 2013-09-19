package main

import (
	"flag"
	"log"
	"os/user"
	"path"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	configFilePath := flag.String("c", path.Join(usr.HomeDir, ".pandik.json"), "Configuration file")
	flag.Parse()

	config, err := parseConfig(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	server, err := NewServer(config)
	if err != nil {
		log.Fatal(err)
	}

	server.Loop()
}

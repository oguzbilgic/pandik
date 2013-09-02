package main

import (
	"flag"
	"fmt"
)

func main() {
	configFilePath := flag.String("c", "~/.pandik.json", "Configuration file")
	flag.Parse()

	config, err := parseConfig(configFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	up := make(chan *Monitor, 50)
	down := make(chan *Monitor, 50)

	for _, monitor := range config.Monitors {
		go monitor.Watch(up, down)
	}

	for {
		select {
		case m := <-down:
			fmt.Println("DOWN:\t " + m.Url)
		case m := <-up:
			fmt.Println("UP:\t " + m.Url)
		}
	}
}

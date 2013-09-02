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

	monitorChan := make(chan *Monitor, 50)

	for _, monitor := range config.Monitors {
		go monitor.Watch(monitorChan)
	}

	var notifiers []Notifier
	notifiers = append(notifiers, notifyViaCLI)
	for _, notifierConf := range config.NotifierConfs {
		notifier, err := newNotifier(notifierConf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		notifiers = append(notifiers, notifier)
	}

	for {
		select {
		case m := <-monitorChan:
			for _, notifier := range notifiers {
				notifier(m)
			}
		}
	}
}

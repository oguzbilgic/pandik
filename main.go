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

	var monitors []Monitor
	for _, monitorConf := range config.MonitorConfs {
		monitor, err := NewMonitor(monitorConf)
		if err != nil {
			fmt.Println(err)
			continue
		}

		go monitor.Watch(monitorChan)
		monitors = append(monitors, *monitor)
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

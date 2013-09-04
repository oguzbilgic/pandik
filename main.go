package main

import (
	"flag"
	"fmt"
	"os/user"
	"path"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	configFilePath := flag.String("c", path.Join(usr.HomeDir, ".pandik.json"), "Configuration file")
	flag.Parse()

	config, err := parseConfig(configFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	logChan := make(chan *MonitorLog, 50)

	var monitors []Monitor
	for _, monitorConf := range config.MonitorConfs {
		monitor, err := NewMonitor(monitorConf)
		if err != nil {
			fmt.Println(err)
			continue
		}

		go monitor.Watch(logChan)
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
		log := <-logChan
		for _, notifier := range notifiers {
			notifier(log)
		}
	}
}

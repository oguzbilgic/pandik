package main

import (
	"fmt"
)

type Notifier func(*MonitorLog)

type NotifierConf struct {
	Type    string
	Address string
}

func newNotifier(nc *NotifierConf) (Notifier, error) {
	switch nc.Type {
	case "cli":
		return notifyViaCLI, nil
	}

	return nil, fmt.Errorf("ERROR:\t Not suppported notifier: %s", nc.Type)
}

func notifyViaCLI(log *MonitorLog) {
	fmt.Printf("%s: %s\n", log.Monitor.Conf.Url, log.Message)
}

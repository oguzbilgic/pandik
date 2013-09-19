package main

import (
	"errors"
	"fmt"
)

type Notifier func(*Log)

type NotifierConf struct {
	Type    string
	Address string
}

func newNotifier(nc *NotifierConf) (Notifier, error) {
	switch nc.Type {
	case "cli":
		return notifyViaCLI, nil
	}

	return nil, errors.New("not suppported notifier: " + nc.Type)
}

func notifyViaCLI(log *Log) {
	fmt.Printf("%s: %s\n", log.Monitor.Conf.Url, log.Message)
}

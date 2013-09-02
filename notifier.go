package main

import (
	"fmt"
)

type Notifier func(*Monitor)

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

func notifyViaCLI(m *Monitor) {
	if m.Up {
		fmt.Println("UP:\t " + m.Conf.Url)
	} else {
		fmt.Println("DOWN:\t " + m.Conf.Url)
	}
}

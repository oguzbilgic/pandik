package main

import (
	"fmt"
	"time"
)

type Monitor struct {
	Type    string
	Url     string
	Up      bool
	Freq    string
	Checker Checker
}

func (m *Monitor) selectChecker() error {
	switch m.Type {
	case "http-status":
		m.Checker = checkHTTPStatus
	default:
		return fmt.Errorf("ERROR:\t Not suppported check type: %s", m.Type)
	}

	return nil
}

func (m *Monitor) Watch(monitorChan chan *Monitor) {
	err := m.selectChecker()
	if err != nil || m.Checker == nil {
		return
	}

	for {
		m.Up, err = m.Checker(m)
		if err != nil {
			panic(err)
		}

		monitorChan <- m
		nextCheck, _ := time.ParseDuration(m.Freq)
		time.Sleep(nextCheck)
	}
}

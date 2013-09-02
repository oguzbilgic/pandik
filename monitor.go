package main

import (
	"fmt"
	"time"
)

type MonitorConf struct {
	Type string
	Url  string
	Freq string
}

type Monitor struct {
	Conf    *MonitorConf
	Checker Checker
	Up      bool
}

func NewMonitor(conf *MonitorConf) (*Monitor, error) {
	switch conf.Type {
	case "http-status":
		return &Monitor{conf, checkHTTPStatus, false}, nil
	}

	return nil, fmt.Errorf("ERROR:\t Not suppported checker: %s", conf.Type)
}

func (m *Monitor) Watch(monitorChan chan *Monitor) {
	for {
		newUp, err := m.Checker(m.Conf)
		if err != nil {
			panic(err)
		}

		if newUp != m.Up {
			monitorChan <- m
			m.Up = newUp
		}

		nextCheck, _ := time.ParseDuration(m.Conf.Freq)
		time.Sleep(nextCheck)
	}
}

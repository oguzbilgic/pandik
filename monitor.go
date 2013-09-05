package main

import (
	"time"
)

type MonitorConf struct {
	Type string
	Url  string
	Freq string
	Data map[string]string
}

type Monitor struct {
	Conf    *MonitorConf
	Checker Checker
	Up      bool
}

func NewMonitor(conf *MonitorConf) (*Monitor, error) {
	checker, err := GetChecker(conf.Type)
	if err != nil {
		return nil, err
	}

	return &Monitor{conf, checker, false}, nil
}

func (m *Monitor) Watch(monitorChan chan *Monitor) {
	for {
		newUp, err := m.Checker(m.Conf)
		if err != nil {
			panic(err)
		}

		m.Up = newUp
		monitorChan <- m
		nextCheck, _ := time.ParseDuration(m.Conf.Freq)
		time.Sleep(nextCheck)
	}
}

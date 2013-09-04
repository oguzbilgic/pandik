package main

import (
	"sort"
	"time"
)

type MonitorConf struct {
	Type string
	Url  string
	Freq string
	Data map[string]string
}

type MonitorLog struct {
	Up      bool
	Time    time.Time
	Message string
	Monitor *Monitor
}

func NewMonitorLog(up bool, message string) *MonitorLog {
	return &MonitorLog{up, time.Now(), message, nil}
}

type MonitorLogs []*MonitorLog

func (logs MonitorLogs) Len() int {
	return len(logs)
}

func (logs MonitorLogs) Swap(i int, j int) {
	logs[i], logs[j] = logs[j], logs[i]
}

func (logs MonitorLogs) Less(i int, j int) bool {
	return logs[i].Time.Before(logs[j].Time)
}

func (logs *MonitorLogs) Add(log *MonitorLog) {
	*logs = append(*logs, log)
	sort.Sort(logs)
}

type Monitor struct {
	Conf    *MonitorConf
	Checker Checker
	Logs    MonitorLogs
}

func NewMonitor(conf *MonitorConf) (*Monitor, error) {
	checker, err := GetChecker(conf.Type)
	if err != nil {
		return nil, err
	}

	return &Monitor{conf, checker, nil}, nil
}

func (m *Monitor) Watch(logChan chan *MonitorLog) {
	for {
		monitorLog := m.Checker(m.Conf)
		monitorLog.Monitor = m

		logChan <- monitorLog

		m.Logs.Add(monitorLog)

		nextCheck, _ := time.ParseDuration(m.Conf.Freq)
		time.Sleep(nextCheck)
	}
}

func (m *Monitor) LastLog() *MonitorLog {
	return m.Logs[m.Logs.Len()-1]
}

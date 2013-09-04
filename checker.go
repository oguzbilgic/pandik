package main

import (
	"fmt"
	"net/http"
)

type Checker func(*MonitorConf) *MonitorLog

func GetChecker(checkerType string) (Checker, error) {
	switch checkerType {
	case "http-status":
		return checkHTTPStatus, nil
	}

	return nil, fmt.Errorf("ERROR:\t Not suppported checker: %s", checkerType)
}

func checkHTTPStatus(mc *MonitorConf) *MonitorLog {
	resp, err := http.Head("http://" + mc.Url)
	if err != nil {
		return NewMonitorLog(false, err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return NewMonitorLog(false, "Http status is "+resp.Status)
	}

	return NewMonitorLog(true, "Http status code is 200")
}

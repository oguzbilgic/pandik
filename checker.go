package main

import (
	"net/http"
	"fmt"
)

type Checker func(*MonitorConf) (bool, error)

func GetChecker(checkerType string) (Checker, error) {
	switch checkerType {
	case "http-status":
		return checkHTTPStatus, nil
	}

	return nil, fmt.Errorf("ERROR:\t Not suppported checker: %s", checkerType)
}

func checkHTTPStatus(mc *MonitorConf) (bool, error) {
	resp, err := http.Head("http://" + mc.Url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false, nil
	}

	return true, nil
}

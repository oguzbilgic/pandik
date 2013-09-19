package main

import (
	"errors"
	"net/http"
)

type Checker func(*MonitorConf) *Log

func GetChecker(checkerType string) (Checker, error) {
	switch checkerType {
	case "http-status":
		return checkHTTPStatus, nil
	}

	return nil, errors.New("not suppported checker: " + checkerType)
}

func checkHTTPStatus(mc *MonitorConf) *Log {
	resp, err := http.Head("http://" + mc.Url)
	if err != nil {
		return NewLog(false, err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return NewLog(false, "Http status is "+resp.Status)
	}

	return NewLog(true, "Http status code is 200")
}

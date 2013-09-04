package main

import (
	"net/http"
)

type Checker func(*MonitorConf) (bool, error)

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

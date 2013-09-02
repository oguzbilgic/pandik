package main

import (
	"net/http"
)

type Checker func(*Monitor) (bool, error)

func checkHTTPStatus(monitor *Monitor) (bool, error) {
	resp, err := http.Get("http://" + monitor.Url)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != 200 {
		return false, nil
	}

	return true, nil
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Config struct {
	Monitors []*Monitor `json:"checks"`
}

type Checker func(*Monitor) (bool, error)

type Monitor struct {
	Type    string
	Url     string
	Freq    string
	Checker Checker
}

func parseConfig(path *string) (*Config, error) {
	configFile, err := ioutil.ReadFile(*path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (m *Monitor) selectChecker() error {
	switch m.Type {
	case "http-status":
		m.Checker = checkHTTPStatus
	default:
		return fmt.Errorf("Not suppported check type: %s", m.Type)
	}

	return nil
}

func (m *Monitor) Watch() {
	if m.Checker == nil {
		return
	}

	for {
		working, err := m.Checker(m)
		if err != nil {
			panic(err)
		}

		if !working {
			fmt.Printf("DOWN: %s - next check in 15s \n", m.Url)
			time.Sleep(15 * time.Second)
		} else {
			nextCheck, _ := time.ParseDuration(m.Freq)
			fmt.Printf("UP: %s - next check in %s \n", m.Url, nextCheck)
			time.Sleep(nextCheck)
		}
	}
}

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

func main() {
	configFile := flag.String("c", "~/.pandik.json", "Configuration file")
	flag.Parse()

	config, err := parseConfig(configFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, monitor := range config.Monitors {
		err := monitor.selectChecker()
		if err != nil {
			fmt.Println(err)
			continue
		}

		println("New monitor: " + monitor.Type + " - " + monitor.Url)
	}

	for _, monitor := range config.Monitors {
		go monitor.Watch()
	}

	for {
		time.Sleep(10 * time.Minute)
	}
}

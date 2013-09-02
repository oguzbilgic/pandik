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
	Up      bool
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
		return fmt.Errorf("ERROR:\t Not suppported check type: %s", m.Type)
	}

	return nil
}

func (m *Monitor) Watch(up chan *Monitor, down chan *Monitor) {
	err := m.selectChecker()
	if err != nil || m.Checker == nil {
		return
	}

	for {
		m.Up, err = m.Checker(m)
		if err != nil {
			panic(err)
		}

		if !m.Up {
			down <- m
			time.Sleep(15 * time.Second)
		} else {
			up <- m
			nextCheck, _ := time.ParseDuration(m.Freq)
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
	configFilePath := flag.String("c", "~/.pandik.json", "Configuration file")
	flag.Parse()

	config, err := parseConfig(configFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	up := make(chan *Monitor, 50)
	down := make(chan *Monitor, 50)

	for _, monitor := range config.Monitors {
		go monitor.Watch(up, down)
	}

	for {
		select {
		case m := <-down:
			fmt.Println("DOWN:\t " + m.Url)
		case m := <-up:
			fmt.Println("UP:\t " + m.Url)
		}
	}
}

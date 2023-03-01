package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type Config struct {
	Application struct {
		DelayedParsing bool `json:"delayed_parsing"`
		IntervalTicker int  `json:"interval_ticker_sec"`
		IntervalGet    int  `json:"interval_get_millisec"`
	} `json:"application"`
	Domains []string `json:"domains"`
	Slugs   []string `json:"slugs"`
}

func main() {
	// Open log files
	errLog, err := os.OpenFile("error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening error log file:", err)
		return
	}
	defer errLog.Close()

	successLog, err := os.OpenFile("success.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening success log file:", err)
		return
	}
	defer successLog.Close()

	// Set log output
	log.SetOutput(io.MultiWriter(os.Stdout, errLog))

	config, err := loadConfig("config.json")
	if err != nil {
		log.Println("Error loading config:", err)
		return
	}

	ticker := time.NewTicker(time.Duration(config.Application.IntervalTicker) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			checkEndpoints(config, successLog)
		}
	}
}

func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func checkEndpoints(config *Config, successLog *os.File) {
	var wg sync.WaitGroup

	for _, domain := range config.Domains {
		for _, slug := range config.Slugs {
			url := fmt.Sprintf("%s/%s", domain, slug)
			wg.Add(1)
			go func(url string) {
				defer wg.Done()

				resp, err := http.Get(url)
				if err != nil {
					log.Printf("Error checking endpoint %s: %v\n", url, err)
					return
				}
				defer resp.Body.Close()

				if resp.StatusCode != http.StatusOK {
					log.Printf("Endpoint %s returned status code %d\n", url, resp.StatusCode)
					return
				}

				// Log successful response
				fmt.Fprintf(successLog, "Endpoint %s returned status code %d\n", url, resp.StatusCode)

				if config.Application.DelayedParsing {
					// Simulate delay
					time.Sleep(time.Duration(config.Application.IntervalGet) * time.Millisecond)

					// Log delayed parsing
					fmt.Printf("Delayed parsing for endpoint %s\n", url)
				}
			}(url)
		}
	}

	wg.Wait()
}

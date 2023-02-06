package internal

import (
	"encoding/json"
	"os"
)

type CheckType string

const (
	StatusCodeCheckType CheckType = "status_code"
	TextCheckType       CheckType = "text"
)

type Config struct {
	Urls []struct {
		Url    string `json:"url"`
		Checks []struct {
			Type   CheckType `json:"type"`
			Params []string  `json:"params"`
		} `json:"checks"`
		MinChecksCnt int `json:"min_checks_cnt"`
	} `json:"urls"`
}

func ReadConfig() (*Config, error) {
	file, err := os.Open("configs/config.json")
	if err != nil {
		return nil, err
	}
	var con Config
	if err := json.NewDecoder(file).Decode(&con); err != nil {
		return nil, err
	}
	return &con, nil
}

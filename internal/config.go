package internal

type Config struct {
	Urls []struct {
		Url    string `json:"url"`
		Checks []struct {
			Name   string   `json:"name"`
			Params []string `json:"params"`
		} `json:"checks"`
		MinChecksCnt int `json:"min_checks_cnt"`
	} `json:"urls"`
}

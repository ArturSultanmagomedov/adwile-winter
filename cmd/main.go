package main

import (
	"adwile-winter/internal"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	file, err := os.Open("configs/config.json")
	if err != nil {
		return
	}
	var con internal.Config
	json.NewDecoder(file).Decode(&con)
	wg := sync.WaitGroup{}
	for _, url := range con.Urls {
		url := url
		wg.Add(1)
		go func() {
			response, err := http.Get(url.Url)
			if err != nil {
				return
			}
			defer response.Body.Close()

			a := 0
			for _, check := range url.Checks {
				switch check.Name {
				case "status_code":
					if response.StatusCode == 200 {
						a++
					}
				case "text":
					b, err := io.ReadAll(response.Body)
					if err != nil {
						return
					}
					if strings.Contains(string(b), "ok") {
						a++
					}
				}
			}
			fmt.Println(url.Url, a)
			wg.Done()
		}()
	}
	wg.Wait()
}

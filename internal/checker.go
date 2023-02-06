package internal

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Checker struct {
	con *Config
}

func NewChecker(con *Config) *Checker {
	return &Checker{con: con}
}

func (c Checker) Check(result chan<- string) {
	wg := &sync.WaitGroup{}
	go func() {
		wg.Wait()
		close(result)
	}()
	for _, url := range c.con.Urls {
		url := url
		wg.Add(1)
		go func() {
			response, err := http.Get(url.Url)
			if err != nil {
				return
			}
			defer response.Body.Close()

			a := 0
			failedTasks := ""
			for _, check := range url.Checks {
				switch check.Type {
				case StatusCodeCheckType:
					if sort.SearchStrings(check.Params, strconv.Itoa(response.StatusCode)) < len(check.Params) {
						a++
					} else {
						failedTasks += ", " + string(check.Type)
					}
				case TextCheckType:
					b, err := io.ReadAll(response.Body)
					if err != nil {
						continue
					}
					s := string(b)
					fail := true
					for _, text := range check.Params {
						if strings.Contains(s, text) {
							a++
							fail = false
							break
						}
					}
					if fail {
						failedTasks += ", " + string(check.Type)
					}
				default:
					fmt.Println("Not yet implemented.")
				}

				if a >= url.MinChecksCnt {
					result <- url.Url + ": ok"
					break
				}
			}
			if a < url.MinChecksCnt {
				result <- fmt.Sprintf("%s: fail (%s)", url.Url, failedTasks[2:])
			}
			wg.Done()
		}()
	}
}

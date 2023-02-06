package main

import (
	"adwile-winter/internal"
	"fmt"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	con, err := internal.ReadConfig()
	if err != nil {
		return err
	}

	result := make(chan string)
	checker := internal.NewChecker(con)
	checker.Check(result)

	for r := range result {
		fmt.Println(r)
	}

	return nil
}

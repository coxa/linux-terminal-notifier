package main

import (
	"github.com/coxa/linux-terminal-notifier/internal/pkg/app"
	"log"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}

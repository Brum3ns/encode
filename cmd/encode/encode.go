package main

import (
	"log"

	"github.com/Brum3ns/encode/pkg/config"
	"github.com/Brum3ns/encode/pkg/options"
	"github.com/Brum3ns/encode/pkg/runner"
)

func main() {
	if option, err := options.Option(); err == nil {
		if ok := runner.NewRunner(config.NewConfigure(option)); ok != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}

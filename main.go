package main

import (
	"errors"

	logger "github.com/yggbrazil/go-toolbox/log"
)

func main() {
	log := logger.New()
	log.Error(errors.New("teste"))
	log.Println("Println")
	log.Warn("Println")
}

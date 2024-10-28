package main

import (
	"github.com/chheller/go-rpc-todo/config"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(config.GetEnvironment().ApplicationConfiguration.LogLevel)
}

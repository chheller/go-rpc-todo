package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Environment struct {
	ApplicationConfiguration *ApplicationConfiguration
}

func (env *Environment) String() string {
	return fmt.Sprintf("TODO")
}

type EnvironmentLoader interface {
	Load(...string) error
}

// Goroutine-safe singleton reference to the parsed environment variables per https://refactoring.guru/design-patterns/singleton/go/example
var env *Environment
var lock = &sync.Mutex{}

func GetEnvironment(load ...func(...string) error) *Environment {
	if env != nil {
		return env
	}

	lock.Lock()
	defer lock.Unlock()
	if env == nil {
		log.Info("Enivironment uninitialized, loading environment variables from .env file")
		var err error
		if len(load) == 0 {
			err = godotenv.Load(".env")
		} else if len(load) == 1 {
			err = load[0](".env")
		} else {
			log.Panic("Too many arguments passed to GetEnvironment, expected 1")
		}
		if err != nil {
			log.Panicf("Error loading .env file\n Error: %v", err)
		}
		env = &Environment{
			ApplicationConfiguration: loadApplicationConfiguration(),
		}
	}
	return env
}

// Panics if the environment variable is not set
func getEnvironmentVariableOrPanic(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Panicf("Missing required environment variable %s", key)
	}
	return value
}

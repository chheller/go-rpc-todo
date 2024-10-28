package config

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type ApplicationConfiguration struct {
	Port                uint32
	LogLevel            log.Level
	LoggerReportCaller  bool
	HttpPrintDebugError bool
	CacheTTL            int64
}

// Loads general application configurations and packages them into a struct. Handles default values,
// parsing strings into valid number types, and panics if any required variables are missing.
func loadApplicationConfiguration() *ApplicationConfiguration {
	var logLevel log.Level
	logLevelStr, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		log.Info("No LOG_LEVEL environment variable found, defaulting to INFO")
	} else {
		var err error
		logLevel, err = log.ParseLevel(logLevelStr)
		if err != nil {
			log.Panicf("LOG_LEVEL must be one of logrus.Level, got %s\nError: %v", os.Getenv("LOG_LEVEL"), err)
		}
	}
	// Note: If we want to use anything outside of the default INFO level log, we need to set it immediately after loading the variable.
	// Any log messages logged before this method is called outside of the INFO scope will be ignored
	log.SetLevel(logLevel)

	var port uint32
	portStr, ok := os.LookupEnv("PORT")
	if !ok {
		log.Info("No PORT environment variable found, defaulting to 8080")
		port = 8080
	} else {
		p, err := strconv.ParseUint(portStr, 10, 32)
		if err != nil {
			log.Panicf("LOG_LEVEL must be an integer, got %s\nError: %v", os.Getenv("LOG_LEVEL"), err)
		}
		port = uint32(p)
	}
	return &ApplicationConfiguration{
		Port:                port,
		LogLevel:            logLevel,
		LoggerReportCaller:  os.Getenv("LOGGER_REPORT_CALLER") == "true",
		HttpPrintDebugError: os.Getenv("HTTP_PRINT_DEBUG_ERROR") == "true",
	}
}

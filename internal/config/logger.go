package config

import (
	log "github.com/sirupsen/logrus"
)

// initLogger configures global logrus logger
func initLogger(loglevel string) (*log.Logger, error) {
	// New logger to pass to home assistant client
	logger := log.New()
	logger.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	level, err := log.ParseLevel(loglevel)
	if err != nil {
		return log.New(), err
	}
	logger.SetLevel(level)

	// Global logger
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	log.SetLevel(level)
	return logger, nil
}

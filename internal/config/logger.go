package config

import (
	log "github.com/sirupsen/logrus"
)

// InitLogger configures global logrus logger
func (w *WitchConfig) InitLogger() (*log.Logger, error) {
	logger := log.New()
	logger.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	level, err := log.ParseLevel(w.LogLevel)
	if err != nil {
		return log.New(), err
	}
	logger.SetLevel(level)
	return logger, nil
}

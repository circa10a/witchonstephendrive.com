package config

import (
	log "github.com/sirupsen/logrus"
)

// InitLogger configures global logrus logger
func (w *WitchConfig) InitLogger() error {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	level, err := log.ParseLevel(w.LogLevel)
	if err != nil {
		return err
	}
	log.SetLevel(level)
	return nil
}

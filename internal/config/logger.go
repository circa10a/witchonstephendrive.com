package config

import (
	log "github.com/sirupsen/logrus"
)

// InitLogger configures global logrus logger
func InitLogger() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

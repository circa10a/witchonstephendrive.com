package config

import (
	"errors"
	"fmt"
	"io"
	"syscall"
	"time"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

const homeAssistantTimeoutSeconds = 10
const homeAssistantRetryCount = 3
const homeAssistantRetryWaitSeconds = 5

// InitHomeAssistantConfig sets the initial home assistant endpoint and REST client
func (w *WitchConfig) InitHomeAssistantConfig(log *log.Logger) {
	// Create client to be used with home assistant
	homeAssistantEndpoint := fmt.Sprintf("%s:%d", w.HomeAssistantHost, w.HomeAssistantPort)
	w.HomeAssistantClient = resty.New().SetLogger(log)
	w.HomeAssistantClient.SetHostURL(homeAssistantEndpoint)
	w.HomeAssistantClient.SetHeader("Content-Type", "application/json")
	w.HomeAssistantClient.SetAuthToken(w.HomeAssistantAPIToken)
	w.HomeAssistantClient.SetTimeout(time.Second * homeAssistantTimeoutSeconds)
	w.HomeAssistantClient.SetRetryCount(homeAssistantRetryCount)
	// Only retry on connection refused or EOF meaning possible reboot
	w.HomeAssistantClient.SetRetryWaitTime(homeAssistantRetryWaitSeconds * time.Second).AddRetryCondition(
		func(_ *resty.Response, err error) bool {
			return errors.Is(err, syscall.ECONNREFUSED) || errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF)
		},
	)
}

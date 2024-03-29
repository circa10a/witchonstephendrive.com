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

const homeAssistantTimeout = 10 * time.Second
const homeAssistantRetryCount = 3
const homeAssistantRetryWait = 5 * time.Second

// initHomeAssistantClient sets the initial home assistant endpoint and REST client
func (w *WitchConfig) initHomeAssistantClient(log *log.Logger) {
	// Create client to be used with home assistant
	homeAssistantEndpoint := fmt.Sprintf("%s:%d", w.HomeAssistantHost, w.HomeAssistantPort)
	w.HomeAssistantClient = resty.New().SetLogger(log)
	w.HomeAssistantClient.SetBaseURL(homeAssistantEndpoint)
	w.HomeAssistantClient.SetHeader("Content-Type", "application/json")
	w.HomeAssistantClient.SetAuthToken(w.HomeAssistantAPIToken)
	w.HomeAssistantClient.SetTimeout(homeAssistantTimeout)
	w.HomeAssistantClient.SetRetryCount(homeAssistantRetryCount)
	// Only retry on connection refused or EOF meaning possible reboot
	w.HomeAssistantClient.SetRetryWaitTime(homeAssistantRetryWait).AddRetryCondition(
		func(_ *resty.Response, err error) bool {
			return errors.Is(err, syscall.ECONNREFUSED) || errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF)
		},
	)
}

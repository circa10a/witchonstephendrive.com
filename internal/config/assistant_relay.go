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

// This should only be as long as the longest sound to play
const assistantRelayTimeoutSeconds = 45
const assistantRelayRetryCount = 1
const assistantRelayRetryWaitSeconds = 5

// InitAssistantRelayConfig sets the initial relay endpoint and relay REST client
func (w *WitchConfig) InitAssistantRelayConfig(log *log.Logger) {
	// Create client to be used with assistant relay
	assitantRelayEndpoint := fmt.Sprintf("%s:%d", w.AssistantRelayHost, w.AssistantRelayPort)
	w.RelayClient = resty.New().SetLogger(log)
	w.RelayClient.SetHostURL(assitantRelayEndpoint)
	w.RelayClient.SetHeader("Content-Type", "application/json")
	w.RelayClient.SetTimeout(time.Second * assistantRelayTimeoutSeconds)
	w.RelayClient.SetRetryCount(assistantRelayRetryCount)
	// Only retry on connection refused or EOF meaning possible reboot
	w.RelayClient.SetRetryWaitTime(assistantRelayRetryWaitSeconds * time.Second).AddRetryCondition(
		func(_ *resty.Response, err error) bool {
			return errors.Is(err, syscall.ECONNREFUSED) || errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF)
		},
	)
}

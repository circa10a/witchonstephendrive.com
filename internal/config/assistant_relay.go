package config

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

// This should only be as long as the longest sound to play
const relayAssistantTimeoutSeconds = 45

// InitAssistantRelayConfig sets the initial relay endpoint and relay REST client
func (w *WitchConfig) InitAssistantRelayConfig() {
	// Create client to be used with assistant relay
	assitantRelayEndpoint := fmt.Sprintf("%s:%d", w.AssistantRelayHost, w.AssistantRelayPort)
	w.RelayClient = resty.New()
	w.RelayClient.SetHostURL(assitantRelayEndpoint)
	w.RelayClient.SetHeader("Content-Type", "application/json")
	w.RelayClient.SetTimeout(time.Second * relayAssistantTimeoutSeconds)
}

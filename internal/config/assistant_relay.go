package config

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// InitAssistantRelayConfig sets the initial relay endpoint and relay REST client
func (w *WitchConfig) InitAssistantRelayConfig() {
	// Create client to be used with assistant relay
	assitantRelayEndpoint := fmt.Sprintf("%s:%d", w.AssistantRelayHost, w.AssistantRelayPort)
	w.RelayClient = resty.New().SetHostURL(assitantRelayEndpoint).SetHeader("Content-Type", "application/json")
}

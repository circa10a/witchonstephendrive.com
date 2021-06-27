package config

import "github.com/amimof/huego"

type WitchConfig struct {
	Port                   int           `envconfig:"PORT" default:"8080"`
	Metrics                bool          `envconfig:"METRICS" default:"true"`
	HueUser                string        `envconfig:"HUE_USER" required:"true"`
	HueLights              []int         `envconfig:"HUE_LIGHTS" required:"true" split_words:"true"`
	AssistantRelayEndpoint string        `envconfig:"ASSISTANT_RELAY_ENDPOINT" default:"127.0.0.1"`
	Bridge                 *huego.Bridge `ignored:"true"`
}

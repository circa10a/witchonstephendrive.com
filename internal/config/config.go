package config

import (
	"github.com/amimof/huego"
	"github.com/go-resty/resty/v2"
)

type WitchConfig struct {
	Port               int           `envconfig:"PORT" default:"8080"`
	Metrics            bool          `envconfig:"METRICS" default:"true"`
	HueUser            string        `envconfig:"HUE_USER" required:"true"`
	HueLights          []int         `envconfig:"HUE_LIGHTS" required:"true" split_words:"true"`
	AssistantRelayHost string        `envconfig:"ASSISTANT_RELAY_HOST" default:"http://127.0.0.1"`
	AssistantRelayPort int           `envconfig:"ASSISTANT_RELAY_PORT" default:"3000"`
	AssistantDevice    string        `envconfig:"ASSISTANT_DEVICE" required:"true"`
	Bridge             *huego.Bridge `ignored:"true"`
	Client             *resty.Client `ignored:"true"`
}

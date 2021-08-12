package config

import (
	"github.com/amimof/huego"
	"github.com/go-resty/resty/v2"
	"github.com/oleiade/lane"
)

// WitchConfig is a global config struct which holds all settings and some stateful objects
type WitchConfig struct {
	APIBaseURL               string        `envconfig:"WITCH_API_BASE_URL" default:"/api/v1"`
	APIEnabled               bool          `envconfig:"WITCH_API_ENABLED" default:"true"`
	Port                     int           `envconfig:"WITCH_PORT" default:"8080"`
	AssistantDevice          string        `envconfig:"WITCH_ASSISTANT_DEVICE" required:"true"`
	AssistantRelayHost       string        `envconfig:"WITCH_ASSISTANT_RELAY_HOST" default:"http://127.0.0.1"`
	AssistantRelayPort       int           `envconfig:"WITCH_ASSISTANT_RELAY_PORT" default:"3000"`
	HueBridge                *huego.Bridge `ignored:"true"`
	HueBridgeRefreshInterval int           `envconfig:"WITCH_HUE_BRIDGE_REFRESH_INTERVAL" default:"21600"`
	HueLights                []int         `envconfig:"WITCH_HUE_LIGHTS" required:"true" split_words:"true"`
	HueLightsStructs         []huego.Light `ignored:"true"`
	HueToken                 string        `envconfig:"WITCH_HUE_TOKEN" required:"true"`
	MetricsEnabled           bool          `envconfig:"WITCH_METRICS_ENABLED" default:"true"`
	RelayClient              *resty.Client `ignored:"true"`
	SoundQuietTimeStart      int           `envconfig:"WITCH_SOUND_QUIET_TIME_START" default:"22"`
	SoundQuietTimeEnd        int           `envconfig:"WITCH_SOUND_QUIET_TIME_END" default:"07"`
	SoundQueueCapacity       int           `envconfig:"WITCH_SOUND_QUEUE_CAPACITY" default:"3"`
	SoundQueuePollInterval   int           `envconfig:"WITCH_SOUND_QUEUE_POLL_INTERVAL" default:"1"`
	SoundQueue               *lane.Deque   `ignored:"true"`
	UIEnabled                bool          `envconfig:"WITCH_UI_ENABLED" default:"true"`
}

package config

import (
	"github.com/amimof/huego"
	"github.com/go-resty/resty/v2"
	"github.com/oleiade/lane"
)

// WitchConfig is a global config struct which holds all settings and some stateful objects
type WitchConfig struct {
	APIBaseURL               string         `envconfig:"WITCH_API_BASE_URL" default:"/api/v1"`
	APIEnabled               bool           `envconfig:"WITCH_API_ENABLED" default:"true"`
	Port                     int            `envconfig:"WITCH_PORT" default:"8080"`
	AssistantDevice          string         `envconfig:"WITCH_ASSISTANT_DEVICE" default:""`
	AssistantRelayHost       string         `envconfig:"WITCH_ASSISTANT_RELAY_HOST" default:"http://127.0.0.1"`
	AssistantRelayPort       int            `envconfig:"WITCH_ASSISTANT_RELAY_PORT" default:"3000"`
	HueBridge                *huego.Bridge  `ignored:"true"`
	HueBridgeRefreshInterval int            `envconfig:"WITCH_HUE_BRIDGE_REFRESH_INTERVAL" default:"6"`
	HueDefaultColors         map[int]string `envconfig:"WITCH_HUE_DEFAULT_COLORS" default:""`
	HueDefaultColorsEnabled  bool           `envconfig:"WITCH_HUE_DEFAULT_COLORS_ENABLED" default:"false"`
	HueDefaultColorsStart    int            `envconfig:"WITCH_HUE_DEFAULT_COLORS_START" default:"22"`
	HueLights                []int          `envconfig:"WITCH_HUE_LIGHTS" required:"true" split_words:"true"`
	HueLightsScheduleEnabled bool           `envconfig:"WITCH_HUE_LIGHTS_SCHEDULE_ENABLED" default:"false"`
	HueLightsStart           int            `envconfig:"WITCH_HUE_LIGHTS_START" default:"18"`
	HueLightsEnd             int            `envconfig:"WITCH_HUE_LIGHTS_END" default:"7"`
	HueLightsStructs         []huego.Light  `ignored:"true"`
	HueToken                 string         `envconfig:"WITCH_HUE_TOKEN" required:"true"`
	MetricsEnabled           bool           `envconfig:"WITCH_METRICS_ENABLED" default:"true"`
	RelayClient              *resty.Client  `ignored:"true"`
	SoundQueue               *lane.Deque    `ignored:"true"`
	SoundQueueCapacity       int            `envconfig:"WITCH_SOUND_QUEUE_CAPACITY" default:"3"`
	SoundQuietTimeEnabled    bool           `envconfig:"WITCH_SOUND_QUIET_TIME_ENABLED" default:"true"`
	SoundQuietTimeEnd        int            `envconfig:"WITCH_SOUND_QUIET_TIME_END" default:"07"`
	SoundQuietTimeStart      int            `envconfig:"WITCH_SOUND_QUIET_TIME_START" default:"22"`
	UIEnabled                bool           `envconfig:"WITCH_UI_ENABLED" default:"true"`
}

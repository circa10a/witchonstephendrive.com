package config

import (
	"time"

	"github.com/amimof/huego"
	"github.com/go-resty/resty/v2"
	"github.com/oleiade/lane"
)

// WitchConfig is a global config struct which holds all settings and some stateful objects
type WitchConfig struct {
	APIBaseURL                  string         `envconfig:"API_BASE_URL" default:"/api/v1"`
	Port                        int            `envconfig:"PORT" default:"8080"`
	HomeAssistantEntityID       string         `envconfig:"HOME_ASSISTANT_ENTITY_ID" default:""`
	HomeAssistantAPIToken       string         `envconfig:"HOME_ASSISTANT_API_TOKEN" default:""`
	HomeAssistantHost           string         `envconfig:"HOME_ASSISTANT_HOST" default:"http://127.0.0.1"`
	HomeAssistantPort           int            `envconfig:"HOME_ASSISTANT_PORT" default:"8123"`
	HueBridge                   *huego.Bridge  `ignored:"true"`
	HueBridgeRefreshInterval    time.Duration  `envconfig:"HUE_BRIDGE_REFRESH_INTERVAL" default:"6h"`
	HueDefaultColors            map[int]string `envconfig:"HUE_DEFAULT_COLORS" default:""`
	HueDefaultColorsEnabled     bool           `envconfig:"HUE_DEFAULT_COLORS_ENABLED" default:"false"`
	HueDefaultColorsStart       int            `envconfig:"HUE_DEFAULT_COLORS_START" default:"22"`
	HueLights                   []int          `envconfig:"HUE_LIGHTS" required:"true" split_words:"true"`
	HueLightsScheduleEnabled    bool           `envconfig:"HUE_LIGHTS_SCHEDULE_ENABLED" default:"false"`
	HueLightsStart              int            `envconfig:"HUE_LIGHTS_START" default:"18"`
	HueLightsEnd                int            `envconfig:"HUE_LIGHTS_END" default:"7"`
	HueLightsStructs            []huego.Light  `ignored:"true"`
	HueToken                    string         `envconfig:"HUE_TOKEN" required:"true"`
	LogLevel                    string         `envconfig:"LOG_LEVEL" default:"info"`
	MetricsEnabled              bool           `envconfig:"METRICS_ENABLED" default:"true"`
	HomeAssistantClient         *resty.Client  `ignored:"true"`
	ShowBanner                  bool           `envconfig:"SHOW_BANNER" default:"true"`
	SoundQueue                  *lane.Deque    `ignored:"true"`
	SoundQueueCapacity          int            `envconfig:"SOUND_QUEUE_CAPACITY" default:"1"`
	SoundQuietTimeEnabled       bool           `envconfig:"SOUND_QUIET_TIME_ENABLED" default:"true"`
	SoundQuietTimeEnd           int            `envconfig:"SOUND_QUIET_TIME_END" default:"07"`
	SoundQuietTimeStart         int            `envconfig:"SOUND_QUIET_TIME_START" default:"22"`
	SoundQueueWaitUntilFinished bool           `envconfig:"SOUND_QUEUE_WAIT_UNTIL_FINISHED" default:"true"`
	UIEnabled                   bool           `envconfig:"UI_ENABLED" default:"true"`
}

package config

import (
	"time"

	"github.com/amimof/huego"
	"github.com/circa10a/go-geofence"
	"github.com/go-resty/resty/v2"
	"github.com/oleiade/lane"
)

// WitchConfig is a global config struct which holds all settings and some stateful objects
type WitchConfig struct {
	HueBridge                   *huego.Bridge     `ignored:"true"`
	HueDefaultColors            map[int]string    `envconfig:"HUE_DEFAULT_COLORS" default:""`
	SoundQueue                  *lane.Deque       `ignored:"true"`
	HomeAssistantClient         *resty.Client     `ignored:"true"`
	HomeAssistantAPIToken       string            `envconfig:"HOME_ASSISTANT_API_TOKEN" default:""`
	LogLevel                    string            `envconfig:"LOG_LEVEL" default:"info"`
	HomeAssistantEntityID       string            `envconfig:"HOME_ASSISTANT_ENTITY_ID" default:""`
	GeofencingFreeGeoIPAPIToken string            `envconfig:"GEOFENCING_FREEGEOIP_API_TOKEN" default:""`
	HomeAssistantHost           string            `envconfig:"HOME_ASSISTANT_HOST" default:"http://127.0.0.1"`
	HueToken                    string            `envconfig:"HUE_TOKEN" required:"true"`
	APIBaseURL                  string            `envconfig:"API_BASE_URL" default:"/api/v1"`
	HueLightsStructs            []huego.Light     `ignored:"true"`
	HueLights                   []int             `envconfig:"HUE_LIGHTS" required:"true" split_words:"true"`
	GeofencingClient            geofence.Geofence `ignored:"true"`
	SoundQuietTimeEnd           int               `envconfig:"SOUND_QUIET_TIME_END" default:"07"`
	HueDefaultColorsStart       int               `envconfig:"HUE_DEFAULT_COLORS_START" default:"22"`
	Port                        int               `envconfig:"PORT" default:"8080"`
	HueLightsStart              int               `envconfig:"HUE_LIGHTS_START" default:"18"`
	HueLightsEnd                int               `envconfig:"HUE_LIGHTS_END" default:"7"`
	HueBridgeRefreshInterval    time.Duration     `envconfig:"HUE_BRIDGE_REFRESH_INTERVAL" default:"6h"`
	HomeAssistantPort           int               `envconfig:"HOME_ASSISTANT_PORT" default:"8123"`
	GeofencingSensitivity       int               `envconfig:"GEOFENCING_SENSITIVITY" default:"3"`
	SoundQueueCapacity          int               `envconfig:"SOUND_QUEUE_CAPACITY" default:"1"`
	SoundQuietTimeStart         int               `envconfig:"SOUND_QUIET_TIME_START" default:"22"`
	HueLightsScheduleEnabled    bool              `envconfig:"HUE_LIGHTS_SCHEDULE_ENABLED" default:"false"`
	ShowBanner                  bool              `envconfig:"SHOW_BANNER" default:"true"`
	MetricsEnabled              bool              `envconfig:"METRICS_ENABLED" default:"true"`
	SoundQuietTimeEnabled       bool              `envconfig:"SOUND_QUIET_TIME_ENABLED" default:"true"`
	HueDefaultColorsEnabled     bool              `envconfig:"HUE_DEFAULT_COLORS_ENABLED" default:"false"`
	GeofencingEnabled           bool              `envconfig:"GEOFENCING_ENABLED" default:"false"`
	SoundQueueWaitUntilFinished bool              `envconfig:"SOUND_QUEUE_WAIT_UNTIL_FINISHED" default:"true"`
	UIEnabled                   bool              `envconfig:"UI_ENABLED" default:"true"`
}

package config

import (
	"sync"
	"time"

	"github.com/amimof/huego"
	"github.com/circa10a/go-geofence"
	"github.com/go-resty/resty/v2"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

// All environment variables for configuration expect WITCH_ prefix
const envVarPrefix = "witch"

// WitchConfig is a global config struct which holds all settings and some stateful objects
type WitchConfig struct {
	APIBaseURL string `envconfig:"API_BASE_URL" default:"/api/v1"`

	HueBridge                *huego.Bridge  `ignored:"true"`
	HueBridgeRefreshInterval time.Duration  `envconfig:"HUE_BRIDGE_REFRESH_INTERVAL" default:"6h"`
	HueDefaultColors         map[int]string `envconfig:"HUE_DEFAULT_COLORS" default:""`
	HueDefaultColorsEnabled  bool           `envconfig:"HUE_DEFAULT_COLORS_ENABLED" default:"false"`
	HueDefaultColorsStart    int            `envconfig:"HUE_DEFAULT_COLORS_START" default:"22"`
	HueLightsStart           int            `envconfig:"HUE_LIGHTS_START" default:"18"`
	HueLightsEnd             int            `envconfig:"HUE_LIGHTS_END" default:"7"`
	HueLightsStructs         []huego.Light  `ignored:"true"`
	HueLightsScheduleEnabled bool           `envconfig:"HUE_LIGHTS_SCHEDULE_ENABLED" default:"false"`
	HueLights                []int          `envconfig:"HUE_LIGHTS" required:"true" split_words:"true"`
	HueToken                 string         `envconfig:"HUE_TOKEN" required:"true"`

	SoundQueue            chan string `ignored:"true"`
	SoundQueueCapacity    int         `envconfig:"SOUND_QUEUE_CAPACITY" default:"1"`
	SoundQuietTimeStart   int         `envconfig:"SOUND_QUIET_TIME_START" default:"22"`
	SoundQuietTimeEnabled bool        `envconfig:"SOUND_QUIET_TIME_ENABLED" default:"true"`
	SoundQuietTimeEnd     int         `envconfig:"SOUND_QUIET_TIME_END" default:"07"`

	GeofencingEnabled        bool              `envconfig:"GEOFENCING_ENABLED" default:"false"`
	GeofencingClient         geofence.Geofence `ignored:"true"`
	GeofencingIPBaseAPIToken string            `envconfig:"GEOFENCING_IPBASE_API_TOKEN" default:""`
	GeofencingRadius         float64           `envconfig:"GEOFENCING_RADIUS" default:"0.5"`

	HomeAssistantClient   *resty.Client `ignored:"true"`
	HomeAssistantAPIToken string        `envconfig:"HOME_ASSISTANT_API_TOKEN" default:""`
	HomeAssistantEntityID string        `envconfig:"HOME_ASSISTANT_ENTITY_ID" default:""`
	HomeAssistantHost     string        `envconfig:"HOME_ASSISTANT_HOST" default:"http://127.0.0.1"`
	HomeAssistantPort     int           `envconfig:"HOME_ASSISTANT_PORT" default:"8123"`

	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`

	MetricsEnabled bool `envconfig:"METRICS_ENABLED" default:"true"`

	Port int `envconfig:"PORT" default:"8080"`

	ShowBanner bool `envconfig:"SHOW_BANNER" default:"false"`

	UIEnabled    bool `envconfig:"UI_ENABLED" default:"true"`
	sync.RWMutex `ignored:"true"`
}

// Returns a new config and inits needed daemons
func New() *WitchConfig {
	w := &WitchConfig{}
	err := envconfig.Process(envVarPrefix, w)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Show HAPPY HALLOWEEN banner
	if w.ShowBanner {
		printBanner()
	}

	// Logger Config
	log, err := initLogger(w.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	// Sounds
	// Home Assistant Config such as endpoint and client
	w.initHomeAssistantClient(log)
	// Creates initial capped sounds queue
	w.SoundQueue = make(chan string, w.SoundQueueCapacity)

	// Geofencing
	// Setup client
	w.initGeofencingClient()

	return w
}

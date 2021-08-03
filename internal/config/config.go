package config

import (
	"github.com/amimof/huego"
	"github.com/go-resty/resty/v2"
	"github.com/oleiade/lane"
)

type WitchConfig struct {
	Port                   int           `envconfig:"PORT" default:"8080"`
	Metrics                bool          `envconfig:"METRICS" default:"true"`
	HueUser                string        `envconfig:"HUE_USER" required:"true"`
	HueLights              []int         `envconfig:"HUE_LIGHTS" required:"true" split_words:"true"`
	AssistantRelayHost     string        `envconfig:"ASSISTANT_RELAY_HOST" default:"http://127.0.0.1"`
	AssistantRelayPort     int           `envconfig:"ASSISTANT_RELAY_PORT" default:"3000"`
	AssistantDevice        string        `envconfig:"ASSISTANT_DEVICE" required:"true"`
	SoundQuietTimeStart    string        `envconfig:"SOUND_QUIET_TIME_START" default:"10:30PM"`
	SoundQuietTimeEnd      string        `envconfig:"SOUND_QUIET_TIME_END" default:"7:00AM"`
	SoundQueueCapacity     int           `envconfig:"SOUND_QUEUE_CAPACITY" default:"3"`
	SoundQueuePollInterval int           `envconfig:"SOUND_QUEUE_POLL_INTERVAL" default:"1"`
	SoundQueue             *lane.Deque   `ignored:"true"`
	Bridge                 *huego.Bridge `ignored:"true"`
	HueLightsStructs       []huego.Light `ignored:"true"`
	RelayClient            *resty.Client `ignored:"true"`
}

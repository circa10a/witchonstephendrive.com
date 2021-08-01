package config

import (
	"context"

	"github.com/amimof/huego"
	"github.com/go-redis/redis/v8"
	"github.com/go-resty/resty/v2"
)

type WitchServerConfig struct {
	Port             int             `envconfig:"PORT" default:"8080"`
	Metrics          bool            `envconfig:"METRICS" default:"true"`
	HueUser          string          `envconfig:"HUE_USER" required:"true"`
	HueLights        []int           `envconfig:"HUE_LIGHTS" required:"true" split_words:"true"`
	RedisHost        string          `envconfig:"REDIS_HOST" default:"127.0.0.1"`
	RedisPort        int             `envconfig:"REDIS_PORT" default:"6379"`
	RedisPassword    string          `envconfig:"REDIS_PASSWORD" default:""`
	RedisChannel     string          `envconfig:"REDIS_CHANNEL" default:"sounds"`
	RedisClient      *redis.Client   `ignored:"true"`
	RedisContext     context.Context `ignored:"true"`
	Bridge           *huego.Bridge   `ignored:"true"`
	HueLightsStructs []huego.Light   `ignored:"true"`
}

type WitchClientConfig struct {
	AssistantRelayHost string `envconfig:"ASSISTANT_RELAY_HOST" default:"http://127.0.0.1"`
	AssistantRelayPort int    `envconfig:"ASSISTANT_RELAY_PORT" default:"3000"`
	AssistantDevice    string `envconfig:"ASSISTANT_DEVICE" required:"true"`
	RedisHost          string `envconfig:"REDIS_HOST" default:"127.0.0.1"`
	RedisPort          int    `envconfig:"REDIS_PORT" default:"6379"`
	RedisPassword      string `envconfig:"REDIS_PASSWORD" default:""`
	RedisChannel       string `envconfig:"REDIS_CHANNEL" default:"sounds"`
	// In seconds
	RedisReadInterval int             `envconfig:"REDIS_READ_INTERVAL" default:"1"`
	RedisClient       *redis.Client   `ignored:"true"`
	RedisContext      context.Context `ignored:"true"`
	RestClient        *resty.Client   `ignored:"true"`
}

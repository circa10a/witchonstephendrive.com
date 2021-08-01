package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/amimof/huego"
	"github.com/circa10a/witchonstephendrive.com/internal/config"
	"github.com/circa10a/witchonstephendrive.com/internal/routes"
	"github.com/go-redis/redis/v8"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start witch server",
	Long: `The witch server hosts the REST API, web UI, changes light states,
and writes sound messages to redis for the witch client to consume.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Setup config
		var witchConfig config.WitchServerConfig
		err := envconfig.Process("", &witchConfig)
		if err != nil {
			log.Fatal(err.Error())
		}

		// Hue Config
		// Find hue bridge ip
		hueBridge, err := huego.Discover()
		if err != nil {
			log.Fatal(err)
		}
		witchConfig.Bridge = hueBridge
		// Authenticate against bridge api
		witchConfig.Bridge.Login(witchConfig.HueUser)
		// Store all light data to be used later
		for _, lightID := range witchConfig.HueLights {
			light, err := witchConfig.Bridge.GetLight(lightID)
			if err != nil {
				log.Fatal(err)
			}
			witchConfig.HueLightsStructs = append(witchConfig.HueLightsStructs, *light)
		}

		// Redis Client
		witchConfig.RedisClient = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", witchConfig.RedisHost, witchConfig.RedisPort),
			Password: witchConfig.RedisPassword,
		})
		witchConfig.RedisContext = context.Background()

		// Echo Server
		e := echo.New()
		// Prometheus metrics
		if witchConfig.Metrics {
			prometheus.NewPrometheus("echo", nil).Use(e)
		}
		// Use logging middleware
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "[${time_rfc3339}] ${status} ${method} ${path} (${remote_ip}) ${latency_human}\n",
			Output: e.Logger.Output(),
		}))
		// Declare routes
		routes.Routes(e, witchConfig, frontendAssetsFS, apiDocAssetsFS)
		// Start App
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", witchConfig.Port)))
	},
}

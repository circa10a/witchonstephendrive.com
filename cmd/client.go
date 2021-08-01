package cmd

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/circa10a/witchonstephendrive.com/internal/config"
	"github.com/circa10a/witchonstephendrive.com/internal/sounds"
	"github.com/go-redis/redis/v8"
	"github.com/go-resty/resty/v2"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(clientCmd)
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start witch client",
	Long:  "The witch client reads messages from redis and triggers playing sounds in order.",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Started witch client")
		// Setup config
		var witchConfig config.WitchClientConfig
		err := envconfig.Process("", &witchConfig)
		if err != nil {
			log.Fatal(err.Error())
		}

		// Create client to be used with assistant relay
		assitantRelayEndpoint := fmt.Sprintf("%v:%v", witchConfig.AssistantRelayHost, witchConfig.AssistantRelayPort)
		witchConfig.RestClient = resty.New().SetHostURL(assitantRelayEndpoint).SetHeader("Content-Type", "application/json")

		// Redis Client
		witchConfig.RedisClient = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", witchConfig.RedisHost, witchConfig.RedisPort),
			Password: witchConfig.RedisPassword,
		})
		witchConfig.RedisContext = context.Background()
		pubsub := witchConfig.RedisClient.Subscribe(witchConfig.RedisContext, witchConfig.RedisChannel)

		for {
			time.Sleep(time.Second * time.Duration(witchConfig.RedisReadInterval))
			msg, err := pubsub.ReceiveMessage(witchConfig.RedisContext)
			if err != nil {
				log.Error(err)
			} else {
				log.Infof("Read sound message: %s from channel: %s. Playing...", msg.Payload, msg.Channel)
				// Call the assistant relay
				resp, err := witchConfig.RestClient.R().SetBody(&sounds.PlaySoundPayload{
					Device: witchConfig.AssistantDevice,
					Source: fmt.Sprintf("%v.mp3", msg.Payload),
					Type:   "custom",
				}).Post("/cast")

				if err != nil {
					log.Error(err)
				} else {
					log.Infof("Played sound: %s, response code: %d", msg.Payload, resp.StatusCode())
				}
			}
		}
	},
}

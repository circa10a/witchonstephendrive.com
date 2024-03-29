# witchonstephendrive.com 🧹

![Build Status](https://github.com/circa10a/witchonstephendrive.com/workflows/build-docker-images/badge.svg)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/circa10a/witchonstephendrive.com)](https://pkg.go.dev/github.com/circa10a/witchonstephendrive.com?tab=overview)
[![Go Report Card](https://goreportcard.com/badge/github.com/circa10a/witchonstephendrive.com)](https://goreportcard.com/report/github.com/circa10a/witchonstephendrive.com)

[![demo](https://yt-embed.herokuapp.com/embed?v=UTl32JWIu6o)](https://youtu.be/UTl32JWIu6o "demo")

A home automation project to control hue lights for Halloween <img src="https://raw.githubusercontent.com/egonelbre/gophers/10cc13c5e29555ec23f689dc985c157a8d4692ab/vector/fairy-tale/witch-too-much-candy.svg" align="right" width="20%" height="20%"/>

- [witchonstephendrive.com](#witchonstephendrivecom---)
  - [Why](#why)
  - [What does it do](#what-does-it-do)
  - [How does it work](#how-does-it-work)
  - [Usage](#usage)
    - [Configuration](#configuration)
    - [Go](#go)
    - [Docker](#docker)
      - [Deploy](#deploy)
      - [Build](#build)
    - [Terraform module](#terraform-module)
    - [Endpoints](#endpoints)
  - [Get colors](#get-colors)
  - [Example color change request](#example-color-change-request)
  - [Get sounds](#get-sounds)
  - [Example sound play request](#example-sound-play-request)
  - [Turn lights on/off](#turn-lights-on/off)
  - [Add new colors](#add-new-colors)
  - [Add new sounds](#add-new-sounds)
  - [Troubleshooting](#troubleshooting)

## Why

- I love Halloween 🎃
- Side projects are the best
- ~~Given how hectic 2020 is, I believe kids won't have a normal Halloween. I want try to produce more smiles from the kids in my neighborhood.~~
- Let's give people a reason to smile every year.

## What does it do

It allows anyone to change the color of the lighting behind the witch silhouette curtain

### Outdoor Pictures

<p float="left">
  <img src="/images/outdoor_green.jpg" width="45%" height="45%"/>
  <img src="/images/outdoor_blue.jpg" width="45%" height="45%"/>
<p/>

### Site Preview

<img src="/images/site_preview.png" width="25%" height="25%"/>

## How does it work

1. Uses [Caddy](https://github.com/caddyserver/caddy) as a reverse proxy to the `witch` app for TLS termination([let's encrypt](https://letsencrypt.org/)).
2. The `witch` app is a Go backend powered by [echo](https://echo.labstack.com/) that serves a vanilla html/css/js front end and has a `/color/:color` route.
3. Once a `/color/:color` route is hit via a `POST` request, the `witch` app uses the [huego](https://github.com/amimof/huego) library for manipulating the state of the philips hue multicolor bulbs. The hue bridge endpoint on your network is automatically discovered.
4. When a `/sound/:sound` route is hit via a `POST` request, the `witch` app writes to an in-memory queue which will then process sounds to play by calling [home assistant](https://www.home-assistant.io/) to play pre-configured halloween sounds through connected google assistant speakers. The reason for the queue is to ensure all sounds are played and do not get interrupted.

## Usage

### Configuration

|                                         |                                                                                                       |           |                    |
|-----------------------------------------|-------------------------------------------------------------------------------------------------------|-----------|--------------------|
| Environment Variable                    | Description                                                                                           | Required  | Default            |
| `WITCH_API_BASE_URL`                    | Base URL for all interactive POST requests                                                            | `false`   | `/api/v1`          |
| `WITCH_GEOFENCING_ENABLED`              | Enable Client IP geofencing enforcing users to be in close proximity. Requires IP Stack API Token     | `false`   | `false`            |
| `WITCH_GEOFENCING_IPBASE_TOKEN`         | [IPBase.com API Token](https://ipbase.com/) to lookup client coordinates                              | `false`   | `""`               |
| `WITCH_GEOFENCING_RADIUS`               | Radius of goefence in kilometers. See [go-geofence](https://github.com/circa10a/go-geofence)          | `false`   | `0.5`              |
| `WITCH_HOME_ASSISTANT_API_TOKEN`        | Home assistant API token to play `/local/<sound>.mp3` files                                           | `false`   | `""`               |
| `WITCH_HOME_ASSISTANT_ENTITY_ID`        | **Sounds only enabled if this is configured**. Name of home assistant speaker(`media_player.speaker)` | `false`   | `""`               |
| `WITCH_HOME_ASSISTANT_HOST`             | Address of home assistant                                                                             | `false`   | `http://127.0.0.1` |
| `WITCH_HOME_ASSISTANT_PORT`             | Listening port of home assistant                                                                      | `false`   | `8123`             |
| `WITCH_HUE_DEFAULT_COLORS`              | Map of light ID/default color to set at configured time. Ex. `var="8:cyan,9:pink"`                    | `false`   | `""`               |
| `WITCH_HUE_DEFAULT_COLORS_ENABLED`      | Enables scheduler to set default colors or not                                                        | `false`   | `false`            |
| `WITCH_HUE_DEFAULT_COLORS_START`        | Local time to set default colors at. Think of this as a nightly "reset"                               | `false`   | `22`               |
| `WITCH_HUE_TOKEN`                       | Philips Hue API Token                                                                                 | `true`    | None               |
| `WITCH_HUE_LIGHTS`                      | Light ID's to change color of. Example(export HUE_LIGHTS="1,2,3")                                     | `true`    | `[]`               |
| `WITCH_HUE_LIGHTS_SCHEDULE_ENABLED`     | Enables start/end times for turning lights on/off. Will also set to default colors if enabled         | `false`   | `false`            |
| `WITCH_HUE_LIGHTS_START`                | Local time to turn on configured lights                                                               | `false`   | `18`               |
| `WITCH_HUE_LIGHTS_END`                  | Local time to turn off configured lights                                                              | `false`   | `7`                |
| `WITCH_LOG_LEVEL`                       | [Logrus](https://github.com/sirupsen/logrus) log level                                                | `false`   | `info`             |
| `WITCH_METRICS_ENABLED`                 | Enables prometheus metrics on `/metrics`                                                              | `false`   | `true`             |
| `WITCH_PORT`                            | Port for web server to listen on                                                                      | `false`   | `8080`             |
| `WITCH_SHOW_BANNER`                     | Displays Happy Halloween banner on startup                                                            | `false`   | `false`            |
| `WITCH_SOUND_QUIET_TIME_ENABLED`        | Enables quiet time functionality during configured hours                                              | `false`   | `true`             |
| `WITCH_SOUND_QUIET_TIME_START`          | Local time to ensure sounds are not played after this hour                                            | `false`   | `22`               |
| `WITCH_SOUND_QUIET_TIME_END`            | Local time to ensure sounds are not played before this hour                                           | `false`   | `07`               |
| `WITCH_SOUND_QUEUE_CAPACITY`            | Maximum depth of sound queue. This is to ensure no spam/long backlog                                  | `false`   | `1`                |
| `WITCH_UI_ENABLED`                      | Enables hosting of UI/static assets on `/`                                                            | `false`   | `true`             |

### Go

```bash
go generate ./...
go build -o witch .
export WITCH_HUE_TOKEN=<YOUR_TOKEN>; export WITCH_HUE_LIGHTS="1,2,3"
./witch
```

### Docker

#### Deploy

> Follow the [Home Assistant docs](https://www.home-assistant.io/) to setup speakers to support sounds by going to http://localhost:8123/
> You will also need to [create an API token in home assistant](https://developers.home-assistant.io/docs/auth_api/#long-lived-access-token) and set WITCH_HOME_ASSISTANT_API_TOKEN environment variable. Sample config in [provided example .env file](.env)

```bash
docker-compose up -d
```

#### Build

```bash
# Auto determine CPU arch
make build-docker
# ARM64
make build-docker-arm64
# ARMv7
make build-docker-armv7
```

### Terraform module

Colors + sounds are not mutually exclusive, you can pass either just a color, just a sound, or both.

```hcl
module "witchonstephendrive" {
  source       = "github.com/circa10a/witchonstephendrive.com//terraform"
  api_base_url = "https://witchonstephendrive.com/api/v1"
  color        = "purple"
  sound        = "stranger-things"
}

output "color_change_response" {
  value = module.witchonstephendrive.color_change_response
}

output "supported_colors" {
  value = module.witchonstephendrive.supported_colors
}

output "sound_play_response" {
  value = module.witchonstephendrive.sound_play_response
}

output "supported_sounds" {
  value = module.witchonstephendrive.supported_sounds
}
```

### Endpoints

> Rate limiting performed by [this caddy plugin](https://github.com/mholt/caddy-ratelimit)

|                         |                                                 |        |              |                     |
|-------------------------|-------------------------------------------------|--------|--------------|---------------------|
| Route                   | Description                                     | Method | Rate Limited | Limit               |
| `/`                     | Serves static content embedded from `./web`     | `GET`  | No           | N/A                 |
| `/api/v1/colors`        | Get supported colors to change to               | `GET`  | No           | N/A                 |
| `/api/v1/color/:color`  | Changes color of hue lights                     | `POST` | Yes          | 10 requests per 10s |
| `/api/v1/sounds`        | Get supported sounds to play                    | `GET`  | No           | N/A                 |
| `/api/v1/sound/:sound`  | Plays sound through configured speaker          | `POST` | Yes          | 10 requests per 10s |
| `/api/v1/lights/:state` | Changes state of configured lights(on/off)      | `POST` | Yes          | 10 requests per 10s |
| `/metrics`              | Serves prometheus metrics using echo middleware | `GET`  | No           | N/A                 |
| `/swagger/index.html`   | Swagger API documentation                       | `GET`  | No           | N/A                 |

## Get colors

```bash
curl -X POST http://localhost:8080/api/v1/colors
```

## Example color change request

```bash
curl -X POST http://localhost:8080/api/v1/color/red
```

## Get sounds

```bash
curl -X POST http://localhost:8080/api/v1/sounds
```

## Example sound play request

```bash
curl -X POST http://localhost:8080/api/v1/sound/werewolf
```

## Turn lights on/off

```bash
# on
curl -X POST http://localhost:8080/api/v1/lights/on
# off
curl -X POST http://localhost:8080/api/v1/lights/off
```

## Add new colors

To add new colors, make a new entry in `./controllers/colors/colors.go`

## Add new sounds

To add new sounds, simply drop a new `.mp3` file in the `./sounds` directory. This is needed to add to the list of supported sounds and will be built into home assistant docker image.

## Troubleshooting

For debug logs:

```bash
export WITCH_LOG_LEVEL=debug
make run
```

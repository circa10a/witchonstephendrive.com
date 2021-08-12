# witchonstephendrive.com 🧹

![Build Status](https://github.com/circa10a/witchonstephendrive.com/workflows/build-docker-images/badge.svg)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/circa10a/witchonstephendrive.com)](https://pkg.go.dev/github.com/circa10a/witchonstephendrive.com?tab=overview)
[![Go Report Card](https://goreportcard.com/badge/github.com/circa10a/witchonstephendrive.com)](https://goreportcard.com/report/github.com/circa10a/witchonstephendrive.com)

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

## Why

- I love Halloween 🎃
- Side projects are the best
- ~~Given how hectic 2020 is, I believe kids won't have a normal Halloween. I want try to produce more smiles from the kids in my neighborhood.~~
- Let's give people a reason to smile every year.

## What does it do

It allows anyone to change the color of the lighting behind the witch silhouette curtain

Here's what the front of my house looks like:

<img src="https://i.imgur.com/hQE6u6h.jpg" width="50%" height="50%"/>

<img src="https://i.imgur.com/Qj296rO.jpg" width="50%" height="50%"/>

Here's what [witchonstephendrive.com](https://witchonstephendrive.com) looks like(with some sweet ghost animations):

<img src="https://i.imgur.com/BSg32cA.png" width="35%" height="35%"/>

## How does it work

1. Uses [Caddy](https://github.com/caddyserver/caddy) as a reverse proxy to the `witch` app for TLS termination([let's encrypt](https://letsencrypt.org/)).
2. The `witch` app is a Go backend powered by [echo](https://echo.labstack.com/) that serves a vanilla html/css/js front end and has a `/color/:color` route.
3. Once a `/color/:color` route is hit via a `POST` request, the `witch` app uses the [huego](https://github.com/amimof/huego) library for manipulating the state of the philips hue multicolor bulbs. The hue bridge endpoint on your network is automatically discovered.
4. When a `/sound/:sound` route is hit via a `POST` request, the `witch` app writes to an in-memory queue which will then process sounds to play by calling the [assistant-relay](https://assistantrelay.com) to play pre-configured halloween sounds through connected nest speakers. The reason for the queue is to ensure all sounds are played and do not overlap.

## Usage

### Configuration

|                             |                                                                          |                                     |           |                    |
|-----------------------------|--------------------------------------------------------------------------|-------------------------------------|-----------|--------------------|
| Name                        | Description                                                              | Environment Variable                | Required  | Default            |
| API_BASE_URL                | Base URL for all interactive POST requests                               | `WITCH_API_BASE_URL`                | `false`   | `/api/v1`          |
| API_ENABLED                 | Enables swagger docs + REST API routes                                   | `WITCH_API_ENABLED`                 | `false`   | `true`             |
| ASSISTANT_DEVICE            | Name of google assistant speaker to play sounds on                       | `WITCH_ASSISTANT_DEVICE`            | `true`    | None               |
| ASSISTANT_RELAY_HOST        | Address of the google assistant relay                                    | `WITCH_ASSISTANT_RELAY_HOST`        | `false`   | `http://127.0.0.1` |
| ASSISTANT_RELAY_PORT        | Listening port of the google assistant relay                             | `WITCH_ASSISTANT_RELAY_PORT`        | `false`   | `3000`             |
| HUE_TOKEN                   | Philips Hue API User/Token                                               | `WITCH_HUE_TOKEN`                    | `true`    | None              |
| HUE_BRIDGE_REFRESH_INTERVAL | How many seconds to wait before rediscovering hue bridge config/ip       | `WITCH_HUE_BRIDGE_REFRESH_INTERVAL` | `true`    | `21600`            |
| HUE_LIGHTS                  | Light ID's to change color of. Example(export HUE_LIGHTS="1,2,3")        | `WITCH_HUE_LIGHTS`                  | `true`    | None               |
| METRICS_ENABLED             | Enables prometheus metrics on `/metrics`                                 | `WITCH_METRICS_ENABLED`             | `false`   | `true`             |
| PORT                        | Port for web server to listen on                                         | `WITCH_PORT`                        | `false`   | `8080`             |
| SOUND_QUIET_TIME_START      | Local time to ensure sounds are not played after this hour               | `WITCH_SOUND_QUIET_TIME_START`      | `false`   | `22`               |
| SOUND_QUIET_TIME_END        | Local time to ensure sounds are not played before this hour              | `WITCH_SOUND_QUIET_TIME_END`        | `false`   | `07`               |
| SOUND_QUEUE_CAPACITY        | Maxiumum depth of soung queue. This is to ensure no spam/long backlog    | `WITCH_SOUND_QUEUE_CAPACITY`        | `false`   | `3`                |
| SOUND_QUEUE_POLL_INTERVAL   | How many seconds to wait between checking sound queue to play sound msgs | `WITCH_SOUND_QUEUE_POLL_INTERVAL`   | `false`   | `1`                |
| UI_ENABLED                  | Enables hosting of UI/static assets on `/`                               | `WITCH_UI_ENABLED`                  | `false`   | `true`             |

### Go

```bash
go build -o witch .
export WITCH_HUE_TOKEN=<YOUR_TOKEN>; export WITCH_HUE_LIGHTS="1,2,3"; export WITCH_ASSISTANT_DEVICE="My speaker"
./witch
```

### Docker

#### Deploy

> Be sure to update values in `.env` to be consumed by `docker-compose.yml`

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

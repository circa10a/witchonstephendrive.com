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

|                             |                                                                          |                               |           |             |
|-----------------------------|--------------------------------------------------------------------------|-------------------------------|-----------|-------------|
| Name                        | Description                                                              | Environment Variable          | Required  | Default     |
| PORT                        | Port for web server to listen on                                         | `PORT`                        | `false`   | `8080`      |
| API_BASE_URL                | Base URL for all interactive POST requests                               | `API_BASE_URL`                | `false`   | `/api/v1`   |
| HUE_USER                    | Philips Hue API User/Token                                               | `HUE_USER`                    | `true`    | None        |
| HUE_LIGHTS                  | Light ID's to change color of. Example(export HUE_LIGHTS="1,2,3")        | `HUE_LIGHTS`                  | `true`    | None        |
| HUE_BRIDGE_REFRESH_INTERVAL | How many seconds to wait before rediscovering hue bridge config/ip       | `HUE_BRIDGE_REFRESH_INTERVAL` | `true`    | `21600`     |
| METRICS                     | Enables prometheus metrics on `/metrics`(unset for false)                | `METRICS`                     | `false`   | `true`      |
| ASSISTANT_DEVICE            | Name of google assistant speaker to play sounds on                       | `ASSISTANT_DEVICE`            | `true`    | None        |
| ASSISTANT_RELAY_HOST        | Address of the google assistant relay                                    | `ASSISTANT_RELAY_HOST`        | `false`   | `127.0.0.1` |
| ASSISTANT_RELAY_PORT        | Listening port of the google assistant relay                             | `ASSISTANT_RELAY_PORT`        | `false`   | `3000`      |
| SOUND_QUIET_TIME_START      | Local time to ensure sounds are not played after this hour               | `SOUND_QUIET_TIME_START`      | `false`   | `22`        |
| SOUND_QUIET_TIME_END        | Local time to ensure sounds are not played before this hour              | `SOUND_QUIET_TIME_END`        | `false`   | `07`        |
| SOUND_QUEUE_CAPACITY        | Maxiumum depth of soung queue. This is to ensure no spam/long backlog    | `SOUND_QUEUE_CAPACITY`        | `false`   | `3`         |
| SOUND_QUEUE_POLL_INTERVAL   | How many seconds to wait between checking sound queue to play sound msgs | `SOUND_QUEUE_POLL_INTERVAL`   | `false`   | `1`         |

### Go

```bash
go build -o witch .
export HUE_USER=<YOUR_TOKEN>; export HUE_LIGHTS="1,2,3"; export ASSISTANT_DEVICE="My speaker"
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

### Endpoints

> Rate limiting performed by [this caddy plugin](https://github.com/mholt/caddy-ratelimit)

|                         |                                                 |        |              |                 |
|-------------------------|-------------------------------------------------|--------|--------------|-----------------|
| Route                   | Description                                     | Method | Rate Limited | Limit           |
| `/`                     | Serves static content in embedded from `./web`  | `GET`  | No           | N/A             |
| `/api/v1/colors`        | Get supported colors to change to               | `GET`  | No           | N/A             |
| `/api/v1/color/:color`  | Changes color of hue lights                     | `POST` | Yes          | 10 req. per 10s |
| `/api/v1/sounds`        | Get support sounds to play                      | `GET`  | No           | N/A             |
| `/api/v1/sound/:sound`  | Changes color of hue lights                     | `POST` | Yes          | 10 req. per 10s |
| `/api/v1/lights/:state` | Changes state of configured lights(on/off)      | `POST` | Yes          | 10 req. per 10s |
| `/metrics`              | Serves prometheus metrics using echo middleware | `GET`  | No           | N/A             |
| `/swagger/index.html`   | Swagger API documentation                       | `GET`  | No           | N/A             |

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

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

1. Uses [Caddy](https://github.com/caddyserver/caddy) as a reverse proxy to the `witch` server app for TLS termination([let's encrypt](https://letsencrypt.org/)).
2. The `witch` server app is a Go backend powered by [echo](https://echo.labstack.com/) that serves a vanilla html/css/js front end and has a `/color/:color` route.
3. Once a `/color/:color` route is hit via a `POST` request, the `witch` app uses the [huego](https://github.com/amimof/huego) library for manipulating the state of the philips hue multicolor bulbs. The hue bridge endpoint on your network is automatically discovered.
4. When a `/sound/:sound` route is hit via a `POST` request, the `witch` server app drops a message into a [Redis](https://redis.io/) pub/sub queue. Once a message is placed in the queue, the `witch` client application reads the message and talks to the local [assistant-relay](https://assistantrelay.com) to play pre-configured halloween sounds through connected nest speakers. The reason for the queue is that if sound casts overlap, some sounds may not play. The `witch` client processes sound messages 1 at a time to ensure everything is played.

## Usage

```log
❯ ./witch
View docs at https://github.com/circa10a/witchonstephendrive.com

Usage:
  witch [command]

Available Commands:
  client      Start witch client
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  server      Start witch server

Flags:
  -h, --help   help for witch

Use "witch [command] --help" for more information about a command.
```

### Configuration

#### Witch server

|                        |                                                                       |                          |           |                    |
|------------------------|-----------------------------------------------------------------------|--------------------------|-----------|--------------------|
| Name                   | Description                                                           | Environment Variable     | Required  | Default            |
| PORT                   | Port for web server to listen on                                      | `PORT`                   | `false`   | `8080`             |
| HUE_USER               | Philips Hue API User/Token                                            | `HUE_USER`               | `true`    | None               |
| HUE_LIGHTS             | Light ID's to change color of. Example(export HUE_LIGHTS="1,2,3")     | `HUE_LIGHTS`             | `true`    | None               |
| METRICS                | Enables prometheus metrics on `/metrics`(unset for false)             | `METRICS`                | `false`   | `true`             |
| ASSISTANT_DEVICE       | Name of google assistant speaker to play sounds on                    | `ASSISTANT_DEVICE`       | `true`    | None               |
| ASSISTANT_RELAY_HOST   | Address of the google assistant relay                                 | `ASSISTANT_RELAY_HOST`   | `false`   | `http://127.0.0.1` |
| ASSISTANT_RELAY_PORT   | Listening port of the google assistant relay                          | `ASSISTANT_RELAY_PORT`   | `false`   | `3000`             |
| REDIS_HOST             | Address of the redis server                                           | `REDIS_HOST`             | `false`   | `127.0.0.1`        |
| REDIS_PORT             | Port the redis server is configured to listen on                      | `REDIS_PORT`             | `false`   | `6379`             |
| REDIS_PASSWORD         | Password used to authenticate against the redis server                | `REDIS_PASSWORD`         | `false`   | `""`               |
| REDIS_CHANNEL          | Redis channel used for the witch server to place sound messages in    | `REDIS_CHANNEL`          | `false`   | `sounds`           |

#### Witch client

|                        |                                                                       |                          |           |                    |
|------------------------|-----------------------------------------------------------------------|--------------------------|-----------|--------------------|
| Name                   | Description                                                           | Environment Variable     | Required  | Default            |
| ASSISTANT_DEVICE       | Name of google assistant speaker to play sounds on                    | `ASSISTANT_DEVICE`       | `true`    | None               |
| ASSISTANT_RELAY_HOST   | Address of the google assistant relay                                 | `ASSISTANT_RELAY_HOST`   | `false`   | `http://127.0.0.1` |
| ASSISTANT_RELAY_PORT   | Listening port of the google assistant relay                          | `ASSISTANT_RELAY_PORT`   | `false`   | `3000`             |
| REDIS_HOST             | Address of the redis server                                           | `REDIS_HOST`             | `false`   | `127.0.0.1`        |
| REDIS_PORT             | Port the redis server is configured to listen on                      | `REDIS_PORT`             | `false`   | `6379`             |
| REDIS_PASSWORD         | Password used to authenticate against the redis server                | `REDIS_PASSWORD`         | `false`   | `""`               |
| REDIS_CHANNEL          | Redis channel used for the witch server to place sound messages in    | `REDIS_CHANNEL`          | `false`   | `sounds`           |

### Go

```bash
go build -o witch .
export HUE_USER=<YOUR_TOKEN>; export HUE_LIGHTS="1,2,3"; export ASSISTANT_DEVICE="My speaker"
# Start server
./witch server
# Start client
./witch client
```

### Docker

> Be sure to update values in `.env`

```bash
docker-compose up -d
```

### Endpoints

> Rate limiting performed by [this caddy plugin](https://github.com/mholt/caddy-ratelimit)

|                       |                                                  |        |              |                |
|-----------------------|--------------------------------------------------|--------|--------------|----------------|
| Route                 | Description                                      | Method | Rate Limited | Limit          |
| `/`                   | Serves static content in embedded from `./web`   | `GET`  | No           | N/A            |
| `/colors`             | Get supported colors to change to                | `GET`  | No           | N/A            |
| `/color/:color`       | Changes color of hue lights                      | `POST` | Yes          | 5 req. per 10s |
| `/sounds`             | Get support sounds to play                       | `GET`  | No           | N/A            |
| `/sound/:sound`       | Changes color of hue lights                      | `POST` | Yes          | 5 req. per 10s |
| `/lights/:state`      | Changes state of configured lights(on/off)       | `POST` | Yes          | 5 req. per 10s |
| `/metrics`            | Serves prometheus metrics using echo middleware] | `GET`  | No           | N/A            |
| `/swagger/index.html` | Swagger API documentation                        | `GET`  | No           | N/A            |

## Get colors

```bash
curl -X POST http://localhost:8080/colors
```

## Example color change request

```bash
curl -X POST http://localhost:8080/color/red
```

## Get sounds

```bash
curl -X POST http://localhost:8080/sounds
```

## Example sound play request

```bash
curl -X POST http://localhost:8080/sound/werewolf
```

## Turn lights on/off

```bash
# on
curl -X POST http://localhost:8080/lights/on
# off
curl -X POST http://localhost:8080/lights/off
```

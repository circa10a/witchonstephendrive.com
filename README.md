# witchonstephendrive.com 🧹

[![PkgGoDev](https://pkg.go.dev/badge/github.com/circa10a/witchonstephendrive.com)](https://pkg.go.dev/github.com/circa10a/witchonstephendrive.com?tab=overview)
[![Go Report Card](https://goreportcard.com/badge/github.com/circa10a/witchonstephendrive.com)](https://goreportcard.com/report/github.com/circa10a/witchonstephendrive.com)

A home automation project to control hue lights for Halloween <img src="https://raw.githubusercontent.com/egonelbre/gophers/10cc13c5e29555ec23f689dc985c157a8d4692ab/vector/fairy-tale/witch-too-much-candy.svg" align="right" width="20%" height="20%"/>

- [witchonstephendrive.com](#witchonstephendrivecom---)
  * [Why](#why)
  * [What does it do](#what-does-it-do)
  * [How does it work](#how-does-it-work)
  * [Usage](#usage)
    + [Configuration](#configuration)
    + [Go](#go)
    + [Docker](#docker)
    + [Endpoints](#endpoints)
  * [Example color change request](#example-color-change-request)

## Why

- I love Halloween 🎃
- Side projects are the best
- Given how hectic 2020 is, I believe kids won't have a normal Halloween. I want try to produce more smiles from the kids in my neighborhood.

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
4. When a `/sound/:sound` route is hit via a `POST` request, the `witch` app talks to the local [assistant-relay](https://assistantrelay.com) to play pre-configured halloween sounds through connected nest speakers.

## Usage

### Configuration

|                        |                                                                       |                          |           |               |
|------------------------|-----------------------------------------------------------------------|--------------------------|-----------|---------------|
| Name                   | Description                                                           | Environment Variable     | Required  | Default       |
| PORT                   | Port for web server to listen on                                      | `PORT`                   | `false`   | `8080`        |
| HUE_USER               | Philips Hue API User/Token                                            | `HUE_USER`               | `true`    | None          |
| HUE_LIGHTS             | Light ID's to change color of. Example(export HUE_LIGHTS="1,2,3")     | `HUE_LIGHTS`             | `true`    | None          |
| METRICS                | Enables prometheus metrics on `/metrics`(unset for false)             | `METRICS`                | `false`   | `true`        |
| ASSISTANT_DEVICE       | Name of google assistant speaker to play sounds on                    | `ASSISTANT_RELAY_DEVICE` | `true`    | None          |
| ASSISTANT_RELAY_HOST   | Address of the google assistant relay                                 | `ASSISTANT_RELAY_HOST`   | `false`   | `127.0.0.1`   |
| ASSISTANT_RELAY_PORT   | Listening port of the google assistant relay                          | `ASSISTANT_RELAY_PORT`   | `false`   | `3000`        |

### Go

```bash
go build -o witch .
export HUE_USER=<YOUR_TOKEN>; export HUE_LIGHTS="1,2,3"
./witch
```

### Docker

> Be sure to update `HUE_USER` and `HUE_LIGHTS` in `docker-compose.yml`

```bash
docker-compose up -d
```

### Endpoints

|                       |                                                                                                    |        |
|-----------------------|----------------------------------------------------------------------------------------------------|--------|
| Route                 | Description                                                                                        | Method |
| `/`                   | Serves static content in embedded from `./web`                                                     | `GET`  |
| `/colors`             | Get supported colors to change to                                                                  | `GET`  |
| `/color/:color`       | Changes color of hue lights                                                                        | `POST` |
| `/sounds`             | Get support sounds to play                                                                         | `GET`  |
| `/sound/:sound`       | Changes color of hue lights                                                                        | `POST` |
| `/metrics`            | Serves prometheus metrics using [echo middleware](https://echo.labstack.com/middleware/prometheus) | `GET`  |
| `/swagger/index.html` | Swagger API documentation                                                                          | `GET`  |

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

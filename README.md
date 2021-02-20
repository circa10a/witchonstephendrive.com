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

## Usage

### Configuration

|             |                                                                       |                      |                        |           |               |
|-------------|-----------------------------------------------------------------------|----------------------|------------------------|-----------|---------------|
| Name        | Description                                                           | Environment Variable | Command Line Argument  | Required  | Default       |
| PORT        | Port for web server to listen on                                      | `PORT`               | NONE                   | `false`   | `8080`        |
| HUE_USER    | Philips Hue API User/Token                                            | `HUE_USER`           | `--hue-user`           | `true`    | None          |
| HUE_LIGHTS  | Light ID's to change color of                                         | `HUE_LIGHTS`         | `--hue-lights`         | `true`    | None          |
| METRICS     | Enables prometheus metrics on `/metrics`(unset for false)             | `METRICS`            | `--metrics`            | `false`   | `true`        |

### Go

```bash
go build -o witch .
./witch --hue-user <YOUR_TOKEN> --lights "1 2 3"
```

### Docker

> Be sure to update `HUE_USER` and `HUE_LIGHTS` in `docker-compose.yml`

```bash
docker-compose up -d
```

### Endpoints

|             |                                                                                                    |        |
|-------------|----------------------------------------------------------------------------------------------------|--------|
| Route       | Description                                                                                        | Method |
| `/`         | Serves static content in `./web`                                                                   | `GET`  |
| `/:color`   | Changes color of hue lights                                                                        | `POST` |
| `/metrics`  | Serves prometheus metrics using [echo middleware](https://echo.labstack.com/middleware/prometheus) | `GET`  |
| `/swagger`  | Swagger API documentation                                                                          | `GET`  |

## Example color change request

```bash
curl -X POST http://localhost:8080/color/red
```

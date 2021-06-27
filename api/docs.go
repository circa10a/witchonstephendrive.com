// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package api

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Caleb Lemoine",
            "email": "caleblemoine@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://raw.githubusercontent.com/circa10a/witchonstephendrive.com/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/color/{color}": {
            "post": {
                "description": "Change hue lights to supported color defined in color map",
                "produces": [
                    "application/json"
                ],
                "summary": "Change hue lights color",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Color to change lights to",
                        "name": "color",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.ColorResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.ColorResponse"
                        }
                    }
                }
            }
        },
        "/sound/{sound}": {
            "post": {
                "description": "Play halloween sound supported sound defined in sound map",
                "produces": [
                    "application/json"
                ],
                "summary": "Play sound via assistant relay",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Sound to play",
                        "name": "sound",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.SoundResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.SoundResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "routes.ColorResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "routes.SoundResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1.0",
	Host:        "witchonstephendrive.com",
	BasePath:    "/",
	Schemes:     []string{"https"},
	Title:       "witchonstephendrive.com",
	Description: "Control my lights for Halloween",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}

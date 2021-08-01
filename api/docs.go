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
                            "$ref": "#/definitions/routes.ColorSuccessfulChangeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.ColorFailedChangeResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routes.ColorFailedChangeResponse"
                        }
                    }
                }
            }
        },
        "/colors": {
            "get": {
                "description": "Get list of supported colors",
                "produces": [
                    "application/json"
                ],
                "summary": "Get available colors to change to",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.ColorsListResponse"
                        }
                    }
                }
            }
        },
        "/lights/{state}": {
            "post": {
                "description": "Only supports states of on/off",
                "produces": [
                    "application/json"
                ],
                "summary": "Change state of configured lights",
                "parameters": [
                    {
                        "type": "string",
                        "description": "State to set lights to (on/off)",
                        "name": "state",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.LightStateSuccessfulChangeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.LightStateFailedChangeResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routes.LightStateFailedChangeResponse"
                        }
                    }
                }
            }
        },
        "/sound/{sound}": {
            "post": {
                "description": "Play halloween sound supported in sound list",
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
                            "$ref": "#/definitions/routes.SoundSuccessfulPlayResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.SoundFailedPlayResponse"
                        }
                    }
                }
            }
        },
        "/sounds": {
            "get": {
                "description": "Get list of supported sounds",
                "produces": [
                    "application/json"
                ],
                "summary": "Get available sounds to play",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.SoundsListResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "routes.ColorFailedChangeResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                },
                "supportedColors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "routes.ColorSuccessfulChangeResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "routes.ColorsListResponse": {
            "type": "object",
            "properties": {
                "supportedColors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "routes.LightStateFailedChangeResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "routes.LightStateSuccessfulChangeResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "routes.SoundFailedPlayResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                },
                "supportedSounds": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "routes.SoundSuccessfulPlayResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "routes.SoundsListResponse": {
            "type": "object",
            "properties": {
                "supportedSounds": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
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
	Description: "Control my halloween decorations",
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

// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/info": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get authenticated user's info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Show user info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.userResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Send credentials and get JWT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Authenticate user and receive JWT",
                "parameters": [
                    {
                        "description": "Credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.loginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.authResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrResponse"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "Send user data and register new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register new user",
                "parameters": [
                    {
                        "description": "New user data",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.signupRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/api.authResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.ErrResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "description": "low-level runtime error",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "status": {
                    "description": "user-level status message",
                    "type": "string"
                }
            }
        },
        "api.authResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/api.userResponse"
                }
            }
        },
        "api.loginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "default": "email@example.com"
                },
                "password": {
                    "type": "string",
                    "default": "password"
                }
            }
        },
        "api.signupRequest": {
            "type": "object",
            "required": [
                "email",
                "password",
                "password_confirmation",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "default": "email@example.com"
                },
                "password": {
                    "type": "string",
                    "default": "password"
                },
                "password_confirmation": {
                    "type": "string",
                    "default": "password"
                },
                "username": {
                    "type": "string",
                    "default": "username"
                }
            }
        },
        "api.userResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "WebChat API",
	Description: "WebChat golang backend server.",
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

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
        "/v1/about": {
            "get": {
                "description": "Returns the supported API versions and the internal build nr",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "General"
                ],
                "summary": "Lists general information about the API",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.About"
                        }
                    }
                }
            }
        },
        "/v1/attachments": {
            "get": {
                "description": "List all downloaded attachments",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attachments"
                ],
                "summary": "List all attachments.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/v1/attachments/{attachment}": {
            "get": {
                "description": "Serve the attachment with the given id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attachments"
                ],
                "summary": "Serve Attachment.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Attachment ID",
                        "name": "attachment",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove the attachment with the given id from filesystem.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attachments"
                ],
                "summary": "Remove attachment.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Attachment ID",
                        "name": "attachment",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/v1/configuration": {
            "get": {
                "description": "List the REST API configuration.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "General"
                ],
                "summary": "List the REST API configuration.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Configuration"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Set the REST API configuration.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "General"
                ],
                "summary": "Set the REST API configuration.",
                "parameters": [
                    {
                        "description": "Configuration",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.Configuration"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/v1/groups/{number}": {
            "get": {
                "description": "List all Signal Groups.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups"
                ],
                "summary": "List all Signal Groups.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Registered Phone Number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.GroupEntry"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new Signal Group with the specified members.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups"
                ],
                "summary": "Create a new Signal Group.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Registered Phone Number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/api.CreateGroup"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/v1/groups/{number}/{groupid}": {
            "get": {
                "description": "List a specific Signal Group.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups"
                ],
                "summary": "List a Signal Group.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Registered Phone Number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Group ID",
                        "name": "groupid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.GroupEntry"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the specified Signal Group.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups"
                ],
                "summary": "Delete a Signal Group.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Registered Phone Number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Group Id",
                        "name": "groupid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/v1/groups/{number}/{groupid}/block": {
            "post": {
                "description": "Block the specified Signal Group.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups"
                ],
                "summary": "Block a Signal Group.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Registered Phone Number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Group ID",
                        "name": "groupid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/v1/groups/{number}/{groupid}/join": {
            "post": {
                "description": "Join the specified Signal Group.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups"
                ],
                "summary": "Join a Signal Group.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Registered Phone Number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Group ID",
                        "name": "groupid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/v1/groups/{number}/{groupid}/quit": {
            "post": {
                "description": "Quit the specified Signal Group.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups"
                ],
                "summary": "Quit a Signal Group.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Registered Phone Number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Group ID",
                        "name": "groupid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/v1/health": {
            "get": {
                "description": "Internally used by the docker container to perform the health check.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "General"
                ],
                "summary": "API Health Check",
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/identities/{number}": {
            "get": {
                "description": "List all identities for the given number.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Identities"
                ],
                "summary": "List Identities",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Registered Phone Number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.IdentityEntry"
                            }
                        }
                    }
                }
            }
        },
        "/v1/identities/{number}/trust/{numberToTrust}": {
            "put": {
                "description": "Trust an identity.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Identities"
                ],
                "summary": "Trust Identity",
                "parameters": [
                    {
                        "description": "Input Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.TrustIdentityRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Registered Phone Number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Number To Trust",
                        "name": "numberToTrust",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/profiles/{number}": {
            "put": {
                "description": "Set your name and optional an avatar.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profiles"
                ],
                "summary": "Update Profile.",
                "parameters": [
                    {
                        "description": "Profile Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.UpdateProfileRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Registered Phone Number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/v1/qrcodelink": {
            "get": {
                "description": "Link device and generate QR code",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "Link device and generate QR code.",
                "responses": {
                    "200": {
                        "description": "Image",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/v1/receive/{number}": {
            "get": {
                "description": "Receives Signal Messages from the Signal Network.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Messages"
                ],
                "summary": "Receive Signal Messages.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Registered Phone Number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Receive timeout in seconds (default: 1)",
                        "name": "timeout",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/v1/register/{number}": {
            "post": {
                "description": "Register a phone number with the signal network.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "Register a phone number.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Registered Phone Number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Additional Settings",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/api.RegisterNumberRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/v1/register/{number}/verify/{token}": {
            "post": {
                "description": "Verify a registered phone number with the signal network.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "Verify a registered phone number.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Registered Phone Number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Additional Settings",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/api.VerifyNumberSettings"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Verification Code",
                        "name": "token",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/v1/send": {
            "post": {
                "description": "Send a signal message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Messages"
                ],
                "summary": "Send a signal message.",
                "deprecated": true,
                "parameters": [
                    {
                        "description": "Input Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.SendMessageV1"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/v2/send": {
            "post": {
                "description": "Send a signal message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Messages"
                ],
                "summary": "Send a signal message.",
                "parameters": [
                    {
                        "description": "Input Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.SendMessageV2"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.About": {
            "type": "object",
            "properties": {
                "build": {
                    "type": "integer"
                },
                "versions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "api.Configuration": {
            "type": "object",
            "properties": {
                "logging": {
                    "$ref": "#/definitions/api.LoggingConfiguration"
                }
            }
        },
        "api.CreateGroup": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "api.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "api.GroupEntry": {
            "type": "object",
            "properties": {
                "blocked": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "internal_id": {
                    "type": "string"
                },
                "invite_link": {
                    "type": "string"
                },
                "members": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "pending_invites": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "pending_requests": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "api.IdentityEntry": {
            "type": "object",
            "properties": {
                "added": {
                    "type": "string"
                },
                "fingerprint": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                },
                "safety_number": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "api.LoggingConfiguration": {
            "type": "object",
            "properties": {
                "Level": {
                    "type": "string"
                }
            }
        },
        "api.RegisterNumberRequest": {
            "type": "object",
            "properties": {
                "captcha": {
                    "type": "string"
                },
                "use_voice": {
                    "type": "boolean"
                }
            }
        },
        "api.SendMessageV1": {
            "type": "object",
            "properties": {
                "base64_attachment": {
                    "type": "string"
                },
                "is_group": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                },
                "recipients": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "api.SendMessageV2": {
            "type": "object",
            "properties": {
                "base64_attachments": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "message": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                },
                "recipients": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "api.TrustIdentityRequest": {
            "type": "object",
            "properties": {
                "verified_safety_number": {
                    "type": "string"
                }
            }
        },
        "api.UpdateProfileRequest": {
            "type": "object",
            "properties": {
                "base64_avatar": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "api.VerifyNumberSettings": {
            "type": "object",
            "properties": {
                "pin": {
                    "type": "string"
                }
            }
        }
    },
    "tags": [
        {
            "description": "Some general endpoints.",
            "name": "General"
        },
        {
            "description": "Register and link Devices.",
            "name": "Devices"
        },
        {
            "description": "Create, List and Delete Signal Groups.",
            "name": "Groups"
        },
        {
            "description": "Send and Receive Signal Messages.",
            "name": "Messages"
        },
        {
            "description": "List and Delete Attachments.",
            "name": "Attachments"
        },
        {
            "description": "Update Profile.",
            "name": "Profiles"
        },
        {
            "description": "List and Trust Identities.",
            "name": "Identities"
        }
    ]
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
	Host:        "127.0.0.1:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Signal Cli REST API",
	Description: "This is the Signal Cli REST API documentation.",
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

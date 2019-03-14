// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This API will give access example string functions",
    "title": "Example Strings API",
    "contact": {
      "email": "rjj.work@gmail.com"
    },
    "version": "0.0.1"
  },
  "host": "localhost",
  "basePath": "/v0",
  "paths": {
    "/concat": {
      "get": {
        "security": [],
        "description": "Returns concatenated string of inputs",
        "tags": [
          "example"
        ],
        "summary": "concat",
        "responses": {
          "200": {
            "description": "No problems"
          },
          "500": {
            "description": "Problems."
          }
        }
      }
    }
  },
  "securityDefinitions": {
    "OAuth2": {
      "type": "oauth2",
      "flow": "accessCode",
      "authorizationUrl": "https://example.com/oauth/authorize",
      "tokenUrl": "https://example.com/oauth/token",
      "scopes": {
        "read": "Grants read access"
      }
    }
  },
  "security": [
    {
      "OAuth2": [
        "read"
      ]
    }
  ],
  "tags": [
    {
      "description": "example implementation to compare with other code generators",
      "name": "strings function example"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This API will give access example string functions",
    "title": "Example Strings API",
    "contact": {
      "email": "rjj.work@gmail.com"
    },
    "version": "0.0.1"
  },
  "host": "localhost",
  "basePath": "/v0",
  "paths": {
    "/concat": {
      "get": {
        "security": [],
        "description": "Returns concatenated string of inputs",
        "tags": [
          "example"
        ],
        "summary": "concat",
        "responses": {
          "200": {
            "description": "No problems"
          },
          "500": {
            "description": "Problems."
          }
        }
      }
    }
  },
  "securityDefinitions": {
    "OAuth2": {
      "type": "oauth2",
      "flow": "accessCode",
      "authorizationUrl": "https://example.com/oauth/authorize",
      "tokenUrl": "https://example.com/oauth/token",
      "scopes": {
        "read": "Grants read access"
      }
    }
  },
  "security": [
    {
      "OAuth2": [
        "read"
      ]
    }
  ],
  "tags": [
    {
      "description": "example implementation to compare with other code generators",
      "name": "strings function example"
    }
  ]
}`))
}

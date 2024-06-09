// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/checkout": {
            "post": {
                "description": "Buy tickets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tickets"
                ],
                "summary": "Buy tickets for an event",
                "parameters": [
                    {
                        "description": "Buy tickets data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecase.BuyTicketsInputDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usecase.BuyTicketsOutputDTO"
                        }
                    }
                }
            }
        },
        "/events": {
            "get": {
                "description": "Get all events",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "List all events",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usecase.ListEventsOutputDTO"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Create a new event",
                "parameters": [
                    {
                        "description": "Event data",
                        "name": "event",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecase.CreateEventInputDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/usecase.CreateEventOutputDTO"
                        }
                    }
                }
            }
        },
        "/events/{eventID}": {
            "get": {
                "description": "Get a single event by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Get event by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "eventID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usecase.GetEventOutputDTO"
                        }
                    }
                }
            }
        },
        "/events/{eventID}/spots": {
            "get": {
                "description": "Get all spots for an event",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "spots"
                ],
                "summary": "List spots for an event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "eventID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usecase.ListSpotsOutputDTO"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "usecase.BuyTicketsInputDTO": {
            "type": "object",
            "properties": {
                "card_hash": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "event_id": {
                    "type": "string"
                },
                "spots": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "ticket_type": {
                    "type": "string"
                }
            }
        },
        "usecase.BuyTicketsOutputDTO": {
            "type": "object",
            "properties": {
                "tickets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/usecase.TicketDTO"
                    }
                }
            }
        },
        "usecase.CreateEventInputDTO": {
            "type": "object",
            "properties": {
                "capacity": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "organization": {
                    "type": "string"
                },
                "partner_id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "rating": {
                    "type": "string"
                }
            }
        },
        "usecase.CreateEventOutputDTO": {
            "type": "object",
            "properties": {
                "capacity": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "organization": {
                    "type": "string"
                },
                "partner_id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "rating": {
                    "type": "string"
                }
            }
        },
        "usecase.EventDTO": {
            "type": "object",
            "properties": {
                "capacity": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "organization": {
                    "type": "string"
                },
                "partner_id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "rating": {
                    "type": "string"
                },
                "spots": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/usecase.SpotDTO"
                    }
                },
                "tickets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/usecase.TicketDTO"
                    }
                }
            }
        },
        "usecase.GetEventOutputDTO": {
            "type": "object",
            "properties": {
                "capacity": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "organization": {
                    "type": "string"
                },
                "partner_id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "rating": {
                    "type": "string"
                },
                "tickets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/usecase.TicketDTO"
                    }
                }
            }
        },
        "usecase.ListEventsOutputDTO": {
            "type": "object",
            "properties": {
                "events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/usecase.EventDTO"
                    }
                }
            }
        },
        "usecase.ListSpotsOutputDTO": {
            "type": "object",
            "properties": {
                "event": {
                    "$ref": "#/definitions/usecase.EventDTO"
                },
                "spots": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/usecase.SpotDTO"
                    }
                }
            }
        },
        "usecase.SpotDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "ticket_id": {
                    "type": "string"
                }
            }
        },
        "usecase.TicketDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "spot_id": {
                    "type": "string"
                },
                "ticket_type": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Events API",
	Description:      "This is a server for managing events. Imersão Full Cycle",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

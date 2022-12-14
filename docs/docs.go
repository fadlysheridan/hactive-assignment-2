// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/orders": {
            "get": {
                "description": "Mendapatkan semua order yang telah di generate",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Mendapatkan semua order yang telah di generate",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.FindAllOrderResponse"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Mendaftarkan/membuat satu order",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Mendaftarkan/membuat satu order",
                "parameters": [
                    {
                        "description": "Order",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CreateOrderResponse"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/orders/{id}": {
            "get": {
                "description": "Mendapatkan satu order yang telah dibuat dengan id",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Mendapatkan satu order yang telah dibuat dengan id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.FindOneOrderResponse"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Mengupdate satu Order yang sudah terbuat berdasarkan id",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Mengupdate satu Order yang sudah terbuat berdasarkan id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Order",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.UpdateOrderResponse"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Menghapus Order yang sudah dibuat berdasarkan id",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Menghapus Order yang sudah dibuat berdasarkan id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.DeleteOrderResponse"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.CreateOrderRequest": {
            "type": "object",
            "properties": {
                "customer_name": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/request.Item"
                    }
                }
            }
        },
        "request.Item": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "itemCode": {
                    "type": "string"
                },
                "item_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "request.UpdateOrderRequest": {
            "type": "object",
            "properties": {
                "customer_name": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/request.Item"
                    }
                }
            }
        },
        "response.CreateOrderResponse": {
            "type": "object",
            "properties": {
                "order": {
                    "$ref": "#/definitions/response.Order"
                }
            }
        },
        "response.DeleteOrderResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "response.FindAllOrderResponse": {
            "type": "object",
            "properties": {
                "order": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.Order"
                    }
                }
            }
        },
        "response.FindOneOrderResponse": {
            "type": "object",
            "properties": {
                "order": {
                    "$ref": "#/definitions/response.Order"
                }
            }
        },
        "response.Item": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "item_code": {
                    "type": "string"
                },
                "item_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "response.Order": {
            "type": "object",
            "properties": {
                "customer_name": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.Item"
                    }
                },
                "order_id": {
                    "type": "integer"
                },
                "orderedAt": {
                    "type": "string"
                }
            }
        },
        "response.UpdateOrderResponse": {
            "type": "object",
            "properties": {
                "order": {
                    "$ref": "#/definitions/response.Order"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

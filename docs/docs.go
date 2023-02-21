// Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/banks": {
            "get": {
                "tags": [
                    "bank"
                ],
                "summary": "Get all banks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Bank"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "bank"
                ],
                "summary": "Create new bank",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/banks/id": {
            "put": {
                "tags": [
                    "bank"
                ],
                "summary": "Update bank by ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Bank"
                        }
                    }
                }
            }
        },
        "/api/banks/{id}": {
            "get": {
                "tags": [
                    "bank"
                ],
                "summary": "Get bank by ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Bank"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "bank"
                ],
                "summary": "Delete bank by ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Bank"
                        }
                    }
                }
            }
        },
        "/api/customers": {
            "get": {
                "tags": [
                    "customer"
                ],
                "summary": "Get all customers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.CustomerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Customer"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "customer"
                ],
                "summary": "Create new customer",
                "parameters": [
                    {
                        "description": "Customer",
                        "name": "string",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CustomerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CustomerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/api/customers/id": {
            "put": {
                "tags": [
                    "customer"
                ],
                "summary": "Update customer by ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Customer"
                        }
                    }
                }
            }
        },
        "/api/customers/{id}": {
            "get": {
                "tags": [
                    "customer"
                ],
                "summary": "Get customer by ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Customer"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "customer"
                ],
                "summary": "Delete customer by ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Customer"
                        }
                    }
                }
            }
        },
        "/api/merchants": {
            "get": {
                "tags": [
                    "merchant"
                ],
                "summary": "Get all merchants",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Merchant"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "merchant"
                ],
                "summary": "Create new merchant",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/merchants/id": {
            "put": {
                "tags": [
                    "merchant"
                ],
                "summary": "Update merchant by ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Merchant"
                        }
                    }
                }
            }
        },
        "/api/merchants/{id}": {
            "get": {
                "tags": [
                    "merchant"
                ],
                "summary": "Get merchant by ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Merchant"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "merchant"
                ],
                "summary": "Delete merchant by ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Merchant"
                        }
                    }
                }
            }
        },
        "/api/payments": {
            "get": {
                "tags": [
                    "payment"
                ],
                "summary": "Get all payments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Payment"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "payment"
                ],
                "summary": "Create new payment",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/payments/{id}": {
            "get": {
                "tags": [
                    "payment"
                ],
                "summary": "Get payment by ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Payment"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Bank": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "string"
                },
                "bank_account_number": {
                    "type": "string"
                },
                "bank_id": {
                    "type": "string"
                }
            }
        },
        "model.Customer": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "model.CustomerRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "user_id": {
                    "type": "string",
                    "example": "1"
                }
            }
        },
        "model.CustomerResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "model.Merchant": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "merchant_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "model.Payment": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "string"
                },
                "bank_account_number": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "payment_id": {
                    "type": "string"
                },
                "receiver_id": {
                    "type": "string"
                },
                "sender_id": {
                    "type": "string"
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
	Title:            "Simple Payment API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

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
        "/admin/promote/{find}": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Promote user admin by id or email",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Promote user admin (You will need send a JWT token of an administration user in authorization header, you can get it in the login route)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search parameter: e-mail or id",
                        "name": "find",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003ctoken\u003e",
                        "description": "JWT Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user_dtos.UserWithoutSensitiveData"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "JWT Login",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user_dtos.LoginResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user": {
            "post": {
                "description": "create user with e-mail and password if the e-mail doesnt already exists and the password is strong",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "Create user",
                        "name": "tags",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user_dtos.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/user_dtos.UserWithoutSensitiveData"
                        }
                    },
                    "408": {
                        "description": "Request Timeout",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/changeEmail/{find}": {
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Change Email by id or email",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Change Email (You will need send a JWT token in authorization header, you can get it in the login route)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search parameter: e-mail or id",
                        "name": "find",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003ctoken\u003e",
                        "description": "JWT Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user_dtos.UserWithoutSensitiveData"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/changePassword/{find}": {
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Change Password by id or email",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Change Password (You will need send a JWT token in authorization header, you can get it in the login route)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search parameter: e-mail or id",
                        "name": "find",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003ctoken\u003e",
                        "description": "JWT Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user_dtos.UserWithoutSensitiveData"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/{find}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Find user by e-mail or id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Find User (You will need send a JWT token in authorization header, you can get it in the login route)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search parameter: e-mail or id",
                        "name": "find",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003ctoken\u003e",
                        "description": "JWT Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user_dtos.UserWithoutSensitiveData"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete user by e-mail or id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete User (You will need send a JWT token in authorization header, you can get it in the login route)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search parameter: e-mail or id",
                        "name": "find",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003ctoken\u003e",
                        "description": "JWT Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils_interfaces.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "user_dtos.CreateUserRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "fulano@fulano.com"
                },
                "password": {
                    "type": "string",
                    "example": "aaaaaaaA#1"
                }
            }
        },
        "user_dtos.LoginResponse": {
            "type": "object",
            "properties": {
                "jwt": {
                    "type": "string",
                    "example": "randomJwt"
                }
            }
        },
        "user_dtos.UserWithoutSensitiveData": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "fulano@fulano.com"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "utils_interfaces.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "variable error message"
                },
                "statusCode": {
                    "type": "integer",
                    "example": -1
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "API",
	Description:      "Authentication API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

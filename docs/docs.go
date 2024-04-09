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
        "/admin/logs": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Find all logs",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Find All Logs (You will need send a JWT token of a admin user, you can get it in the login route)",
                "parameters": [
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
                            "type": "null"
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
        "/admin/logs/{find}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Find all logs of a user by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Find All Logs of a User (You will need send a JWT token of a admin user, you can get it in the login route)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search parameter: id",
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
                            "type": "null"
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
        },
        "/admin/user": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Find all users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Find All Users (You will need send a JWT token of a admin user, you can get it in the login route)",
                "parameters": [
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
                            "type": "null"
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
        "/admin/user/active/{find}": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Active user by e-mail or id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Active User (You will need send a JWT token of a admin user, you can get it in the login route)",
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
        },
        "/admin/user/inative/{find}": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Inative user by e-mail or id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Inative User (You will need send a JWT token of a admin user, you can get it in the login route)",
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
        },
        "/admin/user/{find}": {
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
                    "admin"
                ],
                "summary": "Find User (You will need send a JWT token of a admin user, you can get it in the login route)",
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
                    "admin"
                ],
                "summary": "Delete User (You will need send a JWT token of a admin user, you can get it in the login route)",
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
        "log_entities.Log": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "ip": {
                    "type": "string"
                },
                "method": {
                    "type": "string"
                },
                "operationCode": {
                    "type": "string"
                },
                "route": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                },
                "timestamp": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
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
                },
                "inactive": {
                    "type": "boolean",
                    "example": false
                },
                "isAdmin": {
                    "type": "boolean",
                    "example": true
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

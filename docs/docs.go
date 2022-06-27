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
        "termsOfService": "https://github.com",
        "contact": {
            "name": "conjurer",
            "url": "https:/github.com/dot123",
            "email": "conjurer888888@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/deleteFile/{id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "FileApi"
                ],
                "summary": "删除文件",
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
                        "description": "成功结果",
                        "schema": {
                            "$ref": "#/definitions/ginx.ResponseData"
                        }
                    },
                    "500": {
                        "description": "失败结果",
                        "schema": {
                            "$ref": "#/definitions/ginx.ResponseFail"
                        }
                    }
                }
            }
        },
        "/system/reloadSystem": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SystemApi"
                ],
                "summary": "重启系统",
                "responses": {
                    "200": {
                        "description": "成功结果",
                        "schema": {
                            "$ref": "#/definitions/ginx.ResponseData"
                        }
                    },
                    "500": {
                        "description": "失败结果",
                        "schema": {
                            "$ref": "#/definitions/ginx.ResponseFail"
                        }
                    }
                }
            }
        },
        "/system/serverInfo": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SystemApi"
                ],
                "summary": "服务器状态",
                "responses": {
                    "200": {
                        "description": "成功结果",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/ginx.ResponseData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/helper.Server"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "失败结果",
                        "schema": {
                            "$ref": "#/definitions/ginx.ResponseFail"
                        }
                    }
                }
            }
        },
        "/uploadFile": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "FileApi"
                ],
                "summary": "上传文件",
                "parameters": [
                    {
                        "type": "file",
                        "description": "file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功结果",
                        "schema": {
                            "$ref": "#/definitions/ginx.ResponseData"
                        }
                    },
                    "500": {
                        "description": "失败结果",
                        "schema": {
                            "$ref": "#/definitions/ginx.ResponseFail"
                        }
                    }
                }
            }
        },
        "/user": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserApi"
                ],
                "summary": "修改用户",
                "parameters": [
                    {
                        "type": "string",
                        "name": "avatar",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "user_type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功结果",
                        "schema": {
                            "$ref": "#/definitions/ginx.ResponseData"
                        }
                    },
                    "500": {
                        "description": "失败结果",
                        "schema": {
                            "$ref": "#/definitions/ginx.ResponseFail"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserApi"
                ],
                "summary": "新建用户",
                "parameters": [
                    {
                        "type": "string",
                        "name": "avatar",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "user_type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功结果",
                        "schema": {
                            "$ref": "#/definitions/ginx.ResponseData"
                        }
                    },
                    "500": {
                        "description": "失败结果",
                        "schema": {
                            "$ref": "#/definitions/ginx.ResponseFail"
                        }
                    }
                }
            }
        },
        "/user/info": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserApi"
                ],
                "summary": "获取用户信息",
                "responses": {
                    "200": {
                        "description": "成功结果",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/ginx.ResponseData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/schema.UserData"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "失败结果",
                        "schema": {
                            "$ref": "#/definitions/ginx.ResponseFail"
                        }
                    }
                }
            }
        },
        "/user/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserApi"
                ],
                "summary": "获取用户列表",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "页",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "数量",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "相似用户名",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功结果",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/ginx.ResponseData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/schema.UserQueryResult"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "失败结果",
                        "schema": {
                            "$ref": "#/definitions/ginx.ResponseFail"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserApi"
                ],
                "summary": "删除用户",
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
                        "description": "成功结果",
                        "schema": {
                            "$ref": "#/definitions/ginx.ResponseData"
                        }
                    },
                    "500": {
                        "description": "失败结果",
                        "schema": {
                            "$ref": "#/definitions/ginx.ResponseFail"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ginx.ResponseData": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "ginx.ResponseFail": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "err": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "helper.Cpu": {
            "type": "object",
            "properties": {
                "cores": {
                    "type": "integer"
                },
                "cpus": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                }
            }
        },
        "helper.Disk": {
            "type": "object",
            "properties": {
                "totalGb": {
                    "type": "integer"
                },
                "totalMb": {
                    "type": "integer"
                },
                "usedGb": {
                    "type": "integer"
                },
                "usedMb": {
                    "type": "integer"
                },
                "usedPercent": {
                    "type": "integer"
                }
            }
        },
        "helper.Os": {
            "type": "object",
            "properties": {
                "compiler": {
                    "type": "string"
                },
                "goVersion": {
                    "type": "string"
                },
                "goos": {
                    "type": "string"
                },
                "numCpu": {
                    "type": "integer"
                },
                "numGoroutine": {
                    "type": "integer"
                }
            }
        },
        "helper.Rrm": {
            "type": "object",
            "properties": {
                "totalMb": {
                    "type": "integer"
                },
                "usedMb": {
                    "type": "integer"
                },
                "usedPercent": {
                    "type": "integer"
                }
            }
        },
        "helper.Server": {
            "type": "object",
            "properties": {
                "cpu": {
                    "$ref": "#/definitions/helper.Cpu"
                },
                "disk": {
                    "$ref": "#/definitions/helper.Disk"
                },
                "os": {
                    "$ref": "#/definitions/helper.Os"
                },
                "ram": {
                    "$ref": "#/definitions/helper.Rrm"
                }
            }
        },
        "schema.User": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_type": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schema.UserData": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "introduction": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "schema.UserQueryResult": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.User"
                    }
                },
                "total": {
                    "type": "integer"
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

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "GameAdmin API",
	Description:      "This is a game management background. you can use the api key `special-key` to test the authorization filters.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
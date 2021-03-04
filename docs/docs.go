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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/test/all": {
            "post": {
                "description": "返回请求数据",
                "consumes": [
                    "application/octet-stream"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "CPTS"
                ],
                "summary": "返回请求数据",
                "parameters": [
                    {
                        "description": "数据流",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {},
                    "500": {}
                }
            }
        },
        "/test/person": {
            "post": {
                "description": "上传人物信息",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CPTS"
                ],
                "summary": "上传人物信息",
                "parameters": [
                    {
                        "description": "任务信息",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/cpts.Person"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/cpts.Person"
                        }
                    },
                    "400": {},
                    "500": {}
                }
            }
        },
        "/test/query": {
            "get": {
                "description": "请求参数编码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CPTS"
                ],
                "summary": "请求参数处理",
                "parameters": [
                    {
                        "type": "string",
                        "description": "名字",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "年龄",
                        "name": "age",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/cpts.QueryRequest"
                        }
                    },
                    "400": {},
                    "500": {}
                }
            }
        },
        "/v1/bucket/{bucket_name}": {
            "post": {
                "description": "create bucket",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bucket"
                ],
                "summary": "CreateBucket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bucket name",
                        "name": "bucket_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "bucket action",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.BucketAction"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AkSk"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete bucket",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bucket"
                ],
                "summary": "DeleteBucket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bucket name",
                        "name": "bucket_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {},
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            },
            "head": {
                "description": "head bucket",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bucket"
                ],
                "summary": "HeadBucket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bucket name",
                        "name": "bucket_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "headers": {
                            "Last-Modified": {
                                "type": "string",
                                "description": "last modify"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "404": {},
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/file/{bucket_name}": {
            "get": {
                "description": "download File",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "file"
                ],
                "summary": "DownloadFile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bucket name",
                        "name": "bucket_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "target path",
                        "name": "path",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "upload file",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "UploadFile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bucket name",
                        "name": "bucket_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "file",
                        "name": "path",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.FileTarget"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete file",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "DeleteFile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bucket name",
                        "name": "bucket_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "target path",
                        "name": "path",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            },
            "head": {
                "description": "head file sign info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "SignFile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bucket name",
                        "name": "bucket_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "target path",
                        "name": "path",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.FileTarget"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "cpts.Person": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "score": {
                    "type": "number"
                }
            }
        },
        "cpts.QueryRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.AkSk": {
            "type": "object",
            "required": [
                "ak",
                "sk"
            ],
            "properties": {
                "ak": {
                    "type": "string"
                },
                "sk": {
                    "type": "string"
                }
            }
        },
        "model.BucketAction": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "integer"
                }
            }
        },
        "model.FileTarget": {
            "type": "object",
            "required": [
                "path"
            ],
            "properties": {
                "path": {
                    "type": "string"
                }
            }
        },
        "response.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "sk",
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
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "obs Swagger API",
	Description: "This is a obs server.",
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

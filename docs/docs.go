// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "soberkoder@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/license/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login/": {
            "post": {
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Post detail for a given id",
                "parameters": [
                    {
                        "description": "Login user",
                        "name": "models.User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/photos/": {
            "get": {
                "description": "Get details of All Photo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photos"
                ],
                "summary": "Get Details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            },
            "post": {
                "description": "Post Photo Detail corresponding to the input id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photos"
                ],
                "summary": "Post detail for a given id",
                "parameters": [
                    {
                        "description": "create photo",
                        "name": "models.Photo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            }
        },
        "/photos/{id}": {
            "put": {
                "description": "Update the photo corresponding to the input id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photos"
                ],
                "summary": "Update Photo identified by the given id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of the commment to be updated",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the photo corresponding to the input id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photos"
                ],
                "summary": "Delete Photo identified by the given id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of the commment to be deleted",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            }
        },
        "/photos/{id}/comments": {
            "get": {
                "description": "Get details of All Comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Get Details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            },
            "post": {
                "description": "Post Comment Detail corresponding to the input id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Post detail for a given id",
                "parameters": [
                    {
                        "description": "create comment",
                        "name": "models.Comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            }
        },
        "/photos/{id}/comments/{id}": {
            "put": {
                "description": "Update the comment corresponding to the input id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Update Comment identified by the given id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of the commment to be updated",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the comment corresponding to the input id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Delete Comment identified by the given id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of the commment to be deleted",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            }
        },
        "/register/": {
            "post": {
                "description": "Post New User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Post detail for a given id",
                "parameters": [
                    {
                        "description": "Register user",
                        "name": "models.User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/socMed/": {
            "get": {
                "description": "Get details of All SocialMedia",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "socialMedia"
                ],
                "summary": "Get Details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            },
            "post": {
                "description": "Post SocialMedia Detail corresponding to the input id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "socialMedia"
                ],
                "summary": "Post detail for a given id",
                "parameters": [
                    {
                        "description": "create socmed",
                        "name": "models.SocialMedia",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            }
        },
        "/socMed/{id}": {
            "put": {
                "description": "Update the comment corresponding to the input id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "socialMedia"
                ],
                "summary": "Update SocialMedia identified by the given id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of the commment to be updated",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the comment corresponding to the input id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "socialMedia"
                ],
                "summary": "Delete SocialMedia identified by the given id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of the commment to be deleted",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Comment": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "photo": {
                    "$ref": "#/definitions/models.Photo"
                },
                "photoID": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.Photo": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "photo_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.SocialMedia": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "social_url": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
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
	Title:            "myGram",
	Description:      "API simpel dengan konsep seperti instagram",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
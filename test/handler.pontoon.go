// Code generated by utrack/pontoon. DO NOT EDIT.
// Source: github.com/utrack/pontoon/test

package test

func (s Handler) OpenAPI() string {
	return `{
    "components": {
      "schemas": {
        "test.iterateEmbedded": {
          "description": "Embed comment\n",
          "properties": {
            "PageToken": {
              "description": "PageToken comment\n",
              "type": "string"
            }
          },
          "type": "object"
        },
        "test.iterateRequest": {
          "allOf": [
            {
              "$ref": "#/components/schemas/test.iterateEmbedded"
            }
          ],
          "description": "Request comment\nRequest line 2\n",
          "properties": {
            "Local": {
              "description": "Local field\n",
              "type": "string"
            },
            "Maps": {
              "additionalProperties": {
                "$ref": "#/components/schemas/test.mapped"
              },
              "type": "object"
            },
            "Recursive": {
              "anyOf": [
                {
                  "type": "null"
                },
                {
                  "$ref": "#/components/schemas/test.iterateRequest"
                }
              ]
            },
            "SliceStrings": {
              "items": {
                "type": "string"
              },
              "nullable": true,
              "type": "array"
            }
          },
          "type": "object"
        },
        "test.mapped": {
          "type": "object"
        },
        "test2.IterateResponse": {
          "properties": {
            "resp": {
              "type": "string"
            }
          },
          "type": "object"
        }
      }
    },
    "info": {
      "title": "github.com/utrack/pontoon/test",
      "version": "1-autogen"
    },
    "openapi": "3.1.0",
    "paths": {
      "/v1/products/iterate": {
        "get": {
          "description": "IterateProducts comment\nIncludes imported package\n",
          "parameters": [
            {
              "description": "PageToken comment\n",
              "in": "query",
              "name": "page_token",
              "schema": {
                "description": "PageToken comment\n",
                "type": "string"
              }
            },
            {
              "description": "Local field\n",
              "in": "query",
              "name": "local",
              "schema": {
                "description": "Local field\n",
                "type": "string"
              }
            }
          ],
          "requestBody": {
            "content": {
              "application/json": {
                "schema": {
                  "anyOf": [
                    {
                      "type": "null"
                    },
                    {
                      "$ref": "#/components/schemas/test.iterateRequest"
                    }
                  ]
                }
              }
            }
          },
          "responses": {
            "200": {
              "content": {
                "application/json": {
                  "schema": {
                    "$ref": "#/components/schemas/test2.IterateResponse"
                  }
                }
              },
              "description": "success"
            },
            "default": {
              "description": ""
            }
          },
          "tags": [
            "test.Handler"
          ]
        }
      },
      "/v1/test/return/interface": {
        "get": {
          "parameters": [
            {
              "description": "PageToken comment\n",
              "in": "query",
              "name": "page_token",
              "schema": {
                "description": "PageToken comment\n",
                "type": "string"
              }
            },
            {
              "description": "Local field\n",
              "in": "query",
              "name": "local",
              "schema": {
                "description": "Local field\n",
                "type": "string"
              }
            }
          ],
          "requestBody": {
            "content": {
              "application/json": {
                "schema": {
                  "anyOf": [
                    {
                      "type": "null"
                    },
                    {
                      "$ref": "#/components/schemas/test.iterateRequest"
                    }
                  ]
                }
              }
            }
          },
          "responses": {
            "200": {
              "content": {
                "application/json": {
                  "schema": {}
                }
              },
              "description": "success"
            },
            "default": {
              "description": ""
            }
          },
          "tags": [
            "test.Handler"
          ]
        }
      },
      "/v1/test/return/interface-any": {
        "get": {
          "parameters": [
            {
              "description": "PageToken comment\n",
              "in": "query",
              "name": "page_token",
              "schema": {
                "description": "PageToken comment\n",
                "type": "string"
              }
            },
            {
              "description": "Local field\n",
              "in": "query",
              "name": "local",
              "schema": {
                "description": "Local field\n",
                "type": "string"
              }
            }
          ],
          "requestBody": {
            "content": {
              "application/json": {
                "schema": {
                  "anyOf": [
                    {
                      "type": "null"
                    },
                    {
                      "$ref": "#/components/schemas/test.iterateRequest"
                    }
                  ]
                }
              }
            }
          },
          "responses": {
            "200": {
              "content": {
                "application/json": {
                  "schema": {}
                }
              },
              "description": "success"
            },
            "default": {
              "description": ""
            }
          },
          "tags": [
            "test.Handler"
          ]
        }
      },
      "/v1/test/return/return-nothing": {
        "get": {
          "parameters": [
            {
              "description": "PageToken comment\n",
              "in": "query",
              "name": "page_token",
              "schema": {
                "description": "PageToken comment\n",
                "type": "string"
              }
            },
            {
              "description": "Local field\n",
              "in": "query",
              "name": "local",
              "schema": {
                "description": "Local field\n",
                "type": "string"
              }
            }
          ],
          "requestBody": {
            "content": {
              "application/json": {
                "schema": {
                  "anyOf": [
                    {
                      "type": "null"
                    },
                    {
                      "$ref": "#/components/schemas/test.iterateRequest"
                    }
                  ]
                }
              }
            }
          },
          "responses": {
            "200": {
              "content": {
                "application/json": {}
              },
              "description": "success"
            },
            "default": {
              "description": ""
            }
          },
          "tags": [
            "test.Handler"
          ]
        }
      }
    }
  }`
}

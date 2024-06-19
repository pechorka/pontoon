// Code generated by utrack/pontoon. DO NOT EDIT.
// Source: github.com/utrack/pontoon/test

package test

func (s Handler) OpenAPI() string {
	return `{
    "components": {
      "schemas": {
        "test.dummyStruct": {
          "properties": {
            "DummyField": {
              "type": "string"
            }
          },
          "type": "object"
        },
        "test.iterateEmbedded": {
          "description": "Has an Embed comment",
          "type": "object"
        },
        "test.iterateRequest": {
          "allOf": [
            {
              "$ref": "#/components/schemas/test.iterateEmbedded"
            }
          ],
          "description": "Request comment\nRequest line 2",
          "properties": {
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
        "test.jsonWithArrayOfStructs": {
          "properties": {
            "Ret": {
              "items": {
                "$ref": "#/components/schemas/test.dummyStruct"
              },
              "nullable": true,
              "type": "array"
            }
          },
          "type": "object"
        },
        "test.jsonWithDirectives": {
          "description": "Describes a JSON-marshaled request with additional 'in' directives.",
          "properties": {
            "RequiredOnly": {
              "type": "string"
            },
            "with_default": {
              "default": "1234",
              "type": "string"
            }
          },
          "required": [
            "with_default",
            "RequiredOnly"
          ],
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
          "description": "Includes imported package",
          "operationId": "v1_products_iterate_get",
          "parameters": [
            {
              "description": "A token for the next page.\nPass an empty page_token if you want to request the first page,\nand use a token from the response as page_token to get the next page.",
              "in": "query",
              "name": "page_token",
              "schema": {
                "description": "A token for the next page.\nPass an empty PageToken if you want to request the first page,\nand use a token from the response as PageToken to get the next page.",
                "type": "string"
              }
            },
            {
              "in": "query",
              "name": "foo",
              "schema": {
                "format": "int64",
                "type": "integer"
              }
            },
            {
              "description": "Describes Some Stuff(tm). Required field.",
              "in": "query",
              "name": "local",
              "required": true,
              "schema": {
                "description": "Describes Some Stuff(tm). Required field.",
                "type": "string"
              }
            },
            {
              "in": "query",
              "name": "local_default",
              "required": true,
              "schema": {
                "default": "foobarbaz",
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
          "summary": "Comment",
          "tags": [
            "test.Handler"
          ]
        },
        "post": {
          "description": "Includes imported package",
          "operationId": "v1_products_iterate_post",
          "parameters": [
            {
              "description": "A token for the next page.\nPass an empty page_token if you want to request the first page,\nand use a token from the response as page_token to get the next page.",
              "in": "query",
              "name": "page_token",
              "schema": {
                "description": "A token for the next page.\nPass an empty PageToken if you want to request the first page,\nand use a token from the response as PageToken to get the next page.",
                "type": "string"
              }
            },
            {
              "in": "query",
              "name": "foo",
              "schema": {
                "format": "int64",
                "type": "integer"
              }
            },
            {
              "description": "Describes Some Stuff(tm). Required field.",
              "in": "query",
              "name": "local",
              "required": true,
              "schema": {
                "description": "Describes Some Stuff(tm). Required field.",
                "type": "string"
              }
            },
            {
              "in": "query",
              "name": "local_default",
              "required": true,
              "schema": {
                "default": "foobarbaz",
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
          "summary": "Comment",
          "tags": [
            "test.Handler"
          ]
        }
      },
      "/v1/products/iterate/create": {
        "post": {
          "description": "Includes imported package",
          "operationId": "v1_products_iterate_create_post",
          "parameters": [
            {
              "description": "A token for the next page.\nPass an empty page_token if you want to request the first page,\nand use a token from the response as page_token to get the next page.",
              "in": "query",
              "name": "page_token",
              "schema": {
                "description": "A token for the next page.\nPass an empty PageToken if you want to request the first page,\nand use a token from the response as PageToken to get the next page.",
                "type": "string"
              }
            },
            {
              "in": "query",
              "name": "foo",
              "schema": {
                "format": "int64",
                "type": "integer"
              }
            },
            {
              "description": "Describes Some Stuff(tm). Required field.",
              "in": "query",
              "name": "local",
              "required": true,
              "schema": {
                "description": "Describes Some Stuff(tm). Required field.",
                "type": "string"
              }
            },
            {
              "in": "query",
              "name": "local_default",
              "required": true,
              "schema": {
                "default": "foobarbaz",
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
          "summary": "Comment",
          "tags": [
            "test.Handler"
          ]
        }
      },
      "/v1/test/request/jsonWithDirective": {
        "get": {
          "operationId": "v1_test_request_jsonwithdirective_get",
          "requestBody": {
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/test.jsonWithDirectives"
                }
              }
            },
            "description": "jsonWithDirectives describes a JSON-marshaled request with additional 'in' directives.\n"
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
      },
      "/v1/test/return/interface": {
        "get": {
          "operationId": "v1_test_return_interface_get",
          "parameters": [
            {
              "description": "A token for the next page.\nPass an empty page_token if you want to request the first page,\nand use a token from the response as page_token to get the next page.",
              "in": "query",
              "name": "page_token",
              "schema": {
                "description": "A token for the next page.\nPass an empty PageToken if you want to request the first page,\nand use a token from the response as PageToken to get the next page.",
                "type": "string"
              }
            },
            {
              "in": "query",
              "name": "foo",
              "schema": {
                "format": "int64",
                "type": "integer"
              }
            },
            {
              "description": "Describes Some Stuff(tm). Required field.",
              "in": "query",
              "name": "local",
              "required": true,
              "schema": {
                "description": "Describes Some Stuff(tm). Required field.",
                "type": "string"
              }
            },
            {
              "in": "query",
              "name": "local_default",
              "required": true,
              "schema": {
                "default": "foobarbaz",
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
          "operationId": "v1_test_return_interface-any_get",
          "parameters": [
            {
              "description": "A token for the next page.\nPass an empty page_token if you want to request the first page,\nand use a token from the response as page_token to get the next page.",
              "in": "query",
              "name": "page_token",
              "schema": {
                "description": "A token for the next page.\nPass an empty PageToken if you want to request the first page,\nand use a token from the response as PageToken to get the next page.",
                "type": "string"
              }
            },
            {
              "in": "query",
              "name": "foo",
              "schema": {
                "format": "int64",
                "type": "integer"
              }
            },
            {
              "description": "Describes Some Stuff(tm). Required field.",
              "in": "query",
              "name": "local",
              "required": true,
              "schema": {
                "description": "Describes Some Stuff(tm). Required field.",
                "type": "string"
              }
            },
            {
              "in": "query",
              "name": "local_default",
              "required": true,
              "schema": {
                "default": "foobarbaz",
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
      "/v1/test/return/map": {
        "get": {
          "operationId": "v1_test_return_map_get",
          "parameters": [
            {
              "description": "A token for the next page.\nPass an empty page_token if you want to request the first page,\nand use a token from the response as page_token to get the next page.",
              "in": "query",
              "name": "page_token",
              "schema": {
                "description": "A token for the next page.\nPass an empty PageToken if you want to request the first page,\nand use a token from the response as PageToken to get the next page.",
                "type": "string"
              }
            },
            {
              "in": "query",
              "name": "foo",
              "schema": {
                "format": "int64",
                "type": "integer"
              }
            },
            {
              "description": "Describes Some Stuff(tm). Required field.",
              "in": "query",
              "name": "local",
              "required": true,
              "schema": {
                "description": "Describes Some Stuff(tm). Required field.",
                "type": "string"
              }
            },
            {
              "in": "query",
              "name": "local_default",
              "required": true,
              "schema": {
                "default": "foobarbaz",
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
                    "additionalProperties": {
                      "$ref": "#/components/schemas/test2.IterateResponse"
                    },
                    "type": "object"
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
      "/v1/test/return/return-nothing": {
        "get": {
          "operationId": "v1_test_return_return-nothing_get",
          "parameters": [
            {
              "description": "A token for the next page.\nPass an empty page_token if you want to request the first page,\nand use a token from the response as page_token to get the next page.",
              "in": "query",
              "name": "page_token",
              "schema": {
                "description": "A token for the next page.\nPass an empty PageToken if you want to request the first page,\nand use a token from the response as PageToken to get the next page.",
                "type": "string"
              }
            },
            {
              "in": "query",
              "name": "foo",
              "schema": {
                "format": "int64",
                "type": "integer"
              }
            },
            {
              "description": "Describes Some Stuff(tm). Required field.",
              "in": "query",
              "name": "local",
              "required": true,
              "schema": {
                "description": "Describes Some Stuff(tm). Required field.",
                "type": "string"
              }
            },
            {
              "in": "query",
              "name": "local_default",
              "required": true,
              "schema": {
                "default": "foobarbaz",
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
      },
      "/v1/test/return/slice": {
        "get": {
          "operationId": "v1_test_return_slice_get",
          "parameters": [
            {
              "description": "A token for the next page.\nPass an empty page_token if you want to request the first page,\nand use a token from the response as page_token to get the next page.",
              "in": "query",
              "name": "page_token",
              "schema": {
                "description": "A token for the next page.\nPass an empty PageToken if you want to request the first page,\nand use a token from the response as PageToken to get the next page.",
                "type": "string"
              }
            },
            {
              "in": "query",
              "name": "foo",
              "schema": {
                "format": "int64",
                "type": "integer"
              }
            },
            {
              "description": "Describes Some Stuff(tm). Required field.",
              "in": "query",
              "name": "local",
              "required": true,
              "schema": {
                "description": "Describes Some Stuff(tm). Required field.",
                "type": "string"
              }
            },
            {
              "in": "query",
              "name": "local_default",
              "required": true,
              "schema": {
                "default": "foobarbaz",
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
                    "items": {
                      "$ref": "#/components/schemas/test2.IterateResponse"
                    },
                    "nullable": true,
                    "type": "array"
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
      "/v1/test/return/slice-in-struct": {
        "get": {
          "operationId": "v1_test_return_slice-in-struct_get",
          "responses": {
            "200": {
              "content": {
                "application/json": {
                  "schema": {
                    "$ref": "#/components/schemas/test.jsonWithArrayOfStructs"
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
      }
    },
    "tags": [
      {
        "description": "Struct comment",
        "name": "test.Handler"
      }
    ]
  }`
}

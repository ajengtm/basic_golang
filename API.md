{ "openapi": "3.0.0", "info": { "title": "Basic_Golang", "description":
"This is Auth APP and Fetch App", "version": "1.0.0" }, "servers": \[ {
"url": "https://virtserver.swaggerhub.com/ajengtm/basicGolang/1.0.0",
"description": "SwaggerHub API Auto Mocking" } \], "tags": \[ { "name":
"Auth-App" }, { "name": "Fetch-App" } \], "paths": { "/v1/auth/healthz":
{ "get": { "tags": \[ "Auth-App" \], "description": "This endpoint will
return a \"It Works\" string and Status Code 200. This endpoint can be
used to verify the app is running.`\n`{=tex}", "responses": { "200": {
"description": "It Works", "content": { "application/json": { "schema":
{ "type": "string", "description": "It Works" } } } } } } },
"/v1/auth/upsert-user": { "post": { "tags": \[ "Auth-App" \],
"description": "This endpoint will create user data that is received
from the request body into the database and will return the 4 characters
password for that user.`\n`{=tex}", "requestBody": { "description":
"User to add", "content": { "application/json": { "schema": {
"$ref": "#/components/schemas/UpsertUser"  }  }  }  },  "responses": {  "201": {  "description": "Created",  "content": {  "application/json": {  "schema": {  "$ref":
"\#/components/schemas/User" } } } } } } }, "/v1/auth/login": { "post":
{ "tags": \[ "Auth-App" \], "description": "This endpoint will generate
JWT using Private Claims that contains name, phone, role, and timestamp
of the user that has the correct/matching phone and password`\n`{=tex}",
"requestBody": { "description": "User to add", "content": {
"application/json": { "schema": {
"$ref": "#/components/schemas/LoginUser"  }  }  }  },  "responses": {  "201": {  "description": "Created",  "content": {  "application/json": {  "schema": {  "$ref":
"\#/components/schemas/ResponseJWTToken" } } } }, "406": {
"description": "Not Acceptable", "content": { "application/json": {
"schema": {
"$ref": "#/components/schemas/NotAcceptableResponse"  }  }  }  }  }  }  },  "/v1/auth/check-token{token}": {  "get": {  "tags": [  "Auth-App"  ],  "description": "This endpoint will check the token param of the request, verify the JWT, and return the Private Claims of the JWT.\n",  "parameters": [  {  "name": "token",  "in": "path",  "description": "String jwt token",  "required": true,  "style": "simple",  "explode": false,  "schema": {  "type": "string",  "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRldGVoIiwicGhvbmUiOiI2Mjg1NzcxMjQ3MjgwIiwicm9sZSI6ImFkbWluIiwidGltZXN0YW1wIjoiMjAyMi0wMi0xOFQxNTo1MTo1MCJ9.uLdb5wBPyx_UL6XNTGoTDaEJlmucFhImM-7EjEhFwOw"  }  }  ],  "responses": {  "200": {  "description": "OK",  "content": {  "application/json": {  "schema": {  "$ref":
"\#/components/schemas/CheckTokenResponse" } } } }, "406": {
"description": "Not Acceptable", "content": { "application/json": {
"schema": {
"$ref": "#/components/schemas/NotAcceptableResponse"  }  }  }  }  }  }  },  "/v1/fetch/healthz": {  "get": {  "tags": [  "Fetch-App"  ],  "description": "This endpoint will return a \"It Works\" string and Status Code 200. This endpoint can be used to verify the app is running.\n",  "responses": {  "200": {  "description": "It Works",  "content": {  "application/json": {  "schema": {  "type": "string",  "description": "It Works"  }  }  }  }  }  }  },  "/v1/fetch/check-token{token}": {  "get": {  "tags": [  "Fetch-App"  ],  "description": "This endpoint will check the token param of the request, verify the JWT, and return the Private Claims of the JWT.\n",  "parameters": [  {  "name": "token",  "in": "path",  "description": "String jwt token",  "required": true,  "style": "simple",  "explode": false,  "schema": {  "type": "string",  "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRldGVoIiwicGhvbmUiOiI2Mjg1NzcxMjQ3MjgwIiwicm9sZSI6ImFkbWluIiwidGltZXN0YW1wIjoiMjAyMi0wMi0xOFQxNTo1MTo1MCJ9.uLdb5wBPyx_UL6XNTGoTDaEJlmucFhImM-7EjEhFwOw"  }  }  ],  "responses": {  "200": {  "description": "OK",  "content": {  "application/json": {  "schema": {  "$ref":
"\#/components/schemas/CheckTokenResponse" } } } }, "406": {
"description": "Not Acceptable", "content": { "application/json": {
"schema": {
"$ref": "#/components/schemas/NotAcceptableResponse"  }  }  }  }  }  }  },  "/v1/fetch/resources{token}": {  "get": {  "tags": [  "Fetch-App"  ],  "description": "This endpoint will fetch resources data from https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list, then returns the cleaned data with additional field of the price in USD currency\n",  "parameters": [  {  "name": "token",  "in": "path",  "description": "String jwt token",  "required": true,  "style": "simple",  "explode": false,  "schema": {  "type": "string",  "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRldGVoIiwicGhvbmUiOiI2Mjg1NzcxMjQ3MjgwIiwicm9sZSI6ImFkbWluIiwidGltZXN0YW1wIjoiMjAyMi0wMi0xOFQxNTo1MTo1MCJ9.uLdb5wBPyx_UL6XNTGoTDaEJlmucFhImM-7EjEhFwOw"  }  }  ],  "responses": {  "200": {  "description": "OK",  "content": {  "application/json": {  "schema": {  "$ref":
"\#/components/schemas/Resource" } } } }, "406": { "description": "Not
Acceptable", "content": { "application/json": { "schema": {
"$ref": "#/components/schemas/NotAcceptableResponse"  }  }  }  }  }  }  },  "/v1/fetch/resources/admin{token}": {  "get": {  "tags": [  "Fetch-App"  ],  "description": "This endpoint will fetch commodities data from https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list, clean the data from nil, null, and/or empty values, then returns the aggregated data by the area_provinsi value, weekly, and returns the max, min, avg, and median profit (assuming profit is price * size). This endpoint requires 'admin' role inside the valid JWT in the header of the request.\n",  "parameters": [  {  "name": "token",  "in": "path",  "description": "String jwt token",  "required": true,  "style": "simple",  "explode": false,  "schema": {  "type": "string",  "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRldGVoIiwicGhvbmUiOiI2Mjg1NzcxMjQ3MjgwIiwicm9sZSI6ImFkbWluIiwidGltZXN0YW1wIjoiMjAyMi0wMi0xOFQxNTo1MTo1MCJ9.uLdb5wBPyx_UL6XNTGoTDaEJlmucFhImM-7EjEhFwOw"  }  }  ],  "responses": {  "200": {  "description": "OK",  "content": {  "application/json": {  "schema": {  "$ref":
"\#/components/schemas/AgregateResource" } } } }, "406": {
"description": "Not Acceptable", "content": { "application/json": {
"schema": {
"$ref": "#/components/schemas/NotAcceptableResponse"  }  }  }  }  }  }  }  },  "components": {  "schemas": {  "User": {  "type": "object",  "properties": {  "data": {  "$ref":
"\#/components/schemas/User_data" }, "errors": { "type": "object" },
"meta": { "type": "object" } } }, "UpsertUser": { "required": \[
"phone", "role", "username" \], "type": "object", "properties": {
"phone": { "type": "string", "example": "6280000000000" }, "username": {
"type": "string", "example": "XXX" }, "role": { "type": "string",
"example": "Admin" } } }, "LoginUser": { "required": \[ "password",
"phone" \], "type": "object", "properties": { "phone": { "type":
"string", "example": "6280000000000" }, "password": { "type": "string",
"example": "z14X" } } }, "ResponseJWTToken": { "type": "object",
"properties": { "data": { "type": "object", "example":
"2022eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRkIiwicGhvbmUiOiI2Mjg1NzcxMjQ3MjcwIiwicm9sZSI6ImFkbW9wc2luIiwidGltZXN0YW1wIjoiMjAyMi0wMi0xOVQxNDoxMjo1MyJ9.0zSagz28AtswSX6lFhE8T078Hk6HUrR5Iyn8NeSaM38"
}, "errors": { "type": "object" }, "meta": { "type": "object" } } },
"CheckTokenResponse": { "type": "object", "properties": { "data": {
"$ref": "#/components/schemas/CheckTokenResponse_data"  },  "errors": {  "type": "object"  },  "meta": {  "type": "object"  }  }  },  "NotAcceptableResponse": {  "type": "object",  "properties": {  "data": {  "type": "object"  },  "errors": {  "type": "object",  "example": "Not Authorized"  },  "meta": {  "type": "object"  }  }  },  "Resource": {  "type": "object",  "properties": {  "data": {  "type": "array",  "items": {  "$ref":
"\#/components/schemas/Resource_data" } }, "errors": { "type": "object"
}, "meta": { "type": "object" } } }, "AgregateResource": { "type":
"object", "properties": { "data": { "type": "array", "items": { "\$ref":
"\#/components/schemas/AgregateResource_data" } }, "errors": { "type":
"object" }, "meta": { "type": "object" } } }, "User_data": { "type":
"object", "properties": { "id": { "type": "string", "example": "1" },
"username": { "type": "string", "example": "XXX" }, "phone": { "type":
"string", "example": "6280000000000" }, "role": { "type": "string",
"example": "Admin" }, "password": { "type": "string", "example": "53A2"
}, "timestamp": { "type": "string", "example":
"2022-02-19T14:12:53.000+0000" } } }, "CheckTokenResponse_data": {
"type": "object", "properties": { "id": { "type": "string", "example":
"0" }, "username": { "type": "string", "example": "XXX" }, "phone": {
"type": "string", "example": "6280000000000" }, "role": { "type":
"string", "example": "Admin" }, "password": { "type": "string",
"example": \"\" }, "timestamp": { "type": "string", "example":
"2022-02-19T14:12:53.000+0000" } } }, "Resource_data": { "type":
"object", "properties": { "uuid": { "type": "string", "example":
"6eaa323d-22ca-4a64-8c38-5ba209f21b5d" }, "komoditas": { "type":
"string", "example": "Ikan Tongkol" }, "area_provinsi": { "type":
"string", "example": "Jawa Timur" }, "area_kota": { "type": "string",
"example": "SITUBONDO" }, "size": { "type": "string", "example": "100"
}, "price": { "type": "string", "example": "60000" }, "usd_price": {
"type": "string", "example":
"$4.176266"  },  "tgl_parsed": {  "type": "string",  "example": "2022-02-14T13:57:10.439+0000"  },  "timestamp": {  "type": "string",  "example": "1644847030439"  }  }  },  "AgregateResource_Profit_Tahun X": {  "type": "object",  "properties": {  "Minggu Y": {  "type": "number",  "example": 72000000  }  }  },  "AgregateResource_Profit": {  "type": "object",  "properties": {  "Tahun X": {  "$ref":
"\#/components/schemas/AgregateResource_Profit_Tahun X" } } },
"AgregateResource_data": { "type": "object", "properties": {
"area_provinsi": { "type": "string", "example": "JAWA TENGAH" },
"Profit": { "\$ref": "\#/components/schemas/AgregateResource_Profit" },
"max_profit": { "type": "number", "example": 72000000 }, "min_profit": {
"type": "number", "example": 72000000 }, "average_profit": { "type":
"number", "example": 72000000 }, "median_profit": { "type": "number",
"example": 72000000 } } } } } }

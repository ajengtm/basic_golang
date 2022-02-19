# basic_golang 
Auth-App and Fetch-App

## Setup
* `BASEURL`to define the consistent part of your web address
* `SERVER_PORT` to define which port that the application will listen to 
	* The default port is Auth-App `8089`
    * The default port is Fetch-App `8090`
* `CORS` to define enables controlled access to resources located outside system.
* `CurrencyConverter` to define fetch CurrencyConverter URL & APIkey
* `Efishery` to define fetch resources URL

## Auth-App
`auth-app` will manage new user, password, and JWT generation process using a file based database.

### Auth-App Run
* Run the application directly in local machine 
`go run main.go http-auth`
or
`make start-http-auth`

### Auth-App Summary
Auth-App has 4 endpoints that can be accessed. 
* GET `/v1/auth/healthz` 
-- This endpoint will return a `"It Works"` string and `Status Code` `200`. This endpoint can be used to verify the app is running.
* POST `/v1/auth/upsert-user`
-- This endpoint will create user data that is received from the request body into the database and will return the 4 characters password for that user.
```
//HTTP Request Body (JSON)
{
    "phone": "6280000000000",
    "username": "XXX",
    "role":"admin"
}

//HTTP Response (Application/JSON)
{
    "data": {
        "id": 3,
        "username": "XXX",
        "phone": "6280000000000",
        "role": "admin",
        "password": "<4 characters string>",
        "timestamp": "2022-02-18T19:29:24"
    },
    "errors": {},
    "meta": {}
}
```
* POST `/v1/auth/login`
-- This endpoint will generate JWT using `Private Claims` that contains `name`, `phone`, `role`, and `timestamp` of the user that has the correct/matching `phone` and `password`.
```
//HTTP Request Body (JSON)
{
	"phone": "6280000000000",
	"password": "<4 characters string>"
}

//HTTP Response (Application/JSON)
{
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImJiIiwicGhvbmUiOiI2Mjg1NzcxMjQ3Mjg5Iiwicm9sZSI6ImFkbWluIiwidGltZXN0YW1wIjoiMjAyMi0wMi0xOFQxOToyOToyNCJ9.DCXrwL-GfnObgDxxiZluGBX8mRfJBOLOYiQM8SY57yU",
    "errors": {},
    "meta": {}
}
```
* GET `/v1/auth/check-token`
-- This endpoint will check the `token` param of the request, verify the JWT, and return the `Private Claims` of the JWT.
```
//HTTP Request Param
token: <JWT>

//HTTP Response (Application/JSON)
{
    "data": {
        "id": 0,
        "username": "XXX",
        "phone": "6280000000000",
        "role": "admin",
        "password": "",
        "timestamp": "2022-02-16T22:51:06"
    },
    "errors": {},
    "meta": {}
}
```

## Fetch-App
`fetch-app` will fetch and process resources

### Fetch-App Run
* Run the application directly in local machine 
`go run main.go http-fetch`
or
`make start-http-auth`

### Fetch-App Summary
Fetch-App has 4 endpoints that can be accessed. 
* GET `/v1/fetch/healthz` 
-- This endpoint will return a `"It Works!"` string and `Status Code` `200`. This endpoint can be used to verify the app is running.

* GET `/v1/fetch/check-token`
-- This endpoint will check the `token` param of the request, verify and return the `Private Claims` of the JWT.
```
//HTTP Request Param
check-token:  <JWT>

//HTTP Response (Application/JSON)
{
    "data": {
        "id": 0,
        "username": "teteh",
        "phone": "6285771247280",
        "role": "admin",
        "password": "",
        "timestamp": "2022-02-18T15:51:50"
    },
    "errors": {},
    "meta": {}
}
```
* GET `/v1/fetch/resources`
-- This endpoint will fetch resources data from https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list, then returns the cleaned data with additional field of the price in USD currency
```
//HTTP Request Param (Bearer Token)
token: <JWT>

//HTTP Response (Application/JSON)
{
  "data": [
    {
      "uuid": "6eaa323d-22ca-4a64-8c38-5ba209f21b5d",
      "komoditas": "Ikan Tongkol",
      "area_provinsi": "Ikan Tongkol",
      "area_kota": "SITUBONDO",
      "size": "100",
      "price": "60000",
      "usd_price": "$4.178695",
      "tgl_parsed": "2022-02-14T13:57:10.439Z",
      "timestamp": "1644847030439"
    },
    {
      "uuid": "3f74bc97-a4f4-424c-b053-d9d689323546",
      "komoditas": "Test",
      "area_provinsi": "Test",
      "area_kota": "PADANG PARIAMAN",
      "size": "40",
      "price": "1",
      "usd_price": "$0.000070",
      "tgl_parsed": "2022-02-18T05:07:50.367Z",
      "timestamp": "1645160870367"
    }
  ],
  "errors": {},
  "meta": {}
}
```
* GET `/v1/fetch/resources/admin`
-- This endpoint will fetch commodities data from https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list, clean the data from `nil`, `null`, and/or empty values, then returns the aggregated data by the `area_provinsi` value, weekly, and returns the max, min, avg, and median profit (assuming profit is `price` * `size`). This endpoint requires 'admin' role inside the valid JWT in the header of the request.
```
//HTTP Request Param
token: <JWT>

//HTTP Response (Application/JSON)
{
  "data": [
    {
      "area_provinsi": "JAWA TENGAH",
      "Profit": {
        "Tahun 2022": {
          "Minggu ke 7": 72000000
        }
      },
      "max_profit": 72000000,
      "min_profit": 72000000,
      "average_profit": 72000000,
      "median_profit": 72000000
    },
    {
      "area_provinsi": "BULELENG",
      "Profit": {
        "Tahun 2022": {
          "Minggu ke 7": 4800000
        }
      },
      "max_profit": 4800000,
      "min_profit": 4800000,
      "average_profit": 4800000,
      "median_profit": 4800000
    }
  ],
  "errors": {
    
  },
  "meta": {
    
  }
}
```

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
There are 2 ways to run the app :
* Run the application directly in local machine 
`go run main.go http-auth`
or
`make start-http-auth`

## Fetch-App
`fetch-app` will fetch and process resources

### Fetch-App Run
There are 2 ways to run the app :
* Run the application directly in local machine 
`go run main.go http-fetch`
or
`make start-http-auth`

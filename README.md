# testFramework
A testing framework to test REST APIs provided [here](http://www.groupkt.com/post/c9b0ccb9/country-and-other-related-rest-webservices.htm).

## Requirements
This is written in Go so if you wish to build and run locally, you need to have Go installed. Alternatively, if you don't have Go, you can use [this docker image](https://hub.docker.com/r/zusyed/test-framework/).

## Build Instructions
To build the code, run:
```console
$ go build 
```

To run all tests, run:
```console
$ go test -v 
```

To run specific tests, run:
```console
$ go test -v --run TestGetCountriesBySearch 
```

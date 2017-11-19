# testFramework
A testing framework to test REST APIs provided [here](http://www.groupkt.com/post/c9b0ccb9/country-and-other-related-rest-webservices.htm).

## Requirements
This is written in Go so if you wish to build and run locally, you need to have Go installed. Alternatively if you don't have Go, you can use [this docker image](https://hub.docker.com/r/zusyed/test-framework/).

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

## Design Decisions
The way this is implemented is similar to SDKs which call APIs with the function arguments being the parameters in the URL. To test the actual APIs, unit tests are written with various function arguments.

## Assumptions
Since all the APIs are simple HTTP GET requests, the functions only accept the parameters needed to pass in the URL. There maybe tests needed to run with different HTTP headers but it is assumed that testing with different headers is not required. Also only the HTTP status code and the response body is validated. Other response headers are not validated, although it's not hard to do so with this implementation.

## Limitations
As highlighted earlier, the limitations include passing HTTP headers to the request and validating HTTP headers reurned in the response. Both of which are not hard to add but were excluded for brevity.

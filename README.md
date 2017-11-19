# testFramework
A testing framework to test REST APIs provided [here](http://www.groupkt.com/post/c9b0ccb9/country-and-other-related-rest-webservices.htm).

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

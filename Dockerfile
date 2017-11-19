FROM golang:1.9

ADD . /go/src/github.com/zusyed/testFramework

WORKDIR /go/src/github.com/zusyed/testFramework
RUN go build

ENTRYPOINT ["./gotest.sh"]

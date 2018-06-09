FROM golang:1.10

LABEL Description="Image of our Hero of Parking Spaces API" Vendor="Mike Hall" Version="1.0"

WORKDIR /go/src

RUN go get github.com/golang/dep/cmd/dep
RUN go get github.com/codegangsta/gin

CMD pwd ; ls -al ; gin --port 8888 run api.go

EXPOSE 8888

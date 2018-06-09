FROM golang:stretch

LABEL Description="Image of our Hero of Parking Spaces API" Vendor="Mike Hall" Version="1.0"

WORKDIR /go/src/api

RUN go get github.com/golang/dep/cmd/dep
RUN go get github.com/codegangsta/gin

COPY Gopkg.toml Gopkg.toml
COPY Gopkg.lock Gopkg.lock
RUN dep ensure --vendor-only

COPY entrypoint.sh /root/
RUN chmod 755 /root/entrypoint.sh
RUN chmod +x /root/entrypoint.sh

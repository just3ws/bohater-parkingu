FROM golang:alpine

RUN go get github.com/tools/godep

COPY . /go/src/github.com/just3ws/plamka-bohater

ENV USER mike
ENV HTTP_ADDR :8888
ENV HTTP_DRAIN_INTERVAL 1s
ENV COOKIE_SECRET hLEXCuXkgeERnyYJ

WORKDIR /go/src/github.com/just3ws/plamka-bohater

RUN godep go build

EXPOSE 8888
CMD ./plamka-bohater


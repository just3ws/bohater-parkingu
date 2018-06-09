FROM golang:1.10

RUN go get github.com/golang/dep
RUN go get github.com/pilu/fresh

COPY . /go/src/github.com/just3ws/bohater-parkingu/api
WORKDIR /go/src/github.com/just3ws/bohater-parkingu/api

RUN go get ./
RUN go build

RUN pwd

CMD fresh

EXPOSE 8888

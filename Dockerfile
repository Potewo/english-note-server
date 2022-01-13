FROM golang:1.17-alpine3.15

RUN apk update && apk add git

RUN mkdir /go/src/app
WORKDIR /go/src/app

COPY . /go/src/app

RUN go build

CMD ["/go/src/app/english-note-server"]

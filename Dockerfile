FROM node:lts-alpine as build-stage
WORKDIR /app
COPY ./english-note/package*.json ./
RUN npm install
COPY ./english-note .
RUN npm run build

FROM golang:1.17-alpine3.15

RUN apk update && apk add git

RUN mkdir /go/src/app
WORKDIR /go/src/app
COPY --from=build-stage /app/dist /go/src/app/public

COPY . /go/src/app

RUN go build

CMD ["/go/src/app/english-note-server"]

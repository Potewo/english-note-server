FROM node:lts-alpine as build-stage
WORKDIR /app
COPY ./english-note/package*.json ./
RUN yarn install
COPY ./english-note .
RUN yarn build

FROM golang:1.17-alpine3.15

RUN mkdir /go/src/app
WORKDIR /go/src/app
COPY --from=build-stage /app/dist /go/src/app/public

COPY . /go/src/app

RUN go build


CMD ["/go/src/app/english-note-server"]

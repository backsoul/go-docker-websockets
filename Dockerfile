# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY /public ./public
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /websocket

EXPOSE 8080

CMD [ "/websocket" ]
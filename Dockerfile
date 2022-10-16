FROM golang:latest

WORKDIR /mashi

COPY ./go.mod .
RUN go mod download

COPY . .


RUN go build

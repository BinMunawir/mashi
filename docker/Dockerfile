FROM golang:latest as base

WORKDIR /mashi

COPY go.mod .
RUN go mod download


FROM base as dev

RUN apt update -y && apt upgrade -y && apt install nano -y
RUN alias ll="ls -al --color"

RUN go install github.com/cosmtrek/air@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest


FROM base

EXPOSE 8080


FROM golang:latest as base

WORKDIR /mashi

COPY go.mod .
RUN go mod download


FROM base as dev

RUN apt update -y && apt upgrade -y && apt install nano -y
RUN go install github.com/cosmtrek/air@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go install -v golang.org/x/tools/gopls@latest 
RUN go install -v github.com/ramya-rao-a/go-outline@v0.0.0-20210608161538-9736a4bde949
RUN echo 'alias ll="ls -al --color"' >> ~/.bashrc





FROM base

EXPOSE 8080


FROM golang:latest

WORKDIR /mashi

COPY ./go.mod .
RUN go mod download

RUN go install github.com/cosmtrek/air@latest
COPY . .

# RUN go build

FROM golang:latest

WORKDIR /mashi

COPY ./go.mod .
RUN go mod download

COPY . .
EXPOSE 80
CMD go run .

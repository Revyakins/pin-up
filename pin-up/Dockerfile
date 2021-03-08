FROM golang:alpine

RUN apk update && apk add --no-cache git
WORKDIR /opt/pin-up
COPY . .
RUN go mod download
RUN go get -u github.com/cosmtrek/air
ENTRYPOINT air
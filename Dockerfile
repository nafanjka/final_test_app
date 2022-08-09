FROM golang:latest
ADD . /app
WORKDIR /app
RUN go build app/server.go
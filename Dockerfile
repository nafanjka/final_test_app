FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build app/server.go
EXPOSE 80
ENTRYPOINT /app/server
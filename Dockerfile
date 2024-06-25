FROM golang:1.20.3-buster

WORKDIR /code/

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./

RUN apt-get update && \
  go mod download

CMD ["air"]

FROM golang:1.23.4-bullseye

WORKDIR /code/

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN apt-get update && \
  go mod download

CMD ["air"]

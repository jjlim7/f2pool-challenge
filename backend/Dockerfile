FROM golang:latest as builder

RUN apt-get update
RUN apt install -y curl

RUN GO111MODULE=on

ENV GO_PATH=$HOME/go
ENV PATH="$PATH:$GO_PATH/bin"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy
CMD ["go", "run", "cmd/main.go"]

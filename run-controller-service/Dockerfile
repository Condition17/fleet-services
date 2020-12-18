
# Base build image

FROM golang:alpine as builder_base

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

# Use modules in go compiler
ENV GO111MODULE=on

# Steps that will be cached
COPY go.mod .
COPY go.sum .

RUN go mod download

# Image that only builds the service
FROM builder_base as service_builder
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service *.go

# We start from a fresh Alpine image, to reduce the image size and not ship the Go compiler in our production artifacts.
FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY --from=service_builder /app .

ENTRYPOINT ["./service"]
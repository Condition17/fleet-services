FROM golang:1.14-buster as builder

RUN apt-get update && apt-get upgrade

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

# Steps that will be cached
COPY go.mod .
COPY go.sum .

RUN go get -u
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o service -v .

# Run container
FROM debian:buster

# ---- Install docker ----
RUN apt-get update && apt-get upgrade -y
RUN apt-get -y install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common && \
    curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add - && \
    apt-key fingerprint 0EBFCD88 && \
    add-apt-repository \
       "deb [arch=amd64] https://download.docker.com/linux/debian \
       $(lsb_release -cs) \
       stable" && \
    apt-get update && \
    apt-get install -y docker-ce

# ---- Install prerequisites for NFS mounting ----
RUN apt-get -y update && apt-get -y install nfs-common

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app .

ENTRYPOINT ["./service"]
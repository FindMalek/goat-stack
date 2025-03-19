FROM golang:1.24-alpine AS builder

WORKDIR /go/src/app

RUN apk add --no-cache tzdata
ENV TZ=UTC

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-s -w" -o goat ./cmd/goat



FROM alpine

RUN apk add bash

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /go/src/app/goat /opt/bin/goat
COPY --from=litestream/litestream:latest /usr/local/bin/litestream /usr/local/bin/litestream


RUN mkdir -p /data
# WORKDIR /opt/bin

EXPOSE 8080
COPY scripts/run.sh /scripts/run.sh
COPY etc/litestream.yml /etc/litestream.yml


ENV TZ=UTC
ENV USER=1000

ENTRYPOINT "/scripts/run.sh"

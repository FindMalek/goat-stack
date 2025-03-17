FROM golang:1.24-alpine as builder

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

FROM scratch

WORKDIR /opt/bin

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /go/src/app/goat /opt/bin/goat

ENV TZ=UTC
ENV USER=1000

CMD [ "/opt/bin/goat" ]

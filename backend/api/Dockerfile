FROM golang:1.20-alpine as builder

ENV GO111MODULE=on

RUN apk update \
  && apk upgrade \
  && apk add --no-cache \
  build-base libtool musl-dev ca-certificates dumb-init curl

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app
RUN go mod tidy

RUN go build -o main main.go

FROM alpine:latest

RUN apk update \
  && apk upgrade \
  && apk add --no-cache \
  make ca-certificates curl \
  && update-ca-certificates 2>/dev/null || true

WORKDIR /app
COPY --from=builder /app/main /app/

CMD /app/main
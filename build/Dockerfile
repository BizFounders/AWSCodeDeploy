# syntax=docker/dockerfile:experimental
FROM golang:1.23 AS builder

COPY . /go/src/app/

WORKDIR /go/src/app

RUN go mod tidy && CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags="-w -s" -o /app /go/src/app/

FROM alpine:3.18

RUN apk --no-cache add dumb-init

COPY --from=builder /app /app/app
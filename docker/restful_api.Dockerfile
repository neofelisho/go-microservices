ARG GOLANG_VERSION=1.14
ARG ALPINE_VERSION=3.11

FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} as builder

COPY go.mod /app/
COPY go.sum /app/
WORKDIR /app
RUN go mod download

COPY . /app
WORKDIR /app
RUN go build -ldflags "-s -w" -o /release/restful_api ./cmd/restful_api

FROM alpine:${ALPINE_VERSION}
WORKDIR /app
COPY --from=builder /release .

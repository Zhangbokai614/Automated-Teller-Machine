ARG GO_VERSION=1.16

FROM golang:${GO_VERSION}-alpine AS builder

ENV GOPROXY="https://goproxy.io"
ENV CGO_ENABLED=0
WORKDIR /go/src/app
COPY . .
RUN go build -o ./app ./cmd/main.go

FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app
COPY --from=builder /go/src/app/app .

EXPOSE 8081

ENTRYPOINT ["./app"]
FROM golang:1.20-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o apiserver ./cmd

FROM alpine:latest
RUN apk add --update --no-cache ca-certificates git

COPY --from=builder ["/build/apiserver", "/"]

EXPOSE 8080

ENTRYPOINT ["/apiserver"]

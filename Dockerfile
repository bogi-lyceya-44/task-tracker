FROM golang:1.24.6-alpine3.22 AS builder
WORKDIR /build
ADD go.mod .
COPY . .
RUN go build cmd/main.go
FROM alpine:3.22.1
WORKDIR /build
COPY --from=builder /build/main /build/main
CMD ["./main"]

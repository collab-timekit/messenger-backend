FROM golang:1.24.2 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/app .
COPY --from=builder /app/config /root/config

EXPOSE 8083
CMD ["./app"]
FROM golang:1.20 AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./bin/app ./cmd/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/bin ./bin

CMD ["./bin/app"]
EXPOSE 8000/tcp
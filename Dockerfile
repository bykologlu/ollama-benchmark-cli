FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ollama-benchmark main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/ollama-benchmark .

COPY prompts.txt ./prompts.txt
COPY internal/i18n/lang.json ./internal/i18n/lang.json

CMD ["./ollama-benchmark"]


FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o webapp .

FROM alpine:3.18

RUN mkdir -p /var/log/webapp

WORKDIR /app

COPY --from=builder /app/webapp .
COPY public/ ./public/

EXPOSE 8080

CMD ["./webapp"]
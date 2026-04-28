FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
COPY . .

RUN go get github.com/labstack/echo/v4/middleware@v4.11.1 && \
    go mod tidy && \
    go build -o webapp .

FROM alpine:3.18

RUN mkdir -p /var/log/webapp

WORKDIR /app

COPY --from=builder /app/webapp .
COPY public/ ./public/

EXPOSE 8080

CMD ["sh", "-c", "(sleep 2 && for i in 1 2 3 4 5 6 7 8 9 10; do wget -q -O /dev/null http://localhost:${PORT:-8080}/ 2>/dev/null; sleep 1; done) & exec ./webapp"]
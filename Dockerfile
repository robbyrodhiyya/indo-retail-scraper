FROM golang:1.26.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
    go build -o indo-retail-scraper ./cmd

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/indo-retail-scraper .

CMD ["./indo-retail-scraper"]
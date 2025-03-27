FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY ../. /app/
RUN go build -o ikuai-bypass


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/ikuai-bypass .
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai

# CMD ["./ikuai-bypass", "-c", "/app/config.yml", "-r", "cron", "-m", "ip"]

FROM golang:bullseye as builder

RUN apt-get update && apt-get install -y \
    gcc \
    git \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY . .

RUN go mod tidy

RUN CGO_ENABLED=1 GOOS=linux go build -ldflags '-linkmode "external" -extldflags "-static"' -o app .

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

WORKDIR /root/
COPY --from=0 /app/config/config.yml.example ./config/config.yml
COPY --from=0 /app/app .

CMD ["./app"]

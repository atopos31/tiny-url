FROM golang:1.22.4 AS builder

WORKDIR /app
# Proxy If in China
ARG proxy=https://proxy.golang.org
ENV proxy=$proxy
RUN echo "proxy: $proxy"

ENV GOPROXY=${proxy}

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 开启静态编译
RUN CGO_ENABLED=0 go build -o tiny-url .

FROM alpine:latest

WORKDIR /code

COPY --from=builder /app/tiny-url ./tiny-url
# Web 资源
COPY --from=builder /app/web ./web

RUN chmod +x ./tiny-url

CMD ["./tiny-url"]
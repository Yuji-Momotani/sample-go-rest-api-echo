FROM golang:1.21-alpine

WORKDIR /app
COPY . /app
# COPY go.mod /app
# COPY startup.sh /app
# COPY air.toml /app

RUN go mod download
RUN apk add --no-cache bash

# delvのインストール（デバッグ）
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# airのインストール（ホットリロード）
RUN go install github.com/cosmtrek/air@v1.49.0

ENV PATH="/go/bin:${PATH}"
RUN chmod +x ./startup.sh
EXPOSE 8080

# なぜかビルドするとうまく起動できない。（air経由で起動するしか無理）
# 理由は今のところわかりません...
# RUN go build -o main main.go
# RUN chmod +x /app/main

CMD ["/bin/bash", "./startup.sh"]
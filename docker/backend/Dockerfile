FROM golang:1.17.6-alpine3.15 AS builder

ENV GOPROXY=https://goproxy.io,direct

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories

RUN apk add --no-cache make bash git gcc g++ && mkdir -p /gin-admin

COPY . /gin-admin

WORKDIR /gin-admin

RUN make build

FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories

RUN apk --no-cache add ca-certificates tzdata && mkdir -p /gin-admin

COPY --from=builder /gin-admin/cmd/gin-admin/gin-admin /gin-admin

COPY ./configs /gin-admin/configs

WORKDIR /gin-admin
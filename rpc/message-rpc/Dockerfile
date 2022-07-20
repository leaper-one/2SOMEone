FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 1
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY rpc/message-rpc/etc /app/etc
RUN go build -o /app/message rpc/message-rpc/message.go


FROM ubuntu:20.04 as prof

#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
#COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/message /app/message
COPY --from=builder /app/etc /app/etc

CMD ["./message", "-f", "etc/message.yaml"]

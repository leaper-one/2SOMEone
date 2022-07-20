FROM golang:1.17.10 AS builder

ENV CGO_ENABLED 1
ENV GO111MODULE=on
ENV GOARCH=amd64
ENV GOOS=linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build

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

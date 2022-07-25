FROM golang:1.17.10 AS builder

ENV CGO_ENABLED 1
ENV GO111MODULE=on
ENV GOARCH=amd64
ENV GOOS=linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build

COPY . .
COPY rpc/user-rpc/etc /app/etc
RUN go build -o /app/user rpc/user-rpc/user.go


FROM ubuntu:20.04 as prof

#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
#COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
RUN apt-get -qq update \
    && apt-get -qq install -y --no-install-recommends ca-certificates curl
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/user /app/user
COPY --from=builder /app/etc /app/etc

CMD ["./user", "-f", "etc/user.yaml"]

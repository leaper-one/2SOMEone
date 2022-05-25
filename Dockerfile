FROM golang:1.17.8 as builder

MAINTAINER cunoe

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/2SOMEone

COPY . .

RUN  go build -o ./user/user ./user/main.go ./user/user.go ./user/load_config.go

FROM ubuntu:20.04 as prod

WORKDIR /root/

COPY --from=builder /go/src/2SOMEone/user/user .

CMD ["./user"]
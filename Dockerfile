FROM golang:1.17.8-alpine as builder

RUN apk --no-cache add git
MAINTAINER cunoe

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/2SOMEone

COPY . .

RUN go build -o ./user/linux_$GOARCH/user ./user/main.go ./user/user.go ./user/load_config.go

FROM alpine:latest as prod

ARG GOARCH=amd64

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/2SOMEone/user/linux_$GOARCH/user .

CMD ["./user"]
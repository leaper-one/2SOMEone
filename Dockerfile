FROM golang:1.17.8 as builder

MAINTAINER cunoe

WORKDIR /go/src/2SOMEone

COPY . .

RUN  make build-docker

FROM ubuntu:20.04 as prod

WORKDIR /root/

COPY --from=builder /go/src/2SOMEone/user/docker/user .

CMD ["./user"]

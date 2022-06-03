FROM golang:1.17.8 as builder

MAINTAINER cunoe

WORKDIR /go/src/2SOMEone

COPY . .

RUN  make linux-user

FROM ubuntu:20.04 as prod

WORKDIR /root/

COPY --from=builder /go/src/2SOMEone/user/linux_amd64/user .

CMD ["./user"]

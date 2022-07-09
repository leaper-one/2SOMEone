#
# BUILD
#

FROM golang:1.17.8 as builder
MAINTAINER cunoe
WORKDIR /go/src/2SOMEone
COPY . .
RUN  make linux-user
RUN  make linux-message


#
# IMAGE TARGET
#

FROM ubuntu:20.04 as user
WORKDIR /root/
COPY --from=builder /go/src/2SOMEone/user/linux_amd64/user .
CMD ["./user"]

FROM ubuntu:20.04 as message
WORKDIR /root/
COPY --from=builder /go/src/2SOMEone/message/linux_amd64/message .
CMD ["./message"]

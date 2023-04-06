#FROM ubuntu:14.04
#
#RUN apt-get update \
#    && apt-get install -y software-properties-common \
#    && add-apt-repository ppa:gophers/archive \
#    && apt-get update \
#    && apt-get install -y golang-1.19-go \
#                          git \
#    && rm -rf /var/lib/apt/lists/*
#
#ENV GOPATH /root/go
#ENV GOROOT /usr/lib/go-1.19
#ENV PATH="/usr/lib/go-1.19/bin:$PATH"
#
#COPY ./com/http/main.go /root/main.go
#RUN go build -o /root/httpd /root/main.go \
#    && chmod +x /root/httpd
#WORKDIR /root
#ENTRYPOINT ["/root/httpd"]

FROM golang:1.19-alpine as builder

WORKDIR /go/src
COPY ./com/http/main.go .
RUN go build -o httpd ./main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/httpd .
RUN chmod +x /root/httpd

ENTRYPOINT ["/root/httpd"]
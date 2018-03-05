# Amazon Linux 2 image
FROM amazonlinux:2017.12.0.20180222

MAINTAINER RyoTakashima <r.takashima@miniascape.co.jp>

# Go 1.10 image
# FROM golang:1.10.0-stretch

# install wget tar gzip
RUN yum install -y wget tar gzip

# install go1.10
RUN wget -P /usr/local/ https://storage.googleapis.com/golang/go1.10.linux-amd64.tar.gz
RUN tar zxvf /usr/local/go1.10.linux-amd64.tar.gz -C /usr/local/
RUN rm /usr/local/go1.10.linux-amd64.tar.gz

# GOPATH
RUN export PATH="/usr/local/go/bin:$PATH"
RUN export GOPATH=/go
RUN export PATH=$PATH:$GOPATH/bin

# RUN exportしただけでは何故か環境変数が設定されないので、 ENVコマンドで定義する
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# install mattes/migrate (Database migrations. CLI and Golang library.)
RUN wget -P /usr/local/bin https://github.com/mattes/migrate/releases/download/v3.0.1/migrate.linux-amd64.tar.gz
RUN tar xvzf /usr/local/bin/migrate.linux-amd64.tar.gz -C /usr/local/bin/
RUN mv /usr/local/bin/migrate.linux-amd64 /usr/local/bin/migrate
RUN rm /usr/local/bin/migrate.linux-amd64.tar.gz

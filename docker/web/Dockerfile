# Amazon Linux 2 image
FROM amazonlinux:2017.12.0.20180222

MAINTAINER RyoTakashima <r.takashima@miniascape.co.jp>

# ansible
RUN yum install -y python-setuptools
RUN easy_install pip
RUN pip install ansible
RUN mkdir /etc/ansible/
ADD ansible/* /etc/ansible/

# GOPATH
ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin:/usr/local/go/bin
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# Ansible
WORKDIR /etc/ansible/
RUN ansible-playbook setup.yml

# ワーキングディレクトリ(docker-compose run やexecはこのディレクトリ上で実行される)
WORKDIR /go/src/github.com/wmetaw/go-ddd-on-echo/

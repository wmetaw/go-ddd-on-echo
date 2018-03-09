# Amazon Linux 2 image
FROM wmetaw/amazon-linux2-ec2

MAINTAINER Ryo Takashima

# ansible
RUN easy_install pip
RUN pip install ansible
RUN mkdir /etc/ansible/
ADD ansible/* /etc/ansible/

# GOPATH
ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin:/usr/local/go/bin

# Provisioning by Ansible
WORKDIR /etc/ansible/
RUN ansible-playbook setup.yml

# ワーキングディレクトリ
# (docker-compose run やexecはこのディレクトリ上で実行される)
WORKDIR /go/src/github.com/wmetaw/go-ddd-on-echo/

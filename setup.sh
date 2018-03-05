# install wget tar gzip
yum install -y wget tar gzip

# install go1.10
wget -P /usr/local/ https://storage.googleapis.com/golang/go1.10.linux-amd64.tar.gz
tar zxvf /usr/local/go1.10.linux-amd64.tar.gz -C /usr/local/
rm /usr/local/go1.10.linux-amd64.tar.gz

# GOPATH
export PATH="/usr/local/go/bin:$PATH"
export GOPATH=/go
export PATH=$PATH:$GOPATH/bin

# exportしただけでは何故か環境変数が設定されないので、 ENVコマンドで定義する
mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# install mattes/migrate (Database migrations. CLI and Golang library.)
wget -P /usr/local/bin https://github.com/mattes/migrate/releases/download/v3.0.1/migrate.linux-amd64.tar.gz
tar xvzf /usr/local/bin/migrate.linux-amd64.tar.gz -C /usr/local/bin/
mv /usr/local/bin/migrate.linux-amd64 /usr/local/bin/migrate
rm /usr/local/bin/migrate.linux-amd64.tar.gz

# How to Install

## GOPATHの設定
環境変数にGOPATHを設定する
※既に設定している場合は既存のGOPATHでOK

```bash:Mac
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

```bash:windows
$ echo $GOPATH
C:¥Users¥miniascape100¥go
```

### プロジェクトのclone

```bash:Mac
$ cd $GOPATH/src/github.com/wmetaw/
$ git clone git@github.com:wmetaw/go-ddd-on-echo.git
```

```bash:WIndows
$ cd $GOPATH¥src¥github.com¥wmetaw¥
$ git clone git@github.com:wmetaw/go-ddd-on-echo.git
```

# Docker for [Mac](https://docs.docker.com/docker-for-mac/install/) / [Windows](https://docs.docker.com/docker-for-windows/install/)

```bash:　
# docker-compose.ymlがあるフォルダまで移動
$ cd $GOPATH¥src¥github.com¥wmetaw¥go-ddd-on-echo/docker

# コンテナを立ち上げる
$ docker-compose up

# migrationする
#  (別タブで cd $GOPATH¥src¥github.com¥wmetaw¥go-ddd-on-echo/docker )
$ docker-compose exec web migrate -database 'mysql://root:root@tcp(mysql-server:3306)/dev?charset=utf8&parseTime=True&loc=Local' -path /go/src/github.com/wmetaw/go-ddd-on-echo/config/migrate/ up
```

※ WindowsはHYPER-Vをオンにしたり、BIOS開いて `Virtualization Technology` を`Enable`にする必要があるっぽい。

# Windows10 Home
[Docker Tool Box](https://docs.docker.com/toolbox/toolbox_install_windows/)をダウンロードしてインストール
(※ Git for Windows、VirtualBoxのインストールにチェックを入れる)

※ WindowsはHYPER-Vをオンにしたり、BIOS開いて `Virtualization Technology` を`Enable`にする必要があるっぽい。

### docke コンテナの立ち上げ

```bash
$ cd C:¥Users¥miniascape100¥go¥src¥github.com¥wmetaw¥go-ddd-on-echo¥docker
$ docker-compose up
```

### ipを確認してアクセス

```bash
$ docker-machine ip
192.168.99.100

http://192.168.99.100:1323
```



# Tips
```bash:よく使うコマンド
# リビルド
docker-compose build --no-cache

# 稼働中コンテナ確認
docker-compose ps

# stop
docker-compose stop

# 実行中のコンテナに入る
docker-compose exec web bash

# コンテナを起動してに入る
docker-compose exec web bash

# mysql-serverへの接続
docker-compose ps | grep mysql
docker inspect docker_mysql-server_1 | grep IPAddress
docker-compose exec mysql-server mysql -h 172.18.0.3 -u root -proot

# redis-serverへの接続
docker-compose ps | grep redis
docker inspect docker_redis-server_1 | grep IPAddress
docker-compose exec redis-server redis-cli -h 172.18.0.2

# memcached-serverへの接続(ローカルマシンにtelnetコマンドが必要)
docker-compose ps | grep memcached
docker inspect docker_memcached-server_1 | grep IPAddress
telnet localhost 11211
# telnet 0.0.0.0 11211
# telnet 192.168.99.100 11211

```

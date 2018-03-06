# Usage

```
# docker-compose.ymlがあるファイルまで移動
cd path/to/docker

# コンテナを立ち上げる
docker-compose up

# migrate (別タブで cd path/to/docker )
docker-compose exec web \
migrate -database 'mysql://root:root@tcp(mysql-server:3306)/dev?charset=utf8&parseTime=True&loc=Local' \
-path /go/src/github.com/wmetaw/go-ddd-on-echo/config/migrate/ up
```


```
# リビルド
docker-compose build --no-cache

# 稼働中コンテナ確認
docker-compose ps

# stop
docker-compose stop

# 実行中コンテナに入る
docker-compose exec web bash
```
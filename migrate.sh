migrate -database 'mysql://root:root@tcp(mysql-server:3306)/dev?charset=utf8&parseTime=True&loc=Local' -path $GOPATH/src/github.com/wmetaw/go-ddd-on-echo/config/migrate/ up


docker-compose exec web sh /go/src/github.com/wmetaw/go-ddd-on-echo/migrate.sh

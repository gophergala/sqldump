# sqldump

A small tool for administration of databases. My first 48 hours in golang.

Unfortunately I do not have that much time, so it's just the backend with some basic html. 
Use your fantasy and imagination for fancy output with latest js-technology. 

## prepare

    sudo mysqladmin --defaults-file=/etc/mysql/debian.cnf create gotestdb
    sudo mysql --defaults-file=/etc/mysql/debian.cnf -e "GRANT ALL PRIVILEGES  ON gotestdb.*  TO 'go_user'@'localhost' IDENTIFIED BY 'mypassword'  WITH GRANT OPTION;"
    mysql -p"mypassword" -u go_user gotestdb -e 'create table posts (title varchar(64) default null, start date default null);'
    mysql -p"mypassword" -u go_user gotestdb -e 'insert into posts values("hello","2015-01-01");'
    mysql -p"mypassword" -u go_user gotestdb -e 'insert into posts values("more","2015-01-03");'
    mysql -p"mypassword" -u go_user gotestdb -e 'insert into posts values("end","2015-01-23");'
    mysql -p"mypassword" -u go_user gotestdb -B -e 'select * from posts;'

## install

    go get github.com/go-sql-driver/mysql

## run

    export GOPATH=~/bin/sqldump/
    go run sqldump.go dump.go aux.go

## usage

open http://localhost:8080/

## caveats

- You can't access a database named 'favicon.ico'
- Check for sql injection


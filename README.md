# Goland API

> This repo creates golang API.   

## Drafting big picture
Client > (https) > NGINX > API > DB (PostgreS)  
NGINX does proxy, slow clients serving, TLS termination, caching.
See video 01_api_creating.mp4 at 41:37 for big picture details.

## Drafting database
See video 01_api_creating.mp4 at 57:03 for the draft scheme.

 

## Usage 
### Creating api:

## Installation  
### PostgreS install
See [PostgreSQL packages for Debian](https://wiki.postgresql.org/wiki/Apt)
```bash
apt install postgresql-9.5
```
### PostgreS config file
```bash
cd /etc/postgresql/9.5/main/
postgres --config-file=/postgresql.conf
```
### PostgreS console
```bash
su - postgres
psql # postgres console
\l # list of databases  
\c api # connect to api db
\d # display all the tables in api db
create table users(id int, email varchar(128), password varchar(32));
alter table users add column role int;
alter table users add column status bool;
alter table users add column name varchar(128);
select * from users;  
create table roles(id int, name varchar(32));
select * from roles;
insert into roles values(1, 'Buyer');
insert into roles values(2, 'Admin');
select * from roles;
select * from users;
insert into users values(1,'makuznet@yandex.ru','netlab',2,true,'Max');
select * from users;
select * from roles where id=2;
select users.name, users.email, roles.name as role from users left join roles on users.role=roles.id
# there's a shorter record: 
select u.name, u.email, r.name as role from users u left join roles on u.role=r.id 
# create a postgres user
create user api with password 'netlab';
grant all privileges on lall tables in schema public to api;
\q # exit from the postgres console
psql -h 127.0.0.1 -U api api
# provide 'netlab' password
```
Were not mentioned:
- [Create sequence](https://postgrespro.ru/docs/postgresql/9.6/sql-createsequence)  
- [PostgreSQL foreign key](https://www.postgresqltutorial.com/postgresql-foreign-key)  

### Golang installing
```bash
# for each user it's worth installing its own golang
wget https://dl.google.com/go/go1.11.2.linux-amd64.tar.gz
tar zxvf go1.11.2.linux-amd64.tar.gz

vi .bashrc
    export PATH=$PATH:/root/go/bin
    export GOROOT=/root/go
    export GOPATH='pwd'

go get # download github located module
go version
```
#### Golang code
```bash
vi main.go
package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    "fmt"   
)

func main() {
    fmt.Println("Hello, World!")

    dbinfo := fmt.SPrintf("host=127.0.0.1 user=api password=netlab dbname=api sslmode=disable") # Sprintf writes a line

    db, err :=sql.Open("postgres", dbinfo)
    if err != nil {
        panic(err)
    }

    fmt.Println("# Querying")
    rows, err := db.Quesry("SELECT 1")
    if err != nil {
        panic(err)
    }

    for rows.Next() {
        var one int
        err = rows.Scan(&one)
        if err != nil {
            panic(err)
        }
        fmt.Printf("%d\n", one)
    }

    defer db.Close()
}
```
- [Golang: Package sql](https://golang.org/pkg/database/sql)




## Acknowledgments

This repo was inspired by [rebrainme.com](https://rebrainme.com) team

## See Also
- [PostgreSQL packages for Debian](https://wiki.postgresql.org/wiki/Apt)
- [REST API Tutorial](https://www.restapitutorial.com/lessons/httpmethods.html)
- []()

## License
Follow Ansible, APIDOC, NGINX, Terraform, and other licenses terms and conditions.
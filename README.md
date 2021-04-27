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
sudo apt install postgresql
sudo vi /etc/postgresql/11/main/postgresql.conf

```

### PostgreS console
```bash
sudo -u postgres psql -c 'create database api;' # create a database
sudo -u postgres psql postgres # enter into the postgres console as a postgres user
\c api # connecting to the api database
# create a postgres user
create user api with password 'netlab';
AlTER DATABASE api OWNER TO api;
grant all privileges on all tables in schema public to api;

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
select users.name, users.email, roles.name as role from users left join roles on users.role=roles.id; 
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

vi .bash_profile
    export PATH=$PATH:/Users/makuznet/go/bin
    export GOROOT=/Users/makuznet/go
    export GOPATH='/Users/makuznet/Documents/rebrain/api'

go get github.com/lib/pq # download github located module
go mod init pkg/mod # relative path for a module to get initialized
```

- [Golang: Package sql](https://golang.org/pkg/database/sql)
- [Golang: Package http](https://golang.org/pkg/net/http)



## Acknowledgments

This repo was inspired by [rebrainme.com](https://rebrainme.com) team

## See Also
- [PostgreSQL packages for Debian](https://wiki.postgresql.org/wiki/Apt)
- [REST API Tutorial](https://www.restapitutorial.com/lessons/httpmethods.html)
- [Connecting to a PostgreSQL](https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/)
- []()

## License
Follow Ansible, APIDOC, NGINX, Terraform, and other licenses terms and conditions.
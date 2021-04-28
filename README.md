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

# create table Products
create table products (id int PRIMARY KEY, title varchar(128), price int, description varchar(1024), category varchar(32), image varchar(128));
# PRIMARY KEY — unique id for rows in the table

# alternatively primary key can be added after a table is created
alter table products add primary key(id);

# grant privileges to the user
grant all on TABLE products to api;

# create sequence to increase records id 
create sequence products_id_seq;
select nextval('products_id_seq');
alter table products alter column id set default nextval('products_id_seq');
grant all on sequence products_id_seq to api;

# populate products table
insert into products values(1,'Fjallraven - Foldsack No. 1 Backpack, Fits 15 Laptops',109.95,'Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday','men clothing','https://fakestoreapi.com/img/81fPKd-2AYL._AC_SL1500_.jpg');


\q # exit from the postgres console

# check that Postgres can accept local connections
psql -h 127.0.0.1 -U api api
# provide 'netlab' password
```

Were not mentioned:
- [Create sequence](https://postgrespro.ru/docs/postgresql/9.6/sql-createsequence)  
- [PostgreSQL foreign key](https://www.postgresqltutorial.com/postgresql-foreign-key) 

## Golang
### HTTP function
```bash
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!\n", r.URL.Path[1:])
}
```
%s is r.URL.Path[1:];
r in r.URL.Path[1:] is the var containing the whole url from the request, i.e. 'curl -D - -s http://127.0.0.1:8080/makuznet', will result in extracting 'makuznet' from the url path ([1:] means provide what is written after the first sign of url path) and putting it in to the answ er 'Hi there, I love makuznet!' instead of %s var.

### Fprintf
```bash
fmt.Fprintf(w, "%d\n %s\n %d\n %s\n %s\n %s\n", id, title, price, description, category, image)
```
w — write to the http listener;
%d\n — print integer (%d) and add a new line (\n);
id, title, etc — vars, which come from a request to the Postgres database;

 

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

### Check if an json string is correct 
```bash
curl -D -s http://127.0.0.1:8080/v1/products/ | python -m json.tool
```

## Acknowledgments

This repo was inspired by [rebrainme.com](https://rebrainme.com) team

## See Also
- [PostgreSQL packages for Debian](https://wiki.postgresql.org/wiki/Apt)
- [REST API Tutorial](https://www.restapitutorial.com/lessons/httpmethods.html)
- [Connecting to a PostgreSQL](https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/)
- []()

## License
Follow Ansible, APIDOC, NGINX, Terraform, and other licenses terms and conditions.
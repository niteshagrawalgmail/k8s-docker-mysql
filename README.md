# k8s-docker-mysql

### To run mysql on docker and then ssh to mysql command prompt 

docker run --name mysqldb -e MYSQL_ROOT_PASSWORD=password -d -p 3306:3306 mysql:5.7

docker exec -it mysqldb /bin/bash

mysql -u root -p

### Useful MySQL commands :

CREATE SCHEMA commerce;

CREATE TABLE customers (
	id int NOT NULL AUTO_INCREMENT,
	first_name varchar(255),
	last_name varchar(255),
	PRIMARY KEY (id)
);

INSERT INTO customers (id, first_name, last_name) VALUES ('1','Nitesh','Agrawal');

OR 

INSERT INTO customers (first_name, last_name) VALUES ('foo','bar');

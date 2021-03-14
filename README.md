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

### Docker commands

Command to build the docker image 

Navigate to the root folder and run the below command

```
docker build -t goapp .
```

Command to list all the containers

```
docker container ls -a
```

To list all the running containers 

```
docker container ls
```

Command to list all the images

```
 docker images
```

Command to run the goapp

```
docker run --name myapp -p 3000:3000 goapp:latest
```

Commands to push local images to docker hub

```
https://ropenscilabs.github.io/r-docker-tutorial/04-Dockerhub.html
```

### Git commands

git push

```
git push origin main:main
```

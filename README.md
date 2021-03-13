# k8s-docker-mysql

### To run mysql on docker and then ssh to mysql command prompt 

docker run --name mysqldb -e MYSQL_ROOT_PASSWORD=password -d -p 3306:3306 mysql:5.7

docker exec -it mysqldb /bin/bash

mysql -u root -p


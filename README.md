# API Server

## Usage

- launch local server

```
docker-compose up
```

- login to local mysql container

```sh
docker exec -it api-server_mysql_1 bash
mysql -u root -p
# enter the root password written in docker-compose.yml
```

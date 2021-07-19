harvest is the server that communicates with main DB, storage and stripe endpoints of flyvase.

## Usage

- launch local server

```
docker-compose up
```

- login to local mysql container

```sh
docker exec -it harvest_mysql_1 bash
mysql -u root -p
# enter the root password written in docker-compose.yml
```

- deploy to app engine

```
gcloud app deploy --project YOUR_PROJECT_ID
```

harvest is the server that communicates with main DB, storage and stripe endpoints of flyvase.

## Setup

- build docker image

```
docker build -t harvest:latest .
```

- run docker container (This will launch your server)

```
docker run -it --name harvest --mount type=bind,src=${PWD},dst=/app -p 8080:8080 harvest:latest
```

## Usage

- launch local server

```
docker start -i harvest
```

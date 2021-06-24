harvest is the server that communicates with main DB, storage and stripe endpoints of flyvase.

## Setup

- build docker image

```
docker build -t harvest:latest .
```

- launch docker container

```
docker run -it --name harvest --mount type=bind,src=${PWD},dst=/app -p 8080:8080 harvest:latest
```

- install linter

```
go install honnef.co/go/tools/cmd/staticcheck@latest
```

## Usage

- update dependencies

```
# inside the docker container
go mod download
```

- launch local server

```
# inside the docker container
go run main.go
```

- run static analysis

```
staticcheck ./...
```

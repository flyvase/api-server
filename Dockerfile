FROM golang:1.15.13-alpine

WORKDIR /app

RUN apk update && apk add git && apk add mysql-client
COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8080

CMD [ "/bin/ash" ]
FROM golang:1.15.11-alpine

WORKDIR /app

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8080

CMD [ "/bin/ash" ]
FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download
RUN apt-install bash
RUN  curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash
RUN apt-get update
RUN apt-get install migrate
ENTRYPOINT go run cmd/lo_project/main.go
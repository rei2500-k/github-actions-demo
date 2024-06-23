FROM golang:1.22-bullseye

RUN apt-get update -y

EXPOSE 8080

WORKDIR /var/www

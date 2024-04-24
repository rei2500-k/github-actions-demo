FROM golang:1.22-bullseye

RUN apt-get update -y

WORKDIR /var/www

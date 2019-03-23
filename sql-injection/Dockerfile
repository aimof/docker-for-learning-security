FROM ruby:2.6.2-slim-stretch

LABEL maintainer="aimof (aimof.aimof@gmail.com)"

RUN apt-get update && \
    apt-get upgrade && \
    apt-get -y install build-essential && \
    apt-get clean && \
    mkdir /app
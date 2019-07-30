FROM nginx:latest
MAINTAINER Inno Jia <jdgets111@gmail.com>

RUN apt-get update && apt-get install -y curl
COPY ./dist /usr/share/nginx/html
COPY nginx/default.conf /etc/nginx/conf.d/default.conf

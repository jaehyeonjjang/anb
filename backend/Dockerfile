FROM       alpine
MAINTAINER missbsd@gmail.com

RUN mkdir -p /usr/local/backend/config
RUN mkdir -p /usr/local/backend/webdata
RUN mkdir -p /usr/local/front

COPY ./bin/backend.linux /usr/local/backend/backend
COPY ./config/config.json /usr/local/backend/config/config.json
ADD ./front/www /usr/local/front/www
ADD ./template /usr/local/backend/template

WORKDIR /usr/local/backend
CMD    ./backend

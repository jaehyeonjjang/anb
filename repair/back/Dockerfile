FROM       alpine
MAINTAINER missbsd@gmail.com

RUN  mkdir -p /usr/local/main
COPY ./bin/main.linux /usr/local/main/main
COPY ./.env.yml /usr/local/main/.env.yml
ADD ./dist /usr/local/main/web
ADD ./doc /usr/local/main/doc
RUN  mkdir -p /usr/local/main/webdata
RUN  mkdir -p /usr/local/main/webdata/blueprint
RUN  mkdir -p /usr/local/main/webdata/excel
RUN  mkdir -p /usr/local/main/webdata/file
RUN  mkdir -p /usr/local/main/webdata/temp
RUN  mkdir -p /usr/local/main/webdata/patrol
RUN  mkdir -p /usr/local/main/webdata/periodic
RUN  mkdir -p /usr/local/main/webdata/periodicresult
RUN  mkdir -p /usr/local/main/webdata/apk
RUN  mkdir -p /usr/local/main/webdata/area

COPY ./doc/empty.jpg /usr/local/main/webdata/empty.jpg

WORKDIR /usr/local/main
CMD    ./main --mode production

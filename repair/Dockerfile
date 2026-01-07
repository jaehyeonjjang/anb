FROM golang:1.19 AS gobuilder

WORKDIR /build
COPY ./back /build
RUN go build -a -ldflags '-s' -o main main.go

FROM       ubuntu
MAINTAINER missbsd@gmail.com

RUN  mkdir -p /usr
RUN  mkdir -p /usr/local
RUN  mkdir -p /usr/local/main
COPY --from=gobuilder ./build/main /usr/local/main/main
COPY ./front/dist /usr/local/main/web
COPY ./back/.env.yml /usr/local/main/.env.yml
ADD ./back/doc /usr/local/main/doc
COPY ./back/doc/empty.jpg /usr/local/main/webdata/empty.jpg
RUN  mkdir -p /usr/local/main/webdata
RUN  mkdir -p /usr/local/main/webdata/blueprint
RUN  mkdir -p /usr/local/main/webdata/excel
RUN  mkdir -p /usr/local/main/webdata/file
RUN  mkdir -p /usr/local/main/webdata/temp
RUN  mkdir -p /usr/local/main/webdata/patrol
RUN  mkdir -p /usr/local/main/webdata/periodic
RUN  mkdir -p /usr/local/main/webdata/periodicresult
RUN  mkdir -p /usr/local/main/webdata/apk

WORKDIR /usr/local/main
CMD    ./main --mode production

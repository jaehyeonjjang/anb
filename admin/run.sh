#!/usr/bin/env bash

PID=`ps -ef | grep gradlew | grep -v grep | awk '{print $2}'`
kill -9 $PID
./gradlew bootRun -Pprofile=default &

fswatch --follow-links --recursive -o src | while read num; do
    echo "File change!!!"

	PID=`ps -ef | grep gradlew | grep -v grep | awk '{print $2}'`
	kill -9 $PID
	./gradlew bootRun -Pprofile=default &
done


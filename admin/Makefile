tag=latest

all: run

run:
	./gradlew bootRun -Pprofile=default

allrun:
	fswatch -0 src | ./run.sh &
	./gradlew bootRun -Pprofile=default

build: dummy
	./gradlew build -Pprofile=default

clean: dummy
	./gradlew clean -Pprofile=default

docker: build
	docker build -t yiyuhki/anbadmin:$(tag) .

dockerrun: docker
	docker run -d --name="anbadmin" -p 9000:9000 yiyuhki/anbadmin

push: docker
	docker push yiyuhki/anbadmin:$(tag)

dummy:

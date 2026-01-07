tag=latest

all: server

build: dummy
	go build -o bin/main main.go

frontrun:
	cd ../front && make run

test: dummy
	go test -v ./...

linux:
	env GOOS=linux GOARCH=amd64 go build -o bin/main.linux main.go

dockerbuild:
	GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o bin/main.linux main.go
	cd ../front && npm run release && cp -rf dist ../back/

docker: dockerbuild
	docker build --no-cache -t netb.co.kr:5000/repair:$(tag) .

dockerrun:
	docker run -d --name="repair" -p 9108:9108 netb.co.kr:5000/repair

push: docker
	docker push netb.co.kr:5000/repair:$(tag)

deploy: push
	docker-compose --context repair pull
	docker-compose --context repair up -d

localdeploy: 
	go build -a -ldflags '-s' -o bin/main.linux main.go
	cd ../front && npm run release && cp -rf dist ../back/
	docker build -t netb.co.kr:5000/repair:$(tag) .
	docker push netb.co.kr:5000/repair:$(tag)
	docker-compose pull
	docker-compose up -d

clean:
	rm -f bin/main

dummy:

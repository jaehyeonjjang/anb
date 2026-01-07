tag=latest

all: 

dockerbuild:
	cd front && pnpm install && pnpm run release

docker: dockerbuild
	docker build -t netb.co.kr:5000/repair:$(tag) . --no-cache

dockerrun:
	docker run -d --name="repair" -p 9108:9108 netb.co.kr:5000/repair

push: docker
	docker push netb.co.kr:5000/repair:$(tag)

deploy: push
	docker-compose --context repair pull
	docker-compose --context repair up -d

localdeploy: 
	git pull
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o bin/main.linux main.go
	cd ../front && npm run release && cp -rf dist ../back/
	docker build -t netb.co.kr:5000/repair:$(tag) .
	docker push netb.co.kr:5000/repair:$(tag)
	docker-compose pull
	docker-compose up -d

clean:
	rm -f bin/main

dummy:

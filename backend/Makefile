tag=latest

all: server

server: dummy
	build/model.php build
	build/router.php build
	go build -o bin/backend main.go

fswatch:
	fswatch -0 controllers | xargs -0 -n1 build/notify.sh

run:
	fswatch -0 controllers | xargs -0 -n1 build/notify.sh &
	gin --port 80 --bin bin/backend run main.go

test: dummy
	#mysql -u root anb < build/test.sql
	go test -v ./...

linux:
	env GOOS=linux GOARCH=amd64 go build -o bin/backend.linux main.go

windows:
	env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -o bin/backend.exe main.go

docker: linux
	cp -rf ../front/www ./front/
	docker build -t yiyuhki/backend:$(tag) .

dockerrun: docker
	docker run -d --name="backend" -p 3001:3001 yiyuhki/backend:$(tag)

push: docker
	docker push yiyuhki/backend:$(tag)

clean:
	rm -f bin/backend
	rm -f bin/backend.linux

dummy:

up:
	docker-compose -f ./deploy/docker-compose.yml up -d

build:
	docker-compose build -f ./deploy/docker-compose.yml

go-build:
	cd ./srv
	go get -v .
	go build .

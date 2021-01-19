up:
	docker-compose -f ./deploy/docker-compose.yml up -d

up:
	docker-compose -f ./deploy/docker-compose.yml down

build:
	docker-compose -f ./deploy/docker-compose.yml build

go-build:
	cd ./srv
	go get -v .
	go build .
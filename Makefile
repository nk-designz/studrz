up:
	docker-compose -f ./deploy/docker-compose.yml up -d

build:
	docker-compose -f ./deploy/docker-compose.yml build

go-build:
	cd ./srv
	go get -v .
	go build .

up:
	docker-compose up -f ./deploy/docker-compose.yml

docker-build:
	docker-compose build -f ./deploy/docker-compose.yml

go-build:
	cd ./srv
	go get -v .
	go build .

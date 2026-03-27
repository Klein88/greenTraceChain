VERSION=v3.0.0

build:
	cd src && go mod tidy && go build -o ../bin/chainmaker-explorer.bin

start: stop
	cd scripts && ./startup.sh

stop:
	cd scripts && ./shutdown.sh

docker-build:
	docker build -t explorer-backend -f ./docker/Dockerfile .
	docker tag explorer-backend chainmakerofficial/explorer-backend:${VERSION}

docker-build-dev: build
	docker build -t explorer-backend -f ./docker/dev.Dockerfile .
	docker tag explorer-backend explorer-backend:${VERSION}

docker-compose-start: docker-compose-stop
	cd docker && docker-compose up -d

docker-compose-stop:
	cd docker && docker-compose down

docker-start:
	cd docker && ./start-docker-dev.sh

docker-stop:
	docker stop explorer-backend

all: docker.start test.integration docker.stop

docker.start:
	docker-compose up -d
	sleep 5

docker.stop:
	docker-compose down

docker.restart: docker.stop docker.start
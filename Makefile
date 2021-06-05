build:
	docker build -t file_sharer:latest -f docker/Dockerfile .
run:
	docker-compose -f docker/docker-compose.yml up
clean:
	docker stop file_sharer_db file_sharer
	docker rm docker_db_data file_sharer_db file_sharer
all: clean build run
build:
	docker build -t file_sharer:latest -f docker/Dockerfile .
run:
	docker-compose -f docker/docker-compose.yml up
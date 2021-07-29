build:
	docker build -t file_sharer:latest -f docker/Dockerfile .
run:
	docker-compose -f docker/docker-compose.yml up
clean:
	docker stop file_sharer_db || true && docker rm file_sharer_db || true
	docker stop file_sharer || true && docker rm file_sharer || true
	docker volume rm docker_db_data || true
generate:
	cd internal/provider/
	wire
	cd -
all: clean generate build run
build:
	docker build -t file_sharer:latest .
run:
	docker-compose -p file_sharer -f docker-compose.yml up
clean:
	docker stop file_sharer ; docker rm file_sharer || true
	docker stop file_sharer_db_1; docker rm file_sharer_db_1 || true
	docker volume rm file_sharer_db_data || true
clean_migrations:
	docker rm file_sharer_migrations
generate:
	cd internal/provider/ && wire ; cd ../..
all: clean build run

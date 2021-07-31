build:
	docker build -t file_sharer:latest -f docker/Dockerfile .
run:
	docker-compose -f docker/docker-compose.yml up
clean:
	docker stop file_sharer_db || true && docker rm file_sharer_db || true
	docker stop file_sharer || true && docker rm file_sharer || true
	docker volume rm docker_db_data || true
generate:
	cd internal/provider/ && wire && cd ../..
sync:
	ssh dev2 "mkdir -p ./.vscode_proj && mkdir -p ./.vscode_proj/Filesharer"
	rsync -avzh --delete ./ dev2:/home/ubuntu/.vscode_proj/Filesharer
allDev: sync
	ssh dev2 "cd ./.vscode_proj/Filesharer && make all"
all: clean build run

include ./.env
export

# Start docker
start:
	docker-compose up -d
# Stop docker
stop:
	docker-compose down
# Set Env
all:
	env
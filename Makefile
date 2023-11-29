
local:
	echo "Starting local environment"
	docker compose up -d

down-local:
	docker compose down

run:
	./bin/load_ddl.sh
	go run .

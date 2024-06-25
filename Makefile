jump-consumer:
	docker-compose exec consumer bash

docker-up:
	docker compose up --build --force-recreate
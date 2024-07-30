jump-consumer:
	docker compose exec consumer bash

jump-producer:
	docker compose exec producer bash

docker-up:
	docker compose up --build --force-recreate
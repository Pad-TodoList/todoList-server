CONFIG_FILE=./config/docker/docker-compose.yml


start:
	echo $(CONFIG_FILE)
	docker-compose -f $(CONFIG_FILE) up --build -d

stop:
	docker-compose -f $(CONFIG_FILE) down

restart:
	docker-compose -f $(CONFIG_FILE) down
	docker-compose -f $(CONFIG_FILE) up --build -d

.PHONY: start stop restart

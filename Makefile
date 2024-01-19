dc-reset-to-factory:
	- docker stop $$(docker ps -a -q)
	- docker kill $$(docker ps -q)
	- docker rm $$(docker ps -a -q)
	- docker rmi $$(docker images -q)
	- docker system prune --all --force --volumes

dcup-dev:
	docker-compose up

dcup-build:
	docker-compose build

dcup-prod:
	docker-compose -f ./docker-compose.prod.yml up

dcup-prod-build:
	docker-compose -f ./docker-compose.prod.yml build

dc-down:
	docker-compose down

dc-clear:
	docker-compose down
	docker rmi -f $(docker images | grep doszy)

hosts:
	sudo -- sh -c "echo 127.0.0.1  graph-dev.com >> /etc/hosts"

hosts-dev-prod:
	sudo -- sh -c "echo 127.0.0.1  graph.com >> /etc/hosts"

rm-hosts:
	sudo -- sh -c "sed -i -e '/127.0.0.1 graph-dev.com/d' /etc/hosts"

rm-hosts-dev-prod:
	sudo -- sh -c "sed -i -e '/127.0.0.1 graph.com/d' /etc/hosts"
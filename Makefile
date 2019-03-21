PROJECT = go_api

.PHONY: start
start:
	docker-compose -p $(PROJECT) up -d

.PHONY: restart
restart:
	docker-compose -p $(PROJECT) kill && \
	docker-compose -p $(PROJECT) rm -f && \
	docker-compose -p $(PROJECT) up -d --build

.PHONY: kill
kill:
	docker-compose -p $(PROJECT) kill

.PHONY: down
down:
	docker-compose -p $(PROJECT) down

.PHONY: ps
ps:
	docker-compose -p $(PROJECT) ps

.PHONY: proto
proto:
	cd proto && \
	docker run -v `pwd`:/defs namely/protoc-all -f helloworld.proto -l go -o go

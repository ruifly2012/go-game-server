v ?= latest

.PHONY: build
build:
	make -C arena-service build
	make -C game-service build

.PHONY: proto
proto:
	protoc -I=./proto --go_out=:${GOPATH}/src --micro_out=:${GOPATH}/src ./proto/game/game.proto
	protoc -I=./proto --go_out=:${GOPATH}/src --micro_out=:${GOPATH}/src ./proto/client/client.proto
	protoc -I=./proto --go_out=:${GOPATH}/src --micro_out=:${GOPATH}/src ./proto/pubsub/pubsub.proto

.PHONY: docker
docker:
	make -C apps/game docker

.PHONY: test
test:
	make -C apps/game test

.PHONY: run
run:
	docker-compose up -d

.PHONY: push
push:
	docker tag ultimate hellodudu86/yokai_server_game:$(v)
	docker push hellodudu86/yokai_server_game:$(v)

.PHONY: clean
clean:
	docker rm -f $(shell docker ps -a -q)

.PHONY: stop
stop:
	docker-compose down

.PHONY: init
init:
	go mod download

.PHONY: grpc
grpc:
	cd ./services/langchain && goctl rpc protoc ./rpc/main.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style=go_zero --localize

.PHONY: api
api:
	goctl api go --api=./services/gateway/api/main.api --dir=./services/gateway --style=go_zero --localize

.PHONY: dc-up
	docker-compose -f docker-compose.yml -f docker-compose-nw.yml -f docker-compose-gw.yml up -d

.PHONY: dc-down
	docker-compose -f docker-compose.yml -f docker-compose-nw.yml -f docker-compose-gw.yml down
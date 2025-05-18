PROTO_DIR=proto

gen-proto:
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=paths=source_relative:$(PROTO_DIR)/accounts/accountspb \
		--go-grpc_out=paths=source_relative:$(PROTO_DIR)/accounts/accountspb \
		$(PROTO_DIR)/accounts/accounts.proto

	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=paths=source_relative:$(PROTO_DIR)/users/userspb \
		--go-grpc_out=paths=source_relative:$(PROTO_DIR)/users/userspb \
		$(PROTO_DIR)/users/users.proto
		
build-accounts:
	docker build -t block-lytics/accounts ./accounts

run-accounts:
	docker run -p 50051:50051 block-lytics/accounts

build-gateway:
	docker build -t block-lytics/gateway ./gateway

run-gateway:
	docker run -p 8080:8080 block-lytics/gateway

up:
	docker-compose up --build

PHONY: gen-proto
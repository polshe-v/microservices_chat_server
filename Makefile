ifneq ($(ENV),)
	include $(ENV).env
endif

LOCAL_BIN:=$(CURDIR)/bin
BINARY_NAME=main
CONFIG=$(ENV).env
LOCAL_MIGRATION_DIR=$(MIGRATION_DIR)
LOCAL_MIGRATION_DSN="host=localhost port=$(POSTGRES_PORT_LOCAL) dbname=$(POSTGRES_DB) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) sslmode=disable"

# #################### #
# DEPENDENCIES & TOOLS #
# #################### #

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.1
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.32.0
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.18.0
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@v3.3.6

get-protoc-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

lint:
	GOBIN=$(LOCAL_BIN) bin/golangci-lint run ./... --config .golangci.pipeline.yaml

generate-api:
	make generate-api-v1

generate-api-v1:
	mkdir -p pkg/chat_v1
	protoc --proto_path api/chat_v1 \
	--go_out=pkg/chat_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/chat_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/chat_v1/chat.proto

generate-mocks:
	go generate ./internal/repository
	go generate ./internal/service
	go generate ./internal/client/db

unit-test:
	go test ./internal/api/chat/tests -v
	go test ./internal/service/chat/tests -v

check-env:
ifeq ($(ENV),)
	$(error No environment specified)
endif

# ##### #
# BUILD #
# ##### #

build-app:
	GOOS=linux GOARCH=amd64 go build -o $(LOCAL_BIN)/${BINARY_NAME} cmd/chat/main.go

docker-build: docker-build-app docker-build-migrator

docker-build-app: check-env
	docker buildx build --no-cache --platform linux/amd64 -t chat-server:${APP_IMAGE_TAG} --build-arg="ENV=${ENV}" --build-arg="CONFIG=${CONFIG}" .

docker-build-migrator: check-env
	docker buildx build --no-cache --platform linux/amd64 -t migrator-chat:${MIGRATOR_IMAGE_TAG} -f migrator.Dockerfile --build-arg="ENV=${ENV}" .

# ###### #
# DEPLOY #
# ###### #

docker-deploy: check-env docker-build
	docker compose --env-file=$(ENV).env up -d

local-migration-status: check-env
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up: check-env
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down: check-env
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

# #### #
# STOP #
# #### #

docker-stop: check-env
	docker compose --env-file=$(ENV).env down

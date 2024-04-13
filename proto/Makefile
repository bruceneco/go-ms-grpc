# Variables
RELEASE_VERSION := $(word 1,$(MAKECMDGOALS))
FOLDER_NAMES := order shipping payment

# Rules
.PHONY: setup
setup:
	git fetch --all && git checkout main && git pull origin main
	sudo apt-get install -y protobuf-compiler golang-goprotobuf-dev
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: generate
generate:
	mkdir -p ./go
	protoc --go_out=./go --go_opt=paths=source_relative \
	  --go-grpc_out=./go --go-grpc_opt=paths=source_relative \
	 ./${SERVICE_NAME}/*.proto
	cd go/${SERVICE_NAME} && \
	go mod init \
	  github.com/bruceneco/go-ms-grpc/proto/go/${SERVICE_NAME} || true && \
	go mod tidy

.PHONY: commit
commit:
	git add . && git commit -am "gen(proto): update ${SERVICE_NAME} to ${RELEASE_VERSION}" || true
	git tag -fa proto/go/${SERVICE_NAME}/${RELEASE_VERSION} \
	  -m "proto/go/${SERVICE_NAME}/${RELEASE_VERSION}"

# Default target
.DEFAULT_GOAL := run

.PHONY: run
run: setup
	@for SERVICE_NAME in $(FOLDER_NAMES); do \
    		make generate SERVICE_NAME=$$SERVICE_NAME; \
    		make commit SERVICE_NAME=$$SERVICE_NAME; \
    done
	git push origin refs/tags/proto/go/*/${RELEASE_VERSION}
	git push origin HEAD:main



.PHONY: install test build serve clean pack deploy ship

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get ./...

test: install
	go test ./...

proto-gen:
	 protoc --micro_out=. --go_out=plugins=grpc:. ./proto/*.proto
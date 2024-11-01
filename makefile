# Determine this makefile's path.
# Be sure to place this BEFORE `include` directives, if any.
THIS_FILE := $(lastword $(MAKEFILE_LIST))
BINARY_NAME = go-rpc-todo-app.out
OUT_PATH = bin


compile:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative modules/**/*.proto

build:
	mkdir -p bin
	@$(MAKE) -f $(THIS_FILE) compile
	go build -o ${OUT_PATH}/${BINARY_NAME} main.go

run:
	./${BINARY_NAME}

dev:
	air

clean:
	go clean
	rm ${OUT_PATH}/${BINARY_NAME}
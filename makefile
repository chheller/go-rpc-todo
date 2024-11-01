# Determine this makefile's path.
# Be sure to place this BEFORE `include` directives, if any.
THIS_FILE := $(lastword $(MAKEFILE_LIST))
CERT_PATH = x509
BINARY_NAME = go-rpc-todo-app.out
OUT_PATH = bin
SHELL := /bin/zsh

generate-certs:
	rm -rf ${CERT_PATH}
	mkdir -p ${CERT_PATH}
	openssl req -x509 -out ${CERT_PATH}/localhost.crt -keyout ${CERT_PATH}/localhost.key -newkey rsa:2048 -nodes -sha256 -subj '/CN=localhost' -extensions EXT -config <(  printf "[dn]\nCN=localhost\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:localhost\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")

protoc-compile:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative modules/**/*.proto

build:
	mkdir -p bin
	@$(MAKE) -f $(THIS_FILE) protoc-compile
	go build -o ${OUT_PATH}/${BINARY_NAME} main.go

run:
	.${OUT_PATH}/${BINARY_NAME}

dev:
	air --build.cmd "make protoc-compile && go build -o ${OUT_PATH}/${BINARY_NAME} main.go " --build.bin "${OUT_PATH}/${BINARY_NAME}" --build.include_ext "go,proto" --build.exclude_regex ".*pb\.go"

clean:
	go clean
	rm -rf ${OUT_PATH}

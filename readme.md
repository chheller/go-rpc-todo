# Setup
## Setup Turso connection
### Network DB
- [Install Turso](https://docs.turso.tech/quickstart)
- Run `turso auth login` or `turso auth login --headless` if using WSL
- Run `turso db list`
  
## Prequisite software
This project assumes you use the universal version manager [asdf-vm](https://asdf-vm.com/guide/getting-started.html) 
## Install dependencies
This repository uses [air](), [protoc golang plugins](), and [grpcurl](https://github.com/fullstorydev/grpcurl) and can be installed by makefile target. Likewise, all required tooling can be installed by `asdf-vm`

```sh
make install-dependencies

# chheller in go-rpc-todo on  main [!?] 
# ❯ make install-dependencies
# make[1]: Entering directory '/home/chheller/projects/go-rpc-todo'
# asdf install
# bun 1.1.27 is already installed
# golang 1.23.2 is already installed
# java zulu-17.52.19 is already installed
# nodejs 18.7.0 is already installed
# protoc 28.3 is already installed
# make[1]: Leaving directory '/home/chheller/projects/go-rpc-todo'
# make[1]: Entering directory '/home/chheller/projects/go-rpc-todo'
# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# make[1]: Leaving directory '/home/chheller/projects/go-rpc-todo'
# make[1]: Entering directory '/home/chheller/projects/go-rpc-todo'
# go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
# make[1]: Leaving directory '/home/chheller/projects/go-rpc-todo'
# make[1]: Entering directory '/home/chheller/projects/go-rpc-todo'
# go install github.com/air-verse/air@latest
# make[1]: Leaving directory '/home/chheller/projects/go-rpc-todo'
# make[1]: Entering directory '/home/chheller/projects/go-rpc-todo'
# asdf reshim golang
# make[1]: Leaving directory '/home/chheller/projects/go-rpc-todo'
```


# Running the application
## Generate local certificates
```sh
make generate-certs
```
## Development
```sh
make dev
```
## Build and Run
```sh
make build
make run
```
# Testing

## Test an RPC endpoint
From there, start the application and run the health check

```sh
 grpcurl -proto helloworld/helloworld.proto -cacert x509/localhost.crt -d '{"name": "John Halo" }' localhost:8080 helloworld.Greeter/Say
```
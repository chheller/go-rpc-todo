# Setup
## Setup Turso connection
### Network DB
- [Install Turso](https://docs.turso.tech/quickstart)
- Run `turso auth login` or `turso auth login --headless` if using WSL
- Run `turso db list`
  
## Create self signed certificates
run `make generate-certs` to generate local, self-signed certs in the expected path under `x509/`
## Setup gRPC

Follow the instructions at https://grpc.io/docs/languages/go/quickstart/

### Install Protobuf Compiler
```sh
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO $PB_REL/download/v28.3/protoc-28.3-linux-x86_64.zip
unzip protoc-28.3-linux-x86_64.zip -d $HOME/.local/protoc
export PATH="$PATH:$HOME/.local/protoc/bin"
protoc --version  # Ensure compiler version is 3+
```
## Compile the protobufs
Use the make file `make compile-protobuf` to compile all the protobufs. This is called automatically by `make build`
### Plugins
Ensure you've installed the protoc plugins to compile to Golang

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# If using asdf, run asdf reshim golang instead
export PATH="$PATH:$(go env GOPATH)/bin"
```
# Running the application
Use the makefile scripts to build and run the application. For development, use `make dev` to run air, the filesystem watcher for go. Otherwise, use `make build` and `make run`. 
# Testing
## Installing grpcurl
[Per the website](https://github.com/fullstorydev/grpcurl?tab=readme-ov-file#from-source), grpcurl can be installed via `go install` itself
```sh
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
# If using asdf, run 
asdf reshim golang
```

## Test an RPC endpoint
From there, start the application and run the health check

```sh
 grpcurl -proto helloworld/helloworld.proto -cacert x509/localhost.crt -d '{"name": "John Halo" }' localhost:8080 helloworld.Greeter/Say
```
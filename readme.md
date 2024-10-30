# Setup
## Setup Turso connection
### Network DB
- [Install Turso](https://docs.turso.tech/quickstart)
- Run `turso auth login` or `turso auth login --headless` if using WSL
- Run `turso db list`
  
## Create self signed certificates
```sh
mkdir x509
openssl req -x509 -out x509/localhost.crt -keyout x509/localhost.key \
  -newkey rsa:2048 -nodes -sha256 \
  -subj '/CN=localhost' -extensions EXT -config <( \
   printf "[dn]\nCN=localhost\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:localhost\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")
```
## Setup gRPC /Follow the instructions at https://grpc.io/docs/languages/go/quickstart/
###Install Protobuf Compiler
```sh
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO $PB_REL/download/v28.3/protoc-28.3-linux-x86_64.zip
unzip protoc-28.3-linux-x86_64.zip -d $HOME/.local/protoc
export PATH="$PATH:$HOME/.local/protoc/bin"
protoc --version  # Ensure compiler version is 3+
```
### Plugins
Ensure you've installed the protoc plugins to compile to Golang

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# If using asdf, run asdf reshim golang instead
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Testing
### Installing grpcurl
[Per the website](https://github.com/fullstorydev/grpcurl?tab=readme-ov-file#from-source), grpcurl can be installed via `go install` itself
```sh
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
# If using asdf, run 
asdf reshim golang
```
From there, start the application and run the health check

```sh
 grpcurl -proto helloworld/helloworld.proto -cacert x509/localhost.crt -d '{"name": "John Halo" }' localhost:8080 helloworld.Greeter/Say
```
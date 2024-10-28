# Setup
## Setup Turso connection
### Network DB
- [Install Turso](https://docs.turso.tech/quickstart)
- Run `turso auth login` or `turso auth login --headless` if using WSL
- Run `turso db list`
- 
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
# Figure out how ASDF handles this particular thing 
export PATH="$PATH:$(go env GOPATH)/bin"

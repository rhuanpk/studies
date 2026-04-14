1. `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
1. `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
1. `go install github.com/favadi/protoc-go-inject-tag@latest`

1. `wget https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip`
1. `unzip protoc-3.15.8-linux-x86_64.zip -d /tmp/protoc-3.15.8`
1. `sudo mv /tmp/protoc-3.15.8/bin/protoc /usr/local/bin/`
1. `sudo mv /tmp/protoc-3.15.8/include/google/ /usr/local/include/`
1. `rm -rf /tmp/protoc-3.15.8 protoc-3.15.8-linux-x86_64.zip`
OR
1. `asdf plugin add protoc`
1. `asdf install protoc latest`
1. `asdf set -u protoc latest`

- `which protoc && protoc --version`
- `which protoc-gen-go && protoc-gen-go --version`
- `which protoc-gen-go-grpc && protoc-gen-go-grpc --version`
- `which protoc-go-inject-tag && protoc-go-inject-tag --version`

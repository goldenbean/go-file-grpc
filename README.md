# README

```bash
go mod init go-file-grpc
go mod tidy
go env
go mod vendor
go build -mod=vendor
go get -u github.com/golang/protobuf/{protoc-gen-go,proto} 

protoc --go_out=plugins=grpc:. proto/*.proto
```

[Protoc installation](https://grpc.io/docs/protoc-installation/)
[protobuf github](https://github.com/protocolbuffers/protobuf)

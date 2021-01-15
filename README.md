### Install Package
```
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/golang/protobuf/proto
```
```protobuf
protoc --go_out=plugins=grpc:./src .\src\blog\blogpb\blog.proto
```
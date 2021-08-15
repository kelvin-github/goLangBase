安装
go get -u google.golang.org/grpc
go get google.golang.org/protobuf/reflect/protoreflect

安装编译器(通用)
https://github.com/protocolbuffers/protobuf/releases/tag/v3.9.0

安装 插件
go get github.com/golang/protobuf/protoc-gen-go

命令行(产生 go 中间文件)
protoc --go_out=../services Prod.proto
protoc --go_out=plugins=grpc:../services Prod.proto
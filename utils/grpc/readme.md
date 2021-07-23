## GRPC使用指南

1. 下载grpc和代码生成工具protoc-gen-go

```shell
go get google.golang.org/grpc 
go install google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get google.golang.org/protobuf/reflect/protoreflect@v1.27.1
```



2. 下载protocol buffers插件

https://github.com/protocolbuffers/protobuf/releases

#### windows下载
`protoc-3.17.3-win64.zip`
解压后将bin目录protoc.exe 复制到 `C:\Windows\System32`

3.目录结构分为

```shell
grpc
  -- client
  -- server
  -- proto
```

+ proto

proto文件只需要撰写 *.proto文件
撰写好后cmd运行命令自动生成Go文件（这里我的proto文件叫message.proto，注意对应修改）

```shell
> protoc --go_out=. message.proto
> protoc --go-grpc_out=. message.proto
```



+ client

参照代码client\main.go, 具体移步官方库查看

+ server

参照代码server\main.go, 具体移步官方库查看

4. 分别运行客户端和服务端

  ```shell
  > ..clinet\go run main.go
  > ..server\go run main.go
  ```

### proto文件Demo

```protobuf
syntax = "proto3"; // 声明版本
package message;  // 指定包名

option go_package = ".;message";  // 指定golang包名

// 定义服务
service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// 请求包含的消息
message HelloRequest {
  string name = 1;
}

// 响应包含的消息
message HelloReply {
  string message = 1;
}
```
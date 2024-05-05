# grpcLearning
学习grpc的一些demo

在client 和server中定义了一个grpc客户端一个服务器

1.安装grpc库

go get google.golang.org/grpc



2.安装protoc编译器

在github链接下载合适的proto buf版本https://github.com/protocolbuffers/protobuf/releases

Linux：wget ***

下载好之后，解压，把bin目录下的文件，cp到GOPATH的bin目录下

然后 protoc --version 就能看到版本


3.安装go的编译器插件

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

（这两个插件在安装grpc的时候就下好了，这里是install）


好了之后，可以在gopath的bin目录下看到protoc-gen-go  protoc-gen-go-grpc


protoc --go_out=. --go-grpc_out=. ./proto/user.proto   编译写好的.proto文件


4.可以整一下vscode的vscode-proto-3这个插件，可以高亮.proto文件
package main

import (
	"context"
	"fmt"
	pb "grpcLearning/client/proto"
	"net"

	"google.golang.org/grpc"
)

// 实现一下sayHello方法
// 在生成的hello_grpc_pb.go 中可以看到，sayHello方法定义在结构体UnimplementedSayHelloServer中
type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, hr *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		ResponseMsg: "hello, " + hr.ReguestName + ".Welcome to use the grpc",
	}, nil
}

func main() {

	//开启端口监听
	ls, _ := net.Listen("tcp", "localhost:9090")

	//创建一个grpc服务
	grpcServer := grpc.NewServer()
	//然后，把实现好的sayHello方法注册到grpc连接中
	//注册的方法已经实现好了，直接调用register
	//注册的时候，一定是引用
	pb.RegisterSayHelloServer(grpcServer, &server{})

	// 然后，启动服务
	err := grpcServer.Serve(ls)
	if err != nil {
		fmt.Println("启动服务失败 ", err)
	}
}

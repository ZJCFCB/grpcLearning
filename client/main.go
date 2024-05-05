package main

import (
	"context"
	"fmt"
	pb "grpcLearning/client/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//创建一个与服务器的连接，后面的参数是可选的，表示是否加密
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("error", err)
	}
	defer conn.Close()

	//建立rpc链接
	client := pb.NewSayHelloClient(conn)
	//context.Background()表示当前上下文请求
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{ReguestName: "zjc"})
	//resp中定义好了获取其中变量的函数，在hell.pb.go中
	fmt.Println(resp.GetResponseMsg())
}

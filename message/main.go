package main

import (
	"context"
	"flag"
	"fmt"
	"gRPC-Go-Study/invoke"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

var (
	address = flag.String("addr", "127.0.0.1:8888", "The gRPC server address connect")
)

// 消息服务入口
func main() {
	flag.Parse()
	conn, err := grpc.Dial(*address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}

	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	// stub := api.NewIUserServiceClient(conn)
	stub := invoke.NewIInvokeJavaClient(conn)

	// 使用stub调用方法，设置goroutine调用超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// result, err := stub.Register(ctx, &api.UserRegisterRequest{Email: "wylu@zkjg.com", Username: "wylu", Age: 18})
	result, err := stub.InvokeCrossLanguage(ctx, &invoke.InvokeRequest{Identity: "我是go客户端", Target: "java服务端"})
	if err != nil {
		return
	}
	// fmt.Printf("Client receive data: %v, message: %v, code: %d", result.Data, result.Message, result.Code)
	fmt.Printf("Client receive result: %v, code: %d", result.Result, result.Code)
}

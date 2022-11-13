package main

import (
	"context"
	"fmt"
	"gRPC-Go-Study/user/api"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

// 用户服务入口
func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}
	grpcServer := grpc.NewServer()
	api.RegisterIUserServiceServer(grpcServer, &IUserService{})

	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("Fail to satrt grpc serve, error message = %v", err)
		return
	}
}

// IUserService 需要实现的服务
type IUserService struct {
	api.UnimplementedIUserServiceServer
}

// Register 实现服务内的Register接口
func (IUserService) Register(context context.Context, request *api.UserRegisterRequest) (*api.UserRegisterResponse, error) {
	receive := "收到注册信息，邮箱：" + request.Email + ", 用户名：" + request.Username + ", 年龄：" + strconv.FormatInt(int64(request.Age), 10)
	return &api.UserRegisterResponse{Data: "您已经成功注册，欢迎使用！", Code: 200, Message: receive}, nil
}

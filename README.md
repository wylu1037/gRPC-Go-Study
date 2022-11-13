# 🍱gRPC-Go-Study

## 🪶Prerequisite
### gRPC依赖
```shell
go get google.golang.org/grpc@latest
```
### go语言插件
该插件会根据`.proto`文件生成后缀为`.pb.go`的文件，包含所有`.proto`文件中定义的服务及消息。
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```

### gRPC插件
该插件会生成一个后缀为`_grpc.pb.go`的文件，其中包括：
+ 一种接口类型(stub)，提供给Client调用在`.proto`文件中定义的`service`方法；
+ Service需要实现的接口类型；
```shell
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

### 🍹Validate
验证是否已经具备go的gRPC开发环境。
```shell
protoc --version
# 输出：libprotoc 3.21.9

protoc-gen-go --version
# 输出：protoc-gen-go.exe v1.28.1

protoc-gen-go-grpc --version
# 输出：protoc-gen-go-grpc 1.2.0
```

## 📃proto
示例
```protobuf
syntax = "proto3";

package user;
option go_package="gRPC-Go-Study/user/api";

service IUserService {
  rpc Register (UserRegisterRequest) returns (UserRegisterResponse);
}

message UserRegisterRequest {
  string email = 1;
  string username = 2;
  optional string address = 3;
  int32 age = 4;
}

message UserRegisterResponse {
  string data = 1;
  int32 code = 2;
  string message = 3;
}

```

执行命令生成go的源码文件。
```shell
protoc --go_out=../ --go-grpc_out=../  proto/HelloClient.proto
```
生成的go文件如果出现依赖标红，可执行以下命令。
```shell
go mod tidy
```

## 🍂Service
实现服务端
```go
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
```

## 🍃Client
客户端
```go
package main

import (
	"context"
	"flag"
	"fmt"
	"gRPC-Go-Study/user/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

var (
	address = flag.String("addr", "127.0.0.1:8080", "The gRPC server address connect")
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

	client := api.NewIUserServiceClient(conn)

	// 使用stub调用方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := client.Register(ctx, &api.UserRegisterRequest{Email: "wylu@zkjg.com", Username: "wylu", Age: 18})
	if err != nil {
		return
	}
	fmt.Printf("Client receive data: %v, message: %v, code: %d", result.Data, result.Message, result.Code)
}
```

## 🌿跨语言调用
go做客户端，java做服务端。
### 📃proto
```protobuf
syntax="proto3";

package invoke;
option go_package="gRPC-Go-Study/invoke";
option java_package="cn.edu.hust.grpc.api.user";

service IInvokeJava {
  rpc invokeCrossLanguage (InvokeRequest) returns (InvokeResponse);
}

message InvokeRequest {
  string identity = 1;
  string target = 2;
}

message InvokeResponse {
  string result = 1;
  int32 code = 2;
}

```

### Service
Java的启动类会扫描被注解`@GrpcService`修饰的service，加入到gRPC的服务中去。
```
@GrpcService
public class InvokeService extends IInvokeJavaGrpc.IInvokeJavaImplBase {
	private final Logger logger = LoggerFactory.getLogger(InvokeService.class);

	@Override
	public void invokeCrossLanguage(InvokeRequest request, StreamObserver<InvokeResponse> responseObserver) {
		logger.info("InvokeService 收到参数，identity = {}, target = {}", request.getIdentity(), request.getTarget());
		InvokeResponse response = InvokeResponse.newBuilder()
			.setCode(200)
			.setResult(String.format("identity = %s, target = %s", request.getIdentity(), request.getTarget()))
			.build();
		responseObserver.onNext(response);
		responseObserver.onCompleted();
	}
}
```

### Client
go
```go
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
```


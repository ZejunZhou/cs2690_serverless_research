// main.go
package user

import (
    "context"
    "log"
    "net/http"
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc"
    pb "github.com/delimitrou/DeathStarBench/tree/master/hotelReservation/services/user/proto" // 更新为生成的实际路径
)

func Run() {
    grpcEndpoint := "localhost:8086" // gRPC 服务地址

    mux := runtime.NewServeMux()
    opts := []grpc.DialOption{grpc.WithInsecure()}

    // 将 gRPC-Gateway 注册到 gRPC 服务
    err := pb.RegisterUserHandlerFromEndpoint(context.Background(), mux, grpcEndpoint, opts)
    if err != nil {
        log.Fatalf("Failed to start gRPC-Gateway: %v", err)
    }

    // 运行 HTTP 服务器，监听端口 8081
    if err := http.ListenAndServe(":8087", mux); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}

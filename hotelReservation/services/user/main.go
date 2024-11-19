// main.go
package user

import (
    "context"
    "log"
    "net/http"
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc"
    pb "github.com/delimitrou/DeathStarBench/tree/master/hotelReservation/services/user/proto" // 更新为生成的实际路径
    pb2 "github.com/delimitrou/DeathStarBench/tree/master/hotelReservation/services/reservation/proto"
    pb3 "github.com/delimitrou/DeathStarBench/tree/master/hotelReservation/services/recommendation/proto"
)

func Run() {
    userGrpcEndpoint := "localhost:8086"          // User gRPC 服务地址
    reservationGrpcEndpoint := "localhost:8087"  // Reservation gRPC 服务地址
    recommendationGrpcEndpoint := "localhost:8085"

    mux := runtime.NewServeMux()
    opts := []grpc.DialOption{grpc.WithInsecure()}

    // 注册 User 服务的 gRPC-Gateway
    err := pb.RegisterUserHandlerFromEndpoint(context.Background(), mux, userGrpcEndpoint, opts)
    if err != nil {
        log.Fatalf("Failed to register User gRPC-Gateway: %v", err)
    }

    // 注册 Reservation 服务的 gRPC-Gateway
    err = pb2.RegisterReservationHandlerFromEndpoint(context.Background(), mux, reservationGrpcEndpoint, opts)
    if err != nil {
        log.Fatalf("Failed to register Reservation gRPC-Gateway: %v", err)
    }

    err = pb3.RegisterRecommendationHandlerFromEndpoint(context.Background(), mux, recommendationGrpcEndpoint, opts)
    if err != nil {
        log.Fatalf("Failed to register Recommendation gRPC-Gateway: %v", err)
    }

    // 运行 HTTP 服务器，监听端口 9977
    log.Println("Starting HTTP server on port 9977...")
    if err := http.ListenAndServe(":9977", mux); err != nil {
        log.Fatalf("Failed to serve HTTP: %v", err)
    }
}
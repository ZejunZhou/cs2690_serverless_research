package user

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	pb "github.com/delimitrou/DeathStarBench/tree/master/hotelReservation/services/user/proto"
	pb2 "github.com/delimitrou/DeathStarBench/tree/master/hotelReservation/services/reservation/proto"
	pb3 "github.com/delimitrou/DeathStarBench/tree/master/hotelReservation/services/recommendation/proto"
)

const (
	shmName = "/home/zhouzejun1147/Ironfunctions-ServerlessResearch/hotel_user_shared/grpc_shared_mem" 
	shmSize = 1024              
)

func writeToSharedMemory(data string) error {
	fd, err := syscall.Open(shmName, syscall.O_CREAT|syscall.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("failed to open shared memory: %v", err)
	}
	defer syscall.Close(fd)

	if err := syscall.Ftruncate(fd, shmSize); err != nil {
		return fmt.Errorf("failed to resize shared memory: %v", err)
	}

	mem, err := syscall.Mmap(fd, 0, shmSize, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		return fmt.Errorf("failed to mmap shared memory: %v", err)
	}
	defer syscall.Munmap(mem)

	copy(mem, data)
	log.Println("Data written to shared memory:", data)
	return nil
}

type customResponseWriter struct {
	http.ResponseWriter
	body []byte
}

func (w *customResponseWriter) Write(data []byte) (int, error) {
	w.body = append(w.body, data...)
	return w.ResponseWriter.Write(data)
}

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customWriter := &customResponseWriter{ResponseWriter: w}
		handler.ServeHTTP(customWriter, r) 

		if err := writeToSharedMemory(string(customWriter.body)); err != nil {
			log.Printf("Failed to write to shared memory: %v", err)
		}
	})
}

func RunShared() {
	userGrpcEndpoint := "localhost:8086"          
	reservationGrpcEndpoint := "localhost:8087"  
	recommendationGrpcEndpoint := "localhost:8085"

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterUserHandlerFromEndpoint(context.Background(), mux, userGrpcEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register User gRPC-Gateway: %v", err)
	}

	err = pb2.RegisterReservationHandlerFromEndpoint(context.Background(), mux, reservationGrpcEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register Reservation gRPC-Gateway: %v", err)
	}


	err = pb3.RegisterRecommendationHandlerFromEndpoint(context.Background(), mux, recommendationGrpcEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register Recommendation gRPC-Gateway: %v", err)
	}

	log.Println("Starting HTTP server on port 9978...")
	if err := http.ListenAndServe(":9978", middleware(mux)); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}
}

package main

import (
	"fmt"
	"net/http"
	"strings"
	"syscall"
	"time"
)

const (
	shmName = "/app/grpc_shared_mem" 
	shmSize = 1024                  
)

func readFromSharedMemory() (string, error) {
	fd, err := syscall.Open(shmName, syscall.O_RDONLY, 0666)
	if err != nil {
		return "", fmt.Errorf("failed to open shared memory: %v", err)
	}
	defer syscall.Close(fd)

	mem, err := syscall.Mmap(fd, 0, shmSize, syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		return "", fmt.Errorf("failed to mmap shared memory: %v", err)
	}
	defer syscall.Munmap(mem)

	// Trim trailing null bytes
	data := strings.TrimRight(string(mem), "\x00")
	return data, nil
}

func sendRequestToGateway() {
	url := "http://10.128.0.2:9978/v1/user/check" // gRPC Gateway URL
	go func() {
		_, err := http.Get(url) 
		if err != nil {
			fmt.Printf("Failed to send request to gRPC-Gateway: %v\n", err)
		}
	}()
}

func main() {
	sendRequestToGateway()

	for i := 0; i < 10; i++ { 
		data, err := readFromSharedMemory()
		if err == nil && len(data) > 0 {
			fmt.Println(data)
			return
		}
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("Error: Failed to read valid data from shared memory within timeout.")
}

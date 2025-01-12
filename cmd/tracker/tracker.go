package main

import (
	"log"

	"github.com/xabsvoid/tracker/internal/app/grpc"
	geov1 "github.com/xabsvoid/tracker/internal/app/grpc/geo/v1"
)

func main() {
	server := grpc.NewServer("tcp", ":50051", geov1.NewService())

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

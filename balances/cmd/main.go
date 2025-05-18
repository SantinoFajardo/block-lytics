package main

import (
	"log"
	"net"
	"os"

	"github.com/block-lytics/balances/internal/server"
	balancespb "github.com/block-lytics/proto/balances/balancespb"
	"google.golang.org/grpc"
)

func main(){
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error escuchando: %v", err)
	}

	s := grpc.NewServer()
	balancespb.RegisterBalancesServiceServer(s, &server.BalancesServer{})

	log.Printf("Servidor gRPC escuchando en :%s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}
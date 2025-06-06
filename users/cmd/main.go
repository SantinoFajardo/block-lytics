package main

import (
	"log"
	"net"
	"os"

	userspb "github.com/block-lytics/proto/users/userspb"
	"github.com/block-lytics/users/internal/db"
	"github.com/block-lytics/users/internal/server"
	"google.golang.org/grpc"
)

func main() {
	dbConn := db.InitDB()
	defer dbConn.Close()
	userRepository := db.NewUserRepository(dbConn)
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error escuchando: %v", err)
	}

	s := grpc.NewServer()
	userspb.RegisterUsersServiceServer(s, &server.UsersServer{
		UserRepository: userRepository,
	})

	log.Printf("Servidor gRPC escuchando en :%s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}
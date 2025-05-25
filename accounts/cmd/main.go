package main

import (
	"log"
	"net"
	"os"

	"github.com/block-lytics/accounts/internal/db"
	"github.com/block-lytics/accounts/internal/server"
	accountspb "github.com/block-lytics/proto/accounts/accountspb"
	"google.golang.org/grpc"
)

func main() {
	dbConn := db.InitDB()
	defer dbConn.Close()
	accountRepository := db.NewAccountRepository(dbConn)
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error escuchando: %v", err)
	}

	s := grpc.NewServer()
	accountspb.RegisterAccountsServiceServer(s, &server.AccountsServer{
		AccountRepository: accountRepository,
	})

	log.Println("Servidor gRPC escuchando en :" + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}

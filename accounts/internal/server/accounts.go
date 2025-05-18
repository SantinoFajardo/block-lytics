package server

import (
	"context"
	"log"

	db "github.com/block-lytics/accounts/internal/db"
	accountspb "github.com/block-lytics/proto/accounts/accountspb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountsServer struct {
	accountspb.UnimplementedAccountsServiceServer
	AccountRepository db.AccountRepository
}

func (s *AccountsServer) GetAccountByAdress(ctx context.Context, req *accountspb.GetAccountByAdressRequest) (*accountspb.GetAccountByAdressResponse, error) {
	account, err := s.AccountRepository.GetByAddress(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Error al obtener la cuenta: %v", err)
	}

	log.Printf("Cuenta obtenida: %v", account)

	return &accountspb.GetAccountByAdressResponse{
		Id:      account.Id,
		Address: account.Address,
		Blockchain: account.Blockchain,
		UserId:     account.UserId,
		CreatedAt:  timestamppb.New(account.CreatedAt),
		UpdatedAt:  timestamppb.New(account.UpdatedAt),
	}, nil
}

func (s *AccountsServer) CreateAccount(ctx context.Context, req *accountspb.CreateAccountRequest) (*accountspb.CreateAccountResponse, error) {
	newAccountId, err := s.AccountRepository.Create(req.Address, req.UserId, req.Blockchain)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Error al crear la cuenta: %v", err)
	}

	return &accountspb.CreateAccountResponse{
		Id:      newAccountId,
		Message: "Cuenta creada correctamente",
	}, nil
}

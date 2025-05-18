package server

import (
	"context"
	"log"

	balancespb "github.com/block-lytics/proto/balances/balancespb"
)

type BalancesServer struct {
	balancespb.UnimplementedBalancesServiceServer
}

func (s *BalancesServer) UpdateBalances(ctx context.Context, req *balancespb.UpdateBalancesRequest) (*balancespb.UpdateBalancesResponse, error) {
	log.Printf("Updating balances for account: %s, blockchain: %s, token: %s, amount: %s", req.AccountAddress, req.Blockchain, req.TokenAddress, req.Amount)
	return &balancespb.UpdateBalancesResponse{Success: true}, nil
}



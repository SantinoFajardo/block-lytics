package server

import (
	"context"

	"github.com/block-lytics/balances/internal/db"
	balancespb "github.com/block-lytics/proto/balances/balancespb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BalancesServer struct {
	balancespb.UnimplementedBalancesServiceServer
	BalancesRepo db.BalancesRepository
}

func (s *BalancesServer) UpdateBalances(ctx context.Context, req *balancespb.UpdateBalancesRequest) (*balancespb.UpdateBalancesResponse, error) {
	success, err := s.BalancesRepo.UpdateUserBalances(req.AccountAddress, req.Blockchain, req.TokenAddress, req.Amount)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update balances: %v", err)
	}
	return &balancespb.UpdateBalancesResponse{Success: success}, nil
}



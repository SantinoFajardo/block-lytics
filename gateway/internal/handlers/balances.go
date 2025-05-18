package handlers

import (
	"context"
	"net/http"

	balancespb "github.com/block-lytics/proto/balances/balancespb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UpdateBalancesHandler(c *gin.Context) {
	var req struct {
		AccountAddress string `json:"account_address" binding:"required"`
		Blockchain     string `json:"blockchain" binding:"required"`
		TokenAddress   string `json:"token_address" binding:"required"`
		Amount         int64 `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conn, err := grpc.NewClient("balances:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to connect to balances service"})
		return
	}
	defer conn.Close()

	client := balancespb.NewBalancesServiceClient(conn)

	response, err := client.UpdateBalances(context.Background(), &balancespb.UpdateBalancesRequest{
		AccountAddress: req.AccountAddress,
		Blockchain:     req.Blockchain,
		TokenAddress:   req.TokenAddress,
		Amount:         req.Amount,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else if !response.Success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update balances"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Balances updated successfully!"})
}
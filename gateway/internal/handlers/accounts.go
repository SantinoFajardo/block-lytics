package handlers

import (
	"context"
	"net/http"
	"time"

	accountspb "github.com/block-lytics/proto/accounts/accountspb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetAccountByAdressHandler(c *gin.Context) {
	address := c.Param("address")

	conn, err := grpc.NewClient("accounts:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) // usar TLS en prod
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo conectar a accounts"})
		return
	}
	defer conn.Close()

	client := accountspb.NewAccountsServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.GetAccountByAdress(ctx, &accountspb.GetAccountByAdressRequest{Address: address})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      resp.Id,
		"address": resp.Address,
		"blockchain": resp.Blockchain,
		"user_id": resp.UserId,
		"created_at": resp.CreatedAt.AsTime(),
		"updated_at": resp.UpdatedAt.AsTime(),
	})
}

func CreateAccountHandler(c *gin.Context) {
	var req struct {
		Address    string `json:"address" binding:"required,hexadecimal"`
		Blockchain string `json:"blockchain" binding:"required,oneof=ethereum bitcoin"`
		UserId     int64  `json:"user_id" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conn, err := grpc.NewClient("accounts:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) // usar TLS en prod
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo conectar a accounts"})
		return
	}
	defer conn.Close()

	client := accountspb.NewAccountsServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.CreateAccount(ctx, &accountspb.CreateAccountRequest{
		Address:    req.Address,
		Blockchain: req.Blockchain,
		UserId:     req.UserId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      resp.Id,
		"message": resp.Message,
	})
}

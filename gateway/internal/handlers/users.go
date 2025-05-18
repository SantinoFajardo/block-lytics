package handlers

import (
	"context"
	"log"
	"net/http"

	userspb "github.com/block-lytics/proto/users/userspb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateUserHandler(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conn, err := grpc.NewClient("users:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to connect to users service"})
		return
	}
	defer conn.Close()

	log.Println("Connecting to users service at users:50052")
	client := userspb.NewUsersServiceClient(conn)
	log.Println("Connected to users service")
	response, err := client.CreateUser(context.Background(), &userspb.CreateUserRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": response.Id, "message": response.Message})
}
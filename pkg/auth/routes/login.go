package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stebin13/go-grpc-microservice/pkg/auth/pb"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context, authcli pb.AuthServiceClient) {
	var body LoginRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "binding error",
		})
		return
	}
	res, err := authcli.Login(context.Background(), &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Success": &res,
	})
}

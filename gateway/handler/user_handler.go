package handler

import (
	"github.com/gin-gonic/gin"
	pb "github.com/kianyari/microservice-practice/common/api"
)

type UserHandler struct {
	userClient pb.UserServiceClient
}

func NewUserHandler(
	userClient pb.UserServiceClient,
) *UserHandler {
	return &UserHandler{
		userClient: userClient,
	}
}

func (h *UserHandler) RegisterRoutes(ginEngine *gin.Engine) {
	userGroup := ginEngine.Group("/auth")
	{
		userGroup.POST("/register", h.Register)
		userGroup.POST("/login", h.Login)
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req pb.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.userClient.Register(c, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}
func (h *UserHandler) Login(c *gin.Context) {
	var req pb.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.userClient.Login(c, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}

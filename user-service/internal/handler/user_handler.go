package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/kianyari/microservice-practice/user-service/internal/dto"
	"github.com/kianyari/microservice-practice/user-service/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(
	userService service.UserService,
) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	type RegisterRequest struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	var registerRequest RegisterRequest
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	registerInfo := &dto.RegisterRequest{
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}
	err := h.userService.Register(registerInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})
}

func (h *UserHandler) Login(c *gin.Context) {
	type LoginRequest struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loginInfo := &dto.LoginRequest{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}
	token, err := h.userService.Login(loginInfo)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", token, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

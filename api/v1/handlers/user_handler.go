package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tylorkolbeck/go-cookbook/api/v1/dto"
	"github.com/tylorkolbeck/go-cookbook/auth"
	"github.com/tylorkolbeck/go-cookbook/internal/model"
	"github.com/tylorkolbeck/go-cookbook/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.service.ListUsers()

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not list users"})
		return
	}

	c.JSON(200, users)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := h.service.GetUserByID(id)

	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not marshal user"})
	}

	c.JSON(200, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req model.User
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	newUser, err := h.service.CreateUser(user)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, newUser)
}

func (h *UserHandler) Login(c *gin.Context, authConfig *auth.AuthConfig) {
	var req model.User

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	email := req.Email
	password := req.Password

	token, err := h.service.Login(email, password)
	// print the error
	println("ERROR: ", err)

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

func (h *UserHandler) VerifyEmail(c *gin.Context) {
	token := c.Param("token")

	success, err := h.service.VerifyEmail(token)

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not verify email"})
		return
	}

	if !success {
		c.JSON(404, gin.H{"error": "Could not verify email"})
		return
	}

	c.JSON(200, gin.H{"message": "Email verified"})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateUserRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := h.service.UpdateUser(id, req)

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not update user"})
		return
	}

	c.JSON(200, updatedUser)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	_, err := h.service.DeleteUser(id)

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not delete user"})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted"})
}

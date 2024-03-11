package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tylorkolbeck/go-cookbook/api/v1/dto"
	"github.com/tylorkolbeck/go-cookbook/internal/model"
	"github.com/tylorkolbeck/go-cookbook/internal/service"
)

type CookbookHandler struct {
	service *service.CookbookService
}

func NewCookbookHandler(service *service.CookbookService) *CookbookHandler {
	return &CookbookHandler{
		service: service,
	}
}

// CreateCookbook creates a new cookbook
func (h *CookbookHandler) CreateCookbook(c *gin.Context) {
	var req dto.AddCookbookRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCookbook, err := h.service.Add(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create cookbook"})
		return
	}

	c.JSON(200, newCookbook)
}

// GetCookbook retrieves a single cookbook by its ID
func (h *CookbookHandler) GetCookbook(c *gin.Context) {
	id := c.Param("id")

	cookbook, err := h.service.GetByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cookbook not found"})
		return
	}

	c.JSON(http.StatusOK, cookbook)
}

// ListCookbooks lists all cookbooks, with optional filters
func (h *CookbookHandler) ListCookbooks(c *gin.Context) {
	cookbooks := h.service.Get()
	c.JSON(http.StatusOK, cookbooks)
}

// UpdateCookbook updates an existing cookbook
func (h *CookbookHandler) UpdateCookbook(c *gin.Context) {
	var cookbook model.CookBook
	if err := c.BindJSON(&cookbook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCookbook, err := h.service.Update(c.Param("id"), cookbook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update cookbook"})
		return
	}

	c.JSON(200, newCookbook)
}

// DeleteCookbook deletes a cookbook by its ID
func (h *CookbookHandler) DeleteCookbook(c *gin.Context) {
	id := c.Param("id")
	id, err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cookbook not found"})
		return
	}

	c.JSON(http.StatusOK, id)
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tylorkolbeck/go-cookbook/api/v1/dto"
	"github.com/tylorkolbeck/go-cookbook/internal/service/recipe"
)

type RecipeHandler struct {
	service *recipe.RecipeService
}

func RegisterRecipeRoutes(router *gin.Engine, service *recipe.RecipeService) {
	handler := NewRecipeHandler(service)

	router.GET("/recipes", handler.ListRecipes)
	router.GET("/recipes/:id", handler.GetRecipe)
	router.PUT("/recipes/:id", handler.UpdateRecipe)
	router.DELETE("/recipes/:id", handler.DeleteRecipe)
	router.POST("/recipes", handler.CreateRecipe)
}

func NewRecipeHandler(service *recipe.RecipeService) *RecipeHandler {
	return &RecipeHandler{service: service}
}

// Get all recipes
func (h *RecipeHandler) ListRecipes(c *gin.Context) {
	recipes := h.service.Get()
	c.JSON(http.StatusOK, recipes)
}

// Get a single recipe by its ID
func (h *RecipeHandler) GetRecipe(c *gin.Context) {
	id := c.Param("id")

	recipe, err := h.service.GetByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

// Update a recipe by its ID
func (h *RecipeHandler) UpdateRecipe(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateRecipeRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedRecipe, err := h.service.Update(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update recipe"})
		return
	}

	c.JSON(http.StatusOK, updatedRecipe)
}

// Delete a recipe by its ID
func (h *RecipeHandler) DeleteRecipe(c *gin.Context) {
	id := c.Param("id")

	_, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	_, err = h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete recipe"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Recipe deleted"})
}

// Create a new recipe
func (h *RecipeHandler) CreateRecipe(c *gin.Context) {
	var req dto.CreateRecipeRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRecipe, err := h.service.Add(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create recipe"})
		return
	}

	c.JSON(200, newRecipe)
}

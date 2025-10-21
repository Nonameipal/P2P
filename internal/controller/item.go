package controller

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Nonameipal/P2P/internal/models/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OwnerResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type ItemResponse struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	PricePerDay float64       `json:"pricePerDay"`
	Status      string        `json:"status"`
	Category    string        `json:"category"`
	Owner       OwnerResponse `json:"owner"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
}

func (ctrl *Controller) GetItemByID(c *gin.Context) {
	id := c.Param("id")

	item, err := ctrl.service.GetItemByID(id)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, ItemResponse{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		Category:    strconv.Itoa(item.CategoryID),
		PricePerDay: item.PricePerDay,
		Status:      item.Status,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	})
}

func (ctrl *Controller) GetAllItems(c *gin.Context) {
	items, err := ctrl.service.GetAllItems()
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	var respItems []ItemResponse
	for _, item := range items {
		respItems = append(respItems, ItemResponse{
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			Category:    strconv.Itoa(item.CategoryID),
			PricePerDay: item.PricePerDay,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, respItems)
}

type CreateItemRequest struct {
	Title       string  `json:"title"       binding:"required"`
	Description string  `json:"description" binding:"required"`
	PricePerDay float64 `json:"pricePerDay" binding:"required,gt=0"`
	CategoryID  int     `json:"categoryId"  binding:"required"`
}

func (ctrl *Controller) CreateItem(c *gin.Context) {
	var input CreateItemRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		ctrl.handleError(c, err)
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		ctrl.handleError(c, errors.New("user not authenticated"))
		return
	}

	if err := ctrl.service.CreateItem(domain.Item{
		ID:          uuid.New().String(),
		OwnerID:     userID.(string),
		Title:       input.Title,
		Description: input.Description,
		PricePerDay: input.PricePerDay,
		CategoryID:  input.CategoryID,
		Status:      domain.ItemStatusActive,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, CommonResponse{Message: "Item created successfully!"})

}

type UpdateItemRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	PricePerDay float64 `json:"pricePerDay"`
	CategoryID  int     `json:"categoryId"`
	Status      string  `json:"status"`
}

func (ctrl *Controller) UpdateItem(c *gin.Context) {
	id := c.Param("id")

	var input UpdateItemRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		ctrl.handleError(c, err)
		return
	}


	userID, exists := c.Get("userID")
	if !exists {
		ctrl.handleError(c, errors.New("user not authenticated"))
		return
	}

	if err := ctrl.service.UpdateItemByID(domain.Item{
		ID:          id,
		OwnerID:     userID.(string),
		Title:       input.Title,
		Description: input.Description,
		PricePerDay: input.PricePerDay,
		CategoryID:  input.CategoryID,
		Status:      input.Status,
	}); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "Item updated successfully!"})
}

func (ctrl *Controller) DeleteItem(c *gin.Context) {
	id := c.Param("id")


	userID, exists := c.Get("userID")
	if !exists {
		ctrl.handleError(c, errors.New("user not authenticated"))
		return
	}

	if err := ctrl.service.DeleteItemByID(id, userID.(string)); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "Item deleted successfully!"})
}

// GetMyItems
// @Summary Получить мои товары
// @Description Получить все товары текущего пользователя
// @Tags Items
// @Produce json
// @Success 200 {object} []ItemResponse
// @Failure 401 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /api/v1/my-items [get]
func (ctrl *Controller) GetMyItems(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		ctrl.handleError(c, errors.New("user not authenticated"))
		return
	}

	items, err := ctrl.service.GetMyItems(userID.(string))
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	var respItems []ItemResponse
	for _, item := range items {
		respItems = append(respItems, ItemResponse{
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			Category:    strconv.Itoa(item.CategoryID),
			PricePerDay: item.PricePerDay,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, respItems)
}

// GetItemsByCategory
// @Summary Получить товары по категории
// @Description Получить все товары определенной категории
// @Tags Items
// @Produce json
// @Param category path string true "ID категории"
// @Success 200 {object} []ItemResponse
// @Failure 400 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /api/v1/items/category/{category} [get]
func (ctrl *Controller) GetItemsByCategory(c *gin.Context) {
	category := c.Param("category")

	items, err := ctrl.service.GetItemsByCategory(category)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	var respItems []ItemResponse
	for _, item := range items {
		respItems = append(respItems, ItemResponse{
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			Category:    strconv.Itoa(item.CategoryID),
			PricePerDay: item.PricePerDay,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, respItems)
}

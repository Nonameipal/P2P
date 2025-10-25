package controller

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Nonameipal/P2P/internal/models/domain"
	"github.com/gin-gonic/gin"
)

type OwnerResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type ItemResponse struct {
	ID            int           `json:"id"`
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	PricePerDay   float64       `json:"pricePerDay"`
	Status        string        `json:"status"`
	CategoryID    int           `json:"categoryId"`
	Owner         OwnerResponse `json:"owner"`
	AvailableFrom time.Time     `json:"availableFrom"`
	AvailableTo   time.Time     `json:"availableTo"`
	CreatedAt     time.Time     `json:"createdAt"`
	UpdatedAt     time.Time     `json:"updatedAt"`
}

func (ctrl *Controller) GetItemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.handleError(c, errors.New("invalid item ID"))
		return
	}

	item, err := ctrl.service.GetItemByID(id)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, ItemResponse{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		CategoryID:  item.CategoryID,
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
			CategoryID:  item.CategoryID,
			PricePerDay: item.PricePerDay,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, respItems)
}

type CreateItemRequest struct {
	Title         string  `json:"title"       binding:"required"`
	Description   string  `json:"description" binding:"required"`
	PricePerDay   float64 `json:"pricePerDay" binding:"required,gt=0"`
	CategoryName  string  `json:"categoryName" binding:"required"`
	AvailableFrom string  `json:"availableFrom" binding:"required"`
	AvailableTo   string  `json:"availableTo"   binding:"required"`
}

func (ctrl *Controller) CreateItem(c *gin.Context) {
	var input CreateItemRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		ctrl.handleError(c, errors.New("invalid input format"))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		ctrl.handleError(c, errors.New("user not authenticated"))
		return
	}

	// Проверяем формат дат
	availableFrom, err := time.Parse("02-01-2006", input.AvailableFrom)
	if err != nil {
		ctrl.handleError(c, errors.New("invalid date format for availableFrom. Use DD-MM-YYYY"))
		return
	}

	availableTo, err := time.Parse("02-01-2006", input.AvailableTo)
	if err != nil {
		ctrl.handleError(c, errors.New("invalid date format for availableTo. Use DD-MM-YYYY"))
		return
	}

	if availableTo.Before(availableFrom) {
		ctrl.handleError(c, errors.New("availableTo cannot be before availableFrom"))
		return
	}

	// Получаем категорию по имени
	category, err := ctrl.service.GetCategoryByName(input.CategoryName)
	if err != nil {
		ctrl.handleError(c, errors.New("category not found"))
		return
	}

	if err := ctrl.service.CreateItem(domain.Item{
		OwnerID:       userID.(int),
		Title:         input.Title,
		Description:   input.Description,
		PricePerDay:   input.PricePerDay,
		CategoryID:    category.ID,
		Status:        domain.ItemStatusActive,
		AvailableFrom: availableFrom,
		AvailableTo:   availableTo,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.handleError(c, errors.New("invalid item ID"))
		return
	}

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

	// Получаем текущий товар
	currentItem, err := ctrl.service.GetItemByID(id)
	if err != nil {
		ctrl.handleError(c, errors.New("item not found"))
		return
	}

	// Если указана новая категория, проверяем её существование
	if input.CategoryID != 0 {
		_, err := ctrl.service.GetCategoryByID(input.CategoryID)
		if err != nil {
			ctrl.handleError(c, errors.New("category not found"))
			return
		}
	} else {
		// Если категория не указана, используем текущую
		input.CategoryID = currentItem.CategoryID
	}

	updateItem := domain.Item{
		ID:         id,
		OwnerID:    userID.(int),
		CategoryID: input.CategoryID,
		Status:     currentItem.Status,
	}

	// Обновляем только те поля, которые были указаны
	if input.Title != "" {
		updateItem.Title = input.Title
	} else {
		updateItem.Title = currentItem.Title
	}

	if input.Description != "" {
		updateItem.Description = input.Description
	} else {
		updateItem.Description = currentItem.Description
	}

	if input.PricePerDay > 0 {
		updateItem.PricePerDay = input.PricePerDay
	} else {
		updateItem.PricePerDay = currentItem.PricePerDay
	}

	if input.Status != "" {
		updateItem.Status = input.Status
	}

	updateItem.AvailableFrom = currentItem.AvailableFrom
	updateItem.AvailableTo = currentItem.AvailableTo
	updateItem.UpdatedAt = time.Now()

	if err := ctrl.service.UpdateItemByID(updateItem); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "Item updated successfully!"})
}

func (ctrl *Controller) DeleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.handleError(c, errors.New("invalid item ID"))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		ctrl.handleError(c, errors.New("user not authenticated"))
		return
	}

	userIDStr := strconv.Itoa(userID.(int))
	if err := ctrl.service.DeleteItemByID(id, userIDStr); err != nil {
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

	items, err := ctrl.service.GetMyItems(strconv.Itoa(userID.(int)))
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
			CategoryID:  item.CategoryID,
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
	categoryID, err := strconv.Atoi(c.Param("category"))
	if err != nil {
		ctrl.handleError(c, errors.New("invalid category ID"))
		return
	}

	items, err := ctrl.service.GetItemsByCategory(strconv.Itoa(categoryID))
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
			CategoryID:  item.CategoryID,
			PricePerDay: item.PricePerDay,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, respItems)
}

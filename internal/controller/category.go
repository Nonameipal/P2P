package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Nonameipal/P2P/internal/models/domain"
	"github.com/gin-gonic/gin"
)

type CategoryResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}

type CreateCategoryRequest struct {
	Name        string `json:"name"        binding:"required"`
	Description string `json:"description" binding:"required"`
}

// CreateCategory
// @Summary Создать категорию
// @Description Создать новую категорию товаров (только для админов)
// @Tags Admin
// @Consume json
// @Produce json
// @Param request_body body CreateCategoryRequest true "информация о категории"
// @Success 201 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 401 {object} CommonError
// @Failure 403 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /admin/categories [post]
func (ctrl *Controller) CreateCategory(c *gin.Context) {
	var input CreateCategoryRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		ctrl.handleError(c, err)
		return
	}

	category := domain.Category{
		Name:        input.Name,
		Description: input.Description,
	}

	if err := ctrl.service.CreateCategory(category); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, CommonResponse{Message: "Category created successfully!"})
}

// GetAllCategories
// @Summary Получить все категории
// @Description Получить список всех категорий
// @Tags Categories
// @Produce json
// @Success 200 {object} []CategoryResponse
// @Failure 500 {object} CommonError
// @Router /categories [get]
func (ctrl *Controller) GetAllCategories(c *gin.Context) {
	categories, err := ctrl.service.GetAllCategories()
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	var respCategories []CategoryResponse
	for _, category := range categories {
		respCategories = append(respCategories, CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			CreatedAt:   category.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, respCategories)
}

// GetCategoryByID
// @Summary Получить категорию по ID
// @Description Получить информацию о категории по её ID
// @Tags Categories
// @Produce json
// @Param id path int true "ID категории"
// @Success 200 {object} CategoryResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /categories/{id} [get]
func (ctrl *Controller) GetCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.handleError(c, errors.New("invalid category ID"))
		return
	}

	category, err := ctrl.service.GetCategoryByID(id)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}

// UpdateCategory
// @Summary Обновить категорию
// @Description Обновить информацию о категории (только для админов)
// @Tags Admin
// @Consume json
// @Produce json
// @Param id path int true "ID категории"
// @Param request_body body CreateCategoryRequest true "новая информация о категории"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 401 {object} CommonError
// @Failure 403 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /admin/categories/{id} [put]
func (ctrl *Controller) UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.handleError(c, errors.New("invalid category ID"))
		return
	}

	var input CreateCategoryRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		ctrl.handleError(c, err)
		return
	}

	category := domain.Category{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
	}

	if err := ctrl.service.UpdateCategory(category); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "Category updated successfully!"})
}

// DeleteCategory
// @Summary Удалить категорию
// @Description Удалить категорию (только для админов)
// @Tags Admin
// @Produce json
// @Param id path int true "ID категории"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 401 {object} CommonError
// @Failure 403 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /admin/categories/{id} [delete]
func (ctrl *Controller) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.handleError(c, errors.New("invalid category ID"))
		return
	}

	if err := ctrl.service.DeleteCategory(id); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "Category deleted successfully!"})
}

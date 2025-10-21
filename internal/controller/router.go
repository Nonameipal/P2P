package controller

import (
	"fmt"
	"net/http"

	_ "github.com/Nonameipal/P2P/api/docs"
	"github.com/Nonameipal/P2P/internal/configs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (ctrl *Controller) InitRoutes() error {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", ctrl.ping)

	authG := r.Group("/auth")
	{
		authG.POST("/sign-up", ctrl.SignUp)
		authG.POST("/sign-in", ctrl.SignIn)
		authG.GET("/refresh", ctrl.RefreshTokenPair)
	}

	// Публичные маршруты
	publicG := r.Group("/api/v1")
	{
		publicG.GET("/categories", ctrl.GetAllCategories)
		publicG.GET("/categories/:id", ctrl.GetCategoryByID)
		publicG.GET("/items", ctrl.GetAllItems)
		publicG.GET("/items/:id", ctrl.GetItemByID)
		publicG.GET("/items/category/:category", ctrl.GetItemsByCategory)
	}

	// Маршруты для авторизованных пользователей
	apiV1G := r.Group("/api/v1", ctrl.checkUserAuthentication)
	{
		apiV1G.POST("/items", ctrl.CreateItem)
		apiV1G.PUT("/items/:id", ctrl.UpdateItem)
		apiV1G.DELETE("/items/:id", ctrl.DeleteItem)
		apiV1G.GET("/my-items", ctrl.GetMyItems)
	}

	// Админские маршруты
	adminG := r.Group("/api/v1/admin", ctrl.checkUserAuthentication, ctrl.checkIsAdmin)
	{
		// Управление категориями
		adminG.POST("/categories", ctrl.CreateCategory)
		adminG.PUT("/categories/:id", ctrl.UpdateCategory)
		adminG.DELETE("/categories/:id", ctrl.DeleteCategory)

}
	err := r.Run(fmt.Sprintf(":%s", configs.AppSettings.AppParams.PortRun))
	if err != nil {
		return err
	}

	return nil
}

// Ping
// @Summary     Healthcheck
// @Description Роут проверки сервиса
// @Tags        Ping
// @Produce     json
// @Success     200 {object} CommonResponse
// @Failure     500 {object} CommonError
// @Router      /ping [get]
func (ctrl *Controller) ping(c *gin.Context) {
	c.JSON(http.StatusOK, CommonResponse{
		Message: "Server is running",
	})
}

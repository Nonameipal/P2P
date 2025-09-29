package controller

import (
	"fmt"
	_ "github.com/Nonameipal/P2P/api/docs"
	"github.com/Nonameipal/P2P/internal/configs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func (ctrl *Controller) InitRoutes() error {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", ctrl.ping)

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

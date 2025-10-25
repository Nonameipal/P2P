package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Nonameipal/P2P/internal/models/domain"
	"github.com/Nonameipal/P2P/pkg"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userIDCtx           = "userID"
	userRoleCtx         = "userRole"
)

func (ctrl *Controller) checkUserAuthentication(c *gin.Context) {
	token, err := ctrl.extractTokenFromHeader(c, authorizationHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	userID, isRefresh, userRole, err := pkg.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	if isRefresh {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: "inappropriate token"})
		return
	}

	fmt.Printf("Debug - Token info: UserID=%d, Role=%s, IsRefresh=%v\n", userID, userRole, isRefresh)

	c.Set(userIDCtx, userID)
	c.Set(userRoleCtx, userRole)
}

func (ctrl *Controller) checkIsAdmin(c *gin.Context) {
	role := c.GetString(userRoleCtx)
	fmt.Printf("Debug - Role check: Current role=[%s], Required role=[%s], Equal=%v\n",
		role, domain.AdminRole, role == domain.AdminRole)

	if role == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: "role is not in context"})
		return
	}

	// Приводим обе роли к верхнему регистру для сравнения
	if strings.ToUpper(role) != "ADMIN" {
		fmt.Printf("Debug - Access denied: user role=[%s]\n", role)
		c.AbortWithStatusJSON(http.StatusForbidden, CommonError{Error: "permission denied"})
		return
	}

	c.Next()
}

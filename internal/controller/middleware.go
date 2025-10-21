package controller

import (
	"github.com/Nonameipal/P2P/internal/models/domain"
	"github.com/Nonameipal/P2P/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
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

	c.Set(userIDCtx, userID)
	c.Set(userRoleCtx, string(userRole))
}

func (ctrl *Controller) checkIsAdmin(c *gin.Context) {
	role := c.GetString(userRoleCtx)
	if role == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: "role is not in context"})
		return
	}

	if role != domain.AdminRole {
		c.AbortWithStatusJSON(http.StatusForbidden, CommonError{Error: "permission denied"})
		return
	}

	c.Next()
}

package middleware

import (
	"net/http"
	"st-portier-be/models"

	"github.com/gin-gonic/gin"
)

// RequireRole checks if the user has one of the specified roles
func RequireRole(allowedRoles ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		currentUser := user.(models.User)
		userRoleID := currentUser.RoleID

		// Check if the user's role is in the allowedRoles list
		for _, roleID := range allowedRoles {
			if userRoleID == roleID {
				c.Next()
				return
			}
		}

		// If none of the allowed roles match, deny access
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		c.Abort()
	}
}

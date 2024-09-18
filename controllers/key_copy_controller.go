package controllers

import (
	"net/http"
	"st-portier-be/models"
	"st-portier-be/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetKeyCopy fetches a key-copy by its ID
func GetKeyCopy(c *gin.Context) {
	keyCopyID, _ := strconv.Atoi(c.Param("id"))

	// Fetch the key-copy by ID
	keyCopy, err := services.GetKeyCopyByID(keyCopyID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Key-copy not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": keyCopy})
}

// GetKeyCopiesByLockID retrieves all rooms for a specific floor
func GetKeyCopiesByLockID(c *gin.Context) {
	user, _ := c.Get("user")
	lockID, _ := strconv.Atoi(c.Param("lock_id"))

	// Fetch the lock details
	lock, err := services.GetKeyCopyByID(lockID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Floor not found"})
		return
	}

	// Check if the user has access to the lock
	if user.(models.User).RoleID != models.SuperAdminRoleID && lock.Lock.Room.Floor.Building.CompanyID != user.(models.User).CompanyID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Fetch all rooms for the specified floor
	keyCopies, err := services.GetAllKeyCopiesByLock(lockID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get KeyCopies"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": keyCopies})
}

// GetAllKeyCopies fetches all key-copies for the user's company
func GetAllKeyCopies(c *gin.Context) {
	user, _ := c.Get("user")

	// Fetch all key-copies for the user's company
	keyCopies, err := services.GetAllKeyCopiesForCompany(user.(models.User).CompanyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch key-copies"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": keyCopies})
}

// GetAllKeyCopiesForSuperAdmin fetches all key-copies across all companies
func GetAllKeyCopiesForSuperAdmin(c *gin.Context) {
	// Fetch all key-copies (Super Admin access)
	keyCopies, err := services.GetAllKeyCopiesForSuperAdmin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch key-copies"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": keyCopies})
}

// CreateKeyCopy allows Admin and Normal Users to create key copies for their company
func CreateKeyCopy(c *gin.Context) {
	user, _ := c.Get("user")

	// Only Super Admin and Admin Users can create key copies
	if user.(models.User).RoleID != models.AdminRoleID && user.(models.User).RoleID != models.SuperAdminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var input models.KeyCopy
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to create the key copy, with permission checks
	if err := services.CreateKeyCopy(&input, user.(models.User).CompanyID); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Key copy created successfully", "data": input})
}

// UpdateKeyCopy allows Admin and Normal Users to update key copies for their company
func UpdateKeyCopy(c *gin.Context) {
	user, _ := c.Get("user")
	keyCopyID, _ := strconv.Atoi(c.Param("id"))

	// Only Super Admin and Admin can update key copies
	if user.(models.User).RoleID != models.AdminRoleID && user.(models.User).RoleID != models.SuperAdminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var input models.KeyCopy
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to update the key copy, with permission checks
	if err := services.UpdateKeyCopy(keyCopyID, &input, user.(models.User).CompanyID); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Key copy updated successfully", "data": input})
}

// DeleteKeyCopy allows Admin and Super Admin to delete key copies
func DeleteKeyCopy(c *gin.Context) {
	user, _ := c.Get("user")
	keyCopyID, _ := strconv.Atoi(c.Param("id"))

	// Only Admin and Super Admin can delete key copies
	if user.(models.User).RoleID != models.AdminRoleID && user.(models.User).RoleID != models.SuperAdminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Call the service to delete the key copy, with permission checks
	if err := services.DeleteKeyCopy(keyCopyID, user.(models.User).CompanyID); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Key copy deleted successfully"})
}

package controllers

import (
	"net/http"
	"st-portier-be/models"
	"st-portier-be/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateLock allows Admin or Super Admin to create a new lock for a room
func CreateLock(c *gin.Context) {
	user, _ := c.Get("user")

	// Only Admin or Super Admin can create locks
	if user.(models.User).RoleID != models.SuperAdminRoleID && user.(models.User).RoleID != models.AdminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var input models.Lock
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch the room to ensure it belongs to the user’s company
	room, err := services.GetRoomByID(input.RoomID)
	if err != nil || room.Floor.Building.CompanyID != user.(models.User).CompanyID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Room does not belong to your company"})
		return
	}

	// Call the service to create the lock
	if err := services.CreateLock(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create lock"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lock created successfully", "data": input})
}

// UpdateLock allows Admin or Super Admin to update a lock
func UpdateLock(c *gin.Context) {
	user, _ := c.Get("user")
	lockID, _ := strconv.Atoi(c.Param("id"))

	// Only Admin or Super Admin can update locks
	if user.(models.User).RoleID != models.SuperAdminRoleID && user.(models.User).RoleID != models.AdminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var input models.Lock
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch the lock to ensure it belongs to the user’s company through the room and building
	lock, err := services.GetLockByID(lockID)
	if err != nil || lock.Room.Floor.Building.CompanyID != user.(models.User).CompanyID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Lock does not belong to your company"})
		return
	}

	// Call the service to update the lock
	if err := services.UpdateLock(lockID, &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update lock"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lock updated successfully", "data": input})
}

// DeleteLock allows Super Admin to delete a lock
func DeleteLock(c *gin.Context) {
	user, _ := c.Get("user")
	lockID, _ := strconv.Atoi(c.Param("id"))

	// Only Super Admin can delete locks
	if user.(models.User).RoleID != models.SuperAdminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Call the service to delete the lock
	if err := services.DeleteLock(lockID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete lock"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lock deleted successfully"})
}

// GetLock retrieves a lock by its ID
func GetLock(c *gin.Context) {
	lockID, _ := strconv.Atoi(c.Param("id"))

	// Fetch the lock
	lock, err := services.GetLockByID(lockID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lock not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": lock})
}

// GetAllLocks fetches all rooms for Super Admin
func GetAllLocks(c *gin.Context) {
	// Fetch all rooms without restriction (Super Admin access)
	locks, err := services.GetAllLocks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch locks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": locks})
}

// GetAllLocksByRoomID allows Admin and Normal User to view all locks for a building
func GetAllLocksByRoomID(c *gin.Context) {
	user, _ := c.Get("user")
	roomID, _ := strconv.Atoi(c.Param("room_id"))

	// Fetch the room details
	room, err := services.GetRoomByID(roomID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	// Fetch the floor details
	floor, err := services.GetFloorByID(room.FloorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Floor not found"})
		return
	}

	// Check if the user has access to the floor
	if user.(models.User).RoleID != models.SuperAdminRoleID && floor.Building.CompanyID != user.(models.User).CompanyID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Fetch all locks for the specified building
	locks, err := services.GetAllLocksByRoomID(roomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch locks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": locks})
}

// GetAllLocksForSuperAdmin retrieves all locks for Super Admin
func GetAllLocksForSuperAdmin(c *gin.Context) {
	locks, err := services.GetAllLocksForSuperAdmin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch locks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": locks})
}

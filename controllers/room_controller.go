package controllers

import (
	"net/http"
	"st-portier-be/models"
	"st-portier-be/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateRoom allows Admin or Super Admin to create a new room
func CreateRoom(c *gin.Context) {
	user, _ := c.Get("user")

	// Only Admin or Super Admin can create rooms
	if user.(models.User).RoleID != models.SuperAdminRoleID && user.(models.User).RoleID != models.AdminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var input models.Room
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to create the room
	if err := services.CreateRoom(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create room"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room created successfully", "data": input})
}

// GetRoom retrieves a room by its ID, with access control based on the user's company
func GetRoom(c *gin.Context) {
	user, _ := c.Get("user")
	roomID, _ := strconv.Atoi(c.Param("id"))

	// Fetch the room details
	room, err := services.GetRoomByID(roomID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	// Check if the user can access the room
	floor, _ := services.GetFloorByID(room.FloorID)
	if user.(models.User).RoleID != models.SuperAdminRoleID && floor.Building.CompanyID != user.(models.User).CompanyID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": room})
}

// GetAllRooms fetches all rooms for Super Admin
func GetAllRooms(c *gin.Context) {
	// Fetch all rooms without restriction (Super Admin access)
	rooms, err := services.GetAllRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch rooms"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rooms})
}

// GetRoomsByCompany allows Admin and Normal User to view all floors for their company
func GetRoomsByCompany(c *gin.Context) {
	user, _ := c.Get("user")

	// Get all floors for the user's company
	rooms, err := services.GetAllRoomsByCompanyID(user.(models.User).CompanyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to Get rooms"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rooms})
}

// GetRoomsByFloorID retrieves all rooms for a specific floor
func GetRoomsByFloorID(c *gin.Context) {
	user, _ := c.Get("user")
	floorID, _ := strconv.Atoi(c.Param("floor_id"))

	// Fetch the floor details
	floor, err := services.GetFloorByID(floorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Floor not found"})
		return
	}

	// Check if the user has access to the floor
	if user.(models.User).RoleID != models.SuperAdminRoleID && floor.Building.CompanyID != user.(models.User).CompanyID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Fetch all rooms for the specified floor
	rooms, err := services.GetAllRoomsByFloor(floorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch rooms"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rooms})
}

// UpdateRoom allows Admin or Super Admin to update a room
func UpdateRoom(c *gin.Context) {
	user, _ := c.Get("user")
	roomID, _ := strconv.Atoi(c.Param("id"))

	// Only Admin or Super Admin can update rooms
	if user.(models.User).RoleID != models.SuperAdminRoleID && user.(models.User).RoleID != models.AdminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var input models.Room
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to update the room
	if err := services.UpdateRoom(roomID, &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update room"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room updated successfully", "data": input})
}

// DeleteRoom allows Super Admin to delete a room
func DeleteRoom(c *gin.Context) {
	user, _ := c.Get("user")
	roomID, _ := strconv.Atoi(c.Param("id"))

	// Only Super Admin can delete rooms
	if user.(models.User).RoleID != models.SuperAdminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if err := services.DeleteRoom(roomID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete room"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room deleted successfully"})
}

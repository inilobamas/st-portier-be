package controllers

import (
	"net/http"
	"st-portier-be/models"
	"st-portier-be/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateFloor allows Admin or Super Admin to create a new floor
func CreateFloor(c *gin.Context) {
	user, _ := c.Get("user")

	// Only Admin or Super Admin can create floors
	if userRole := user.(models.User).RoleID; userRole != models.SuperAdminRoleID && userRole != models.AdminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var input models.Floor
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the building belongs to the user's company
	if building, err := services.GetBuildingByID(int(input.BuildingID)); err != nil || building.CompanyID != user.(models.User).CompanyID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Building not found or access denied"})
		return
	}

	// Call the service to create the floor
	if err := services.CreateFloor(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create floor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Floor created successfully", "data": input})
}

// GetFloors allows Admin and Normal User to view all floors for their company
func GetFloors(c *gin.Context) {
	user, _ := c.Get("user")

	// Get all floors for the user's company
	floors, err := services.GetAllFloorsByCompany(user.(models.User).CompanyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to Get floors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": floors})
}

// GetFloor allows Admin and Normal User to view their company, Super Admin can view any company
func GetFloor(c *gin.Context) {
	user, _ := c.Get("user")
	floorID, _ := strconv.Atoi(c.Param("id"))

	// Fetch the floor by ID
	floor, err := services.GetFloorByID(floorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Floor not found"})
		return
	}

	// Super Admin can view any floor, so no additional checks needed
	if user.(models.User).RoleID == models.SuperAdminRoleID {
		c.JSON(http.StatusOK, gin.H{"data": floor})
		return
	}

	// For Admin and Normal Users, check if the floor's building belongs to their company
	building, err := services.GetBuildingByID(int(floor.BuildingID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Building not found"})
		return
	}

	// Check if the building belongs to the user's company
	if building.CompanyID != user.(models.User).CompanyID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// If the building belongs to the user's company, return the floor
	c.JSON(http.StatusOK, gin.H{"data": floor})
}

// GetFloorsByBuildingID allows Admin and Normal User to view all floors for a building
func GetFloorsByBuildingID(c *gin.Context) {
	user, _ := c.Get("user")
	buildingID, _ := strconv.Atoi(c.Param("building_id"))

	// Fetch the building details
	building, err := services.GetBuildingByID(buildingID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Building not found"})
		return
	}

	// Check if the user has access to the building
	if user.(models.User).RoleID != models.SuperAdminRoleID && building.CompanyID != user.(models.User).CompanyID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Fetch all floors for the specified building
	floors, err := services.GetAllFloorsByBuilding(buildingID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch floors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": floors})
}

// UpdateFloor allows Admin or Super Admin to update a floor
func UpdateFloor(c *gin.Context) {
	user, _ := c.Get("user")
	floorID, _ := strconv.Atoi(c.Param("id"))

	// Only Admin or Super Admin can update floors
	if user.(models.User).RoleID != models.SuperAdminRoleID && user.(models.User).RoleID != models.AdminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var input models.Floor
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the building belongs to the user's company
	if building, err := services.GetBuildingByID(int(input.BuildingID)); err != nil || building.CompanyID != user.(models.User).CompanyID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Building not found or access denied"})
		return
	}

	// Call the service to update the floor
	if err := services.UpdateFloor(floorID, &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update floor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Floor updated successfully", "data": input})
}

// DeleteFloor allows Super Admin to delete a floor
func DeleteFloor(c *gin.Context) {
	user, _ := c.Get("user")
	floorID, _ := strconv.Atoi(c.Param("id"))

	// Only Super Admin can delete floors
	if user.(models.User).RoleID != models.SuperAdminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if err := services.DeleteFloor(floorID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete floor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Floor deleted successfully"})
}

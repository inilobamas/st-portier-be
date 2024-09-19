package controllers

import (
	"net/http"
	"st-portier-be/models"
	"st-portier-be/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateBuilding allows Admin or Super Admin to create a new building
func CreateBuilding(c *gin.Context) {
	user, _ := c.Get("user")

	// Only Admin or Super Admin can create buildings
	if user.(models.User).RoleID != models.SuperAdminRoleID && user.(models.User).RoleID != models.AdminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var input models.Building
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.(models.User).RoleID != models.SuperAdminRoleID {
		// Set the company ID of the employee to the logged-in user's company
		input.CompanyID = user.(models.User).CompanyID
	}

	// Call the service to create the building
	if err := services.CreateBuilding(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create building"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Building created successfully", "data": input})
}

// GetAllBuildings fetches all rooms for Super Admin
func GetAllBuildings(c *gin.Context) {
	user, _ := c.Get("user")
	buildingID, _ := strconv.Atoi(c.Param("id"))

	var buildings *models.Building[]
	var err error

	// Super Admin can view any company
	if user.(models.User).RoleID == models.SuperAdminRoleID {
		buildings, err = services.GetAllBuildings()
	} else {
		// Admin and Normal User can only view their own company
		buildings, err = services.GetAllBuildingsByCompany(user.(models.User).CompanyID, buildingID)
	}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Building not found"})
		return
	}


	c.JSON(http.StatusOK, gin.H{"data": buildings})
}

// GetBuildingsByCompany allows Admin and Normal User to view all buildings for their company
func GetBuildingsByCompany(c *gin.Context) {
	user, _ := c.Get("user")

	// Get all buildings for the user's company
	buildings, err := services.GetAllBuildingsByCompany(user.(models.User).CompanyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to Get buildings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": buildings})
}

// GetBuilding allows Admin and Normal User to view their company, Super Admin can view any company
func GetBuilding(c *gin.Context) {
	user, _ := c.Get("user")
	buildingID, _ := strconv.Atoi(c.Param("id"))

	var building *models.Building
	var err error

	// Super Admin can view any company
	if user.(models.User).RoleID == models.SuperAdminRoleID {
		building, err = services.GetBuildingByID(buildingID)
	} else {
		// Admin and Normal User can only view their own company
		building, err = services.GetBuildingByIDAndUserCompany(user.(models.User).CompanyID, buildingID)
	}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Building not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": building})
}

// UpdateBuilding allows Admin or Super Admin to update a building
func UpdateBuilding(c *gin.Context) {
	user, _ := c.Get("user")
	buildingID, _ := strconv.Atoi(c.Param("id"))

	// Only Admin or Super Admin can update buildings
	if user.(models.User).RoleID != models.SuperAdminRoleID && user.(models.User).RoleID != models.AdminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var input models.Building
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to update the building
	if err := services.UpdateBuilding(buildingID, &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update building"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Building updated successfully", "data": input})
}

// DeleteBuilding allows Super Admin to delete a building
func DeleteBuilding(c *gin.Context) {
	user, _ := c.Get("user")
	buildingID, _ := strconv.Atoi(c.Param("id"))

	// Only Super Admin can delete buildings
	if user.(models.User).RoleID != models.SuperAdminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if err := services.DeleteBuilding(buildingID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete building"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Building deleted successfully"})
}
